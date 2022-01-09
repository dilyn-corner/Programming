## Concurrency is not Parallelism

Go has some great concurrency support. This support is manifested in goroutines
and channels. So, in order to master concurrency, we have to discuss channels.

### Goroutines
A goroutine is a function that is capable of running concurrently with other
functions. All you have to do to create a goroutine is use the keyword `go`,
followed by the function:

```
package main

import "fmt"

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
    }
}

func main() {
    go f(0)
    var input string
    fmt.Scanln(&input)
}
```

How many goroutines are in this program? If you said one, you're wrong!

`go f(0)` is one goroutine. The other goroutine is implicit; it's `main()`.

A goroutine essentially makes it so that, instead of waiting for the line to
finish executing before moving on, once the line is reached we immediately move
to the next line of our code. In fact, because we don't wait for the function to
complete, our program would immediately exit without the `fmt.Scanln(&input)`
function! Test it for yourself; without this line, the program will exit before
all of the numbers get printed!

Goroutines are lightweight, and you can create essentially an arbitrary number
of them. Indeed, we can modify the above program to make even greater use of
goroutines:

```
func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
```

However, if you run this program it might appear as though the goroutines are
running sequentially instead of concurrently. If you sprinkle in some random
wait times you'll see that in fact, they *are* happening concurrently:

```
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
```

`f` will print numbers from 0 to 10, and will wait anywhere between 0 and 250ms
after each one. You'll see now that the goroutines are running simultaneously!


### Channels

Channels are key to goroutine communication. Channels ensure that two goroutines
can synchronize their execution by allowing them to communicate between each
other. Here is an example:

```
package main

import (
    "fmt"
    "time"
)

func pinger(c chan string) {
    for i := 0; ; i++ {
        c <- "ping"
    }
}

func printer(c chan string) {
    for {
        msg := <- c
        fmt.Println(msg)
        time.Sleep(time.Second * 1)
    }
}

func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}
```

This simple program will print "ping" forever (can be stopped by hitting enter).
A channel type is represented by the keyword `chan` (so it is much like any
other type in that way), followed by the type of the things that are passed on
that channel (so in this case, `c chan string` means that c is a channel which
has strings travel through it). 

The `<-` operator is used to send and receive messages on a channel.

Syntactically, `c <- "ping"` means "send 'ping'". 

`msg := <- c` means "receive a message and store it in `msg`". 

Instead of writing `fmt.Println(msg)` we could have written `fmt.Println(<-c)`,
eliminating the need for `msg` (thus, channels behave in expected ways).

This channel `c` acts to synchronize two goroutines. `pinger()` will attempt to
send a message across the channel (`"ping"`), but it will wait until `printer()`
is ready to receive the message. In other words, `pinger()` is being blocked
until a relevant condition happens which unblocks it. Blocking. Beautiful.

We can add a third function to this program and see what happens:
```
func ponger(c chan string) {
    for i := 0; ; i++ {
        c <- "pong"
    }
}
```

And modify `main()`:

```
func main() {
    var c chan string = make(chan string)

    go pinger(c)
    go ponger(c)
    go printer(c)

    var input string
    fmt.Scanln(&input)
}
```

And with this, the program will take turns printing "ping" and "pong". You can
think about it like this:

The goroutines run concurrently, but they are called sequentially (because we
read and execute line to line), so `pinger()` is the first goroutine to execute,
followed by `ponger()`. When `pinger()` goes, the channel `c` holds `"ping"`.
`ponger()` is thus blocked from sending its message across the channel `c`
because the channel is currently occupied (`ponger()` is being blocked). When
the `printer()` goroutine executes, the message in `c` gets taken out and
printed. Because the channel is now empty, `ponger()` is unblocked and sends
"pong" down the channel `c`. And this repeats, ad infinitum.


The channel `c` we have created here is known as a bidirectional channel; it can
both send and receive things (in this case, strings). However, we can limit a
channel to being only a sender or a receiver:

```func pinger(c chan<- string)```

This means that `c` can only be sent to. Attempting to receive from `c` will
result in a compile time error. Similarly, we can change `printer()`:

``` func printer(c <-chan string)```

And thus `c` can only receive, not send!

### Select Statements

Because channels are special, we cannot use `switch`. Instead, we use `select`:

```
func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        for {
            c1 <- "from 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func() {
        for {
            c2 <- "from 2"
        }
    }()

    go func() {
        for {
            select {
                case msg1 := <- c1:
                    fmt.Println(msg1)
                case msg2 := <- c2:
                    fmt.Println(msg2)
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}
```

This program will print "from 1" every 2 seconds and "from 2" every 3 seconds.
`select` picks the first channel that is ready and receives from it (or it sends
to it). If more than one channel is ready, it randomly picks one. If none of the
channels are ready, the statement blocks until one becomes available, unless
there is a `default` case, in which `default` is chosen immediately.


A common use of the `select` statement is to implement a timeout:
```
select {
    case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
    case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
    case <- time.After(time.Second):
        fmt.Println("timeout")
}
```

`time.After()` creates a channel and after the given duration will send the
current time on it. Here is with a default case:
```
select {
    case msg1 := <- c1:
        fmt.Println("Message 1", msg1)
    case msg2 := <- c2:
        fmt.Println("Message 2", msg2)
    case <- time.After(time.Second):
        fmt.Println("timeout")
    default:
        fmt.Println("nothing ready")
}
```


When making channels using the `make()` builtin, it's possible to pass a second
parameter!
```c := make(chan int, 1)```

This creates a *buffered channel* with a capacity of 1. Normally channels are
synchronous; both sides of the channel will wait until the other side is ready.
A buffered channel is asynchronous; sending or receiving a message will not wait
unless the channel is already full!



These notes come essentially verbatim from [here](https://www.golang-book.com/books/intro/10)!
