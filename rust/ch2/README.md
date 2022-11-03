## Programming a Guessing Game

In this chapter we will do our first hands-on project to explore the syntax of
Rust and some important concepts and features of the language. Here we will
learn about `let`, `match`, methods, associated functions, external crates, and
a fair bit more. These things will all be built upon in later chapters. Here, we
will be focusing on the fundamentals.

Here we will create a program which will generate a random integer between 1 and
100, and then it will prompt the player to guess a number. The program will
indicate if the guess is too low or too high, and if they guessed correctly the
program will congratulate the player on a job well done and exit.

First, we need to setup our project. Do `cargo new guessing_game` to generate
the boilerplate.

### Processing a guess

Replace the content of `guessing_game/src/main.rs` with:

```
use std::io;

fn main() {
    println!("Guess the number!");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {guess}");
}
```

This snippet contains a fair bit of new stuff, so let's get cracking!

Rust's standard library contains a lot of useful things for us to use. One
specific example is a way of getting input from users. This is defined in the
`io` library. In order to add this library to our program, we use:

```
use std::io;
```

Note that by default, Rust has a set of items defined in the standard library
which it brings into the scope of every Rust program. This set of items is
called the *prelude*, and you can learn more about it
[here](https://doc.rust-lang.org/stable/std/prelude/index.html).


#### Storing Values with Variables

Another new thing added here we haven't seen before is

```
let mut guess = String::new();
```

`let` is used in Rust to create a variable. Variables are immutable in Rust,
meaning that once they have a value this value will not change. For instance,
`let apples = 5` creates a variable named `apples` and assigns it the value `5`.
If we wanted to make a variable *mutable*, we use `mut` when we declare it.
Thus, `guess` is a new mutable variable (more on mutability in `../ch3`).

We immediately bind a value to our `guess` variable using `=`. This value is
the result of calling `String::new()`, a function that returns a new instance of
`String`. `String` is a *string type* provided by `std` that is a growable,
UTF-8 encoded text.

The `::` syntax indicates that `new` is an associated function of the `String`
type. An *associated function* is a function that's implemented on a type. The
`new` function thus creates a new and empty string. `new` is a function on many
types.

Thus, the `let mut guess = String::new();` line creates a *mutable variable*
that is *currently bound* to a *new*, *empty* instance of a *`String`*.

#### Receiving User Input

Now we get to use the input/output functions in Rust! Because of `use std::io`,
we can handle user input with the lines:

```
    io::stdin()
        .read_line(&mut guess)
```

Note that if we hadn't used `use std::io` at the top of this program, we could
call the function from the standard library directly by instead using
`std::io::stdin()`.

The `stdin` function returns an instance of `std::io::Stdin`, which is a type
that represents a handle to the standard input for your terminal.

`.read_line(&mut guess)` calls the `read_line` method on the standard input
handle to get an input from the user (lots of methods in here!). `&mut guess` is
the argument passed to `read_line`, and this tells `read_line` what string to
store in the `guess` variable! `read_line`'s full job is to take whatever the
user types into standard input and *append it* into a string, __not overwriting
that string's contents__.

The `&` indicates that this argument is a *reference*. In short, by using a
reference we can let multiple parts of our program use data stored in memory
without copying that data into memory multiple times for each part of the code
to access. References are a complex feature, and one of Rust's major advantages
is how safe and easy it is to use references. References are explained more in
`../ch4`, but for now all that matters is that because references are
*immutable* by default, we have to use `&mut guess` instead of `&guess` to mark
the reference as a mutable one.

#### Handling Potential Failure with the Result Type

The next part of the previous method,

```
    .expect("Failed to read line");
```

could have been written as

```
io::stdin().read_line(&mut guess).expect("Failed to read line");
```

However, one line is difficult to read -- so this choice is purely a design one.
Often, it's wise to introduce a newline and other whitespace to break up long
lines when you call a method with the `.method_name()` syntax.

As mentioned earlier, `read_line` appends the input to the string passed to it.
It also returns a `Result` value. `Result` is an *enumeration*, often called an
*enum*, which is a type that can be in one of multiple possible states. Each
possible state is called a *variant*. Enums are covered in more detail in
`../ch6`. The `Result` type is used to encode error-handling information.

`Result`'s variants are `Ok` and `Err`. `Ok` indicates success, and inside of
`Ok` is the successfully generated value. `Err` indicates failure, and `Err`
contains information about how or why the operation failed.

Like any other type, values of the `Result` variant have defined methods. An
instance of `Result` has an `expect` method which can be called. If `Result` is
an `Err`, `expect` will cause the program to crash and display the message
passed as an argument to `expect`. If the `read_line` method returns an `Err`,
it would likely be the result of an error coming from the underlying operating
system. If the `Result` is `Ok`, `expect` will take the return value that `Ok`
is holding and return just that value ot you so that it can be used. In this
case, that value is the number of bytes in the user's input.

If you don't call `expect`, the code will build but give a warning:

```
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
warning: unused `Result` that must be used
  --> src/main.rs:10:5
   |
10 |     io::stdin().read_line(&mut guess);
   |     ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
   |
   = note: `#[warn(unused_must_use)]` on by default
   = note: this `Result` may be an `Err` variant, which should be handled

warning: `guessing_game` (bin "guessing_game") generated 1 warning
    Finished dev [unoptimized + debuginfo] target(s) in 0.59s
```

Rust is warning that the `Result` value from `read_line` is not being used,
which means that a possible error isn't being handled. The *correct way* to
supress this warning is to actually write error handling.

#### Printing values with `println!` Placeholders

The last line,

```
    println!("You guessed: {guess}");
```

prints the string (`guess`) that now contains the user's input. The `{ }` is a
placeholder; `{ }` holds a value in place. You can print more than one value
using curly brackets: the first set of curly brackets holds the first value
listed after the format string, the second set holds the second value,a dn so
on. Printin gmultiple values in one call to `println!` would look like:

```
let x = 5;
let y = 10;
println!("x = {} and y = {}", x, y);
```

which will print `x = 5 and y = 10`.

#### Testing the First Part

Testing the first part of the guessing game can be done using `cargo run`,
outputting (with some light interaction):

```
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
    Finished dev [unoptimized + debuginfo] target(s) in 6.44s
     Running `target/debug/guessing_game`
Guess the number!
Please input your guess.
6
You guessed: 6
```

So thus far, we have success! We accept input, and we print something to the
screen. Now for the more salient logic!

### Generating a Secret Number

Now we have to generate the secret number players will be trying to guess. The
secret number should be different every time (how boring if it weren't!). Random
number generation is not currently in the standard library, but there is a crate
with this functionality available which can be used. It's usefully named `rand`.

#### Using a Crate to Get More Functionality

A crate is just another Rust project. Crates can be libraries or binaries. In
this case, `rand` is a crate containing a library which our binary program can
use. In order to use this crate, we will have to modify our
`guessing_game/Cargo.toml` to include the `rand` crate as a dependency. Add the
following line in the `[dependencies]` section header:

```
rand = "0.8.3"
```

Here we are specifying that the crate `random` we use should be the version
`0.8.3`. Technically we are using a shorthand of `^0.8.3`, which means *any
version at least `0.8.3` but less than `0.9.0`*. We are leveraging the meaning
of semantic versioning here; adhering to the `Major.Minor.Patch` convention
should ensure that we use any avaialable updated version of the `rand` crate
which is still API compatible with the `0.8.0` release. Now we can try to build
our project:

```
    Updating crates.io index
  Downloaded rand_chacha v0.3.1
  Downloaded ppv-lite86 v0.2.16
  Downloaded rand v0.8.5
  Downloaded rand_core v0.6.4
  Downloaded cfg-if v1.0.0
  Downloaded getrandom v0.2.8
  Downloaded libc v0.2.137
  Downloaded 7 crates (791.9 KB) in 0.83s
   Compiling libc v0.2.137
   Compiling cfg-if v1.0.0
   Compiling ppv-lite86 v0.2.16
   Compiling getrandom v0.2.8
   Compiling rand_core v0.6.4
   Compiling rand_chacha v0.3.1
   Compiling rand v0.8.5
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
    Finished dev [unoptimized + debuginfo] target(s) in 7.36s
```

Cargo fetches the latest versions of our specified dependencies *and those
dependencies dependencies*. They are fetched from
[Crates.io](https://crates.io), which is where people post their open source
Rust projects for others to use.

#### Ensuring Reproducible Builds with the `Cargo.lock` File

Cargo ensures that rebuilds of a project will always build the same artifact for
anyone who builds it. It does this by using a `Cargo.lock` file, which will hold
all of the dependency crates to a particular version. This way, if `rand 0.8.6`
were released and actually had a regression, it wouldn't impact our project.
Until, of course, we generate a new `Cargo.lock` file (for instance, by
explicitly upgrading the crate in our `[dependencies]` section of the
`Cargo.toml`).

#### Updating a Crate to Get a New Version

When you *do* want to update a crate, `cargo update` will do the work for you!
This command ignores the `Cargo.lock` file and find the versions of crates
specified in the `Cargo.toml`, updating the `Cargo.lock` file.

If the `rand` crate had two new versions, `0.8.6` and `0.9.0`, only the `0.8.6`
crate would be added - Cargo ignores the `0.9.0` release. If instead you wanted
to jump to the `0.9.0` release, you will have to modify the `Cargo.toml` file to
specify the version.

### Generating a Random Number

Update `guessing_game/src/main.rs` to the following:

```
use std::io;
use rand::Rng;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    println!("The secret number is: {secret_number}");

    println!("Please input your guess.");

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    println!("You guessed: {guess}");
}
```

The first addition is `use rand::Rng;`. The `Rng` *trait* defines the methods
that random number generators implement, adn this trait must be *in scope* for
us to use this method; traits are covered in more detail in `../ch10`.

Then we add a few lines in the middle. The first line, `rand::thread_rng` is
called to give us the particular random number generator that we're going to
use: the one that is local to the current thread of execution and seeded by the
operating system. From there, we call the `gen_range` method on this random
number generator function. This method is defined by the `Rng` trait brought
into scope with the `use rand::Rng` statement. The `gen_range` method takes as
an argument a range expression and generates a random number in that range. The
expression we use here is of the form `start..=end`, which is inclusive of both
the lower and upper bounds. Thus, we specify `1..=100` to generate a random
number between 1 and 100.

Note: it's never really obvious which traits can be used with which methods and
functions from any given crate. Thus, each crate has (and should have)
documentation with instructions for use. The `cargo doc --open` command will
build the documentation for every dependency locally and open it in your
browser. You can find the other features of the `rand` crate running this
command!

The other new line, `println!("The secret number is: {secret_number}");`, is
only really useful here during development work; the game wouldn't be very fun
if it immediately gave away the answer, after all!

### Comparing the Guess to the Secret Number

Now that we have the major components of the game, fetching user input and
generating a random number, we can implement the meat of the game: comparison.

This next code block is written to intentionally cause a failure to demonstrate
an important feature of Rust; we will fix this soon.

Note that the `// --snip--`s you may see in code blocks are not generally added
to the files, but for readability.

Update `guessing_game/src/main.rs`:

```
use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    // --snip--

    println!("You guessed: {guess}");

    match guess.cmp(&secret_number) {
        Ordering::Less => println!("Too small!"),
        Ordering::Greater => println!("Too big!"),
        Ordering::Equal => println!("You win!"),
    }
}
```

First we add `use std::cmp::Ordering;` which brings a useful type into scope for
our program. The `Ordering` type is another enum and has the variants `Less`,
`Greater`, and `Equal`. These are the three outcomes possible when comparing two
values.

Then we add some new lines at the bottom which leverage the newly introduced
`Ordering` type. The `cmp` method compares two values and can be called on
anything which can be compared. It takes a reference to your point of comparison
(here, `&secret_number`), and compares it with `guess`. It will then return a
variant of the `Ordering` enum brought into scope by the `use` statement. We
leverage the `match` expression to decide what to do based on the variant of
`Ordering`; if the output of `guess.cmp` *`match`es* the variant `Less`, then
the result of the `match` will be `println!("Too small!");`. Likewise for the
other two variants.

A `match` expression is made up of *arms* (lovely term of art). An arm consists
of a *pattern* to *match* against, and the code that should be run if the value
given to `match` fits that arm's pattern. The patterns are checked in turn.
Patterns and the `match` construct are __powerful__ Rust features covered in
`../ch6` and `../ch18`.

Use `cargo build` to try and build our code. Note that it won't compile:

```
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
error[E0308]: mismatched types
   --> src/main.rs:22:21
    |
22  |     match guess.cmp(&secret_number) {
    |                 --- ^^^^^^^^^^^^^^ expected struct `String`, found integer
    |                 |
    |                 arguments to this function are incorrect
    |
    = note: expected reference `&String`
               found reference `&{integer}`
note: associated function defined here
   --> /home/dilyn/.rustup/toolchains/stable-x86_64-unknown-linux-gnu/lib/rustlib/src/rust/library/core/src/cmp.rs:785:8
    |
785 |     fn cmp(&self, other: &Self) -> Ordering;
    |        ^^^

For more information about this error, try `rustc --explain E0308`.
error: could not compile `guessing_game` due to previous error
```

The core of the error is that there are *mismatching types*. Rust uses a strong,
static typing system. However, it does have type inference. So Rust inferred
from `let mut guess = String::new()` that `guess` is a `String` and didn't force
us to write this explicitly. However, `secret_number` was inferred to be a
number type (and rightfully so). Rust supports several number types, and the
ones which have a value between 1 and 100 include `i32` (a 32-bit number), `u32`
(an unsigned 32-bit number), `i64` (a 64-bit number), and a few others. Because
we didn't specify a type, Rust defaults to an `i32`. Thus, the reason for the
error is that we are trying to compare a `String` to an `i32`; Rust cannot
compare a string and a number type, and will not convert between them for you.

We can convert the string inputted by players to an `i32` to allow the guess to
be compared to the generated number. We can do this by modifying
`guessing_game/src/main.rs` to include:

```
    // --snip--

    let mut guess = String::new();

    io::stdin()
        .read_line(&mut guess)
        .expect("Failed to read line");

    let guess: u32 = guess.trim().parse().expect("Please type a number!");

    println!("You guessed: {guess}");

    match guess.cmp(&secret_number) {
        Ordering::Less => println!("Too small!"),
        Ordering::Greater => println!("Too big!"),
        Ordering::Equal => println!("You win!"),
    }
```

The line added is `let guess: u32 = guess.trim().parse().expect("Please type a number!");`

"But wait!", you may cry; "don't we already use a variable named `guess`??". Why
certainly it does, astute reader. But Rust helpfully allows us to *shadow* the
previous value of `guess` with a new one. Shadowing allows us to reuse the
`guess` variable name instead of forcing us to create, say, two new variables
like `guess_string` and `guess_number`. This is covered more in `../ch3`, but
just know that this shadowing allows us to convert a value from one type to
another. The work is done in the `: u32` bit, and we use various methods
(`.trim()`, `.parse()`, `.expect()`) to clean up the guess and error check.

The `trim` method on a `String` eliminates any white space at the beginning
and end, as our string should only contain numerical data. Because our program
relies on user input and submitting this input requires pressed Return, a
newline character is appended to the user's input. Thus, we have to trim it.
`trim` eliminates the newline character, changing the user's input from `foo\n`
to `foo`.

The `parse` method converst a string to another type. Here, we use it to convert
from a string to a number. We tell Rust what it should be converted into with
`let guess: u32`. Note that Rust has several types for numbers; here we use
`u32` as it is a sane choice for small, positive numbers. Additionally, because
the `u32` typed `guess` is compared later with `cmp` to `&secret_number`, Rust
infers that `&secret_number` should in fact *also* be a `u32`!

`parse` will only work on characters one would *assume* can be converted into
numbers. Thus it won't work on crazy strings with special characters, letters,
or emojis. Thus, `parse` can fail and, as a consequence, returns a `Result`
type, much like the `read_line` method does. We again handle the possibility of
an `Err` return using `expect`.

Now trying running the program!

```
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
    Finished dev [unoptimized + debuginfo] target(s) in 0.18s
     Running `target/debug/guessing_game`
Guess the number!
The secret number is: 26
Please input your guess.
13
You guessed: 13
Too small!
```

Now most of the game works. However, it seems very punishing; a player would
have to be quite lucky in order to correctly guess the generated number the
first time. We can make the game a bit more friendly by allowing players to
retry until they get it right. We can do this by adding a loop!

### Allowing Multiple Guesses with Looping

The `loop` keyboard creates an infinite loop:

```
    // --snip--

    println!("The secret number is: {secret_number}");

    loop {
        println!("Please input your guess.");

        // --snip--

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => println!("You win!"),
        }
    }
}
```

The big change here is that we've moved the game's guessing logic into a `loop {
}`. Try running the game, and you might find something pretty dystopian: the
player can never actually stop playing!

```
    Finished dev [unoptimized + debuginfo] target(s) in 0.00s
     Running `target/debug/guessing_game`
Guess the number!
The secret number is: 29
Please input your guess.
50
You guessed: 50
Too big!
Please input your guess.
25
You guessed: 25
Too small!
Please input your guess.
37
You guessed: 37
Too big!
Please input your guess.
31
You guessed: 31
Too big!
Please input your guess.
27
You guessed: 27
Too small!
Please input your guess.
29
You guessed: 29
You win!
Please input your guess.
quit
thread 'main' panicked at 'Please type a number!: ParseIntError { kind: InvalidDigit }', src/main.rs:21:47
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```

Typing `quit` does indeed quit the game, but you'll also find that entering any
non-number input also quits. Quitting by panicing is what power users may call
"suboptimal" or "bad". So we should fix this!

#### Quitting after a Correct Guess

We can program the game to exit cleanly when the player wins the game by
guessing the correct number. We do this by adding a `break` statement to
`guessing_game/src/main.rs`:

```
        // --snip--

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
```

Adding `break` after "You win!" makes the program exit our infinite `loop` when
the correct number is guessed. Exiting the `loop` here means that we also exit our
program, as it is the last bit of code in `main`.

Note that this only handles one part of the problem we identified earlier; *any*
non parseable entry by the player crashes the game. We should fix this so that
players can have a good time even when violating the rules of the game.

### Handling Invalid Input

We can resolve the problem of disingenuous players by making the game ignore
nonnumeric input:

```
        // --snip--

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {guess}");

        // --snip--
```

We have made our code a bit more robust here. In shadowing `guess` we introduce
a new `match`. This `match` will replace the behavior of the `expect` method we
were previously using; now, instead of crashing when an error is encountered, we
now cleanly handle the error. Because `parse` returns a `Result` type, which is
an enum, we can use our knowledge of the variants available for `Result` to form
the arms of our `match`.

If `parse` succeeds, we just do what we would have done with `expect`: return
the number. If `parse` fails, it will return an `Err` value containing
information about the error. Note the `_`. This is a catchall value. Here, we
are saying we want to match *all* `Err` values, no matter what that `Err` value
is. This coursely handles errors; we could make it more fine-grained if we
wanted to make error messages more explicit, but we don't need to here. The
`continue` informs the program that it should move to *the next iteration* of
the `loop` we are in, essentially giving the user a do-over!

Go ahead and test the new code to make sure it works as expected and have some
fun:

```
   Compiling guessing_game v0.1.0 (/mnt/share/git/programming/rust/ch2/guessing_game)
    Finished dev [unoptimized + debuginfo] target(s) in 0.17s
     Running `target/debug/guessing_game`
Guess the number!
The secret number is: 77
Please input your guess.
50
You guessed: 50
Too small!
Please input your guess.
75
You guessed: 75
Too small!
Please input your guess.
87
You guessed: 87
Too big!
Please input your guess.
81
You guessed: 81
Too big!
Please input your guess.
78
You guessed: 78
Too big!
Please input your guess.
77
You guessed: 77
You win!
```

Hooray!

Of course, the game still prints the secret number, so we should fix that. You
know, to make it fun.

The final program we end up with is:

```
use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed: {guess}");

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
```

### Summary

We learned a few new concepts here: `let`, `match`, functions, methods, external
crates, typing, and more. All of these will be discussed and demonstrated in
more detail later. [Chapter 3](../ch3) covers the usual concepts you might
expect: variables, data types, and functions, as well as their usage in Rust.
[Chapter 5](../ch5) will discuss structs and method syntax, and [Chapter
6](../ch6) explains how enums work.
