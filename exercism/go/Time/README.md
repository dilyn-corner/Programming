There's something worth talking about for this section specifically that super
tripped me up: methods.

Methods are odd. In a way, they're like functions. EXCEPT, they're like
functions which have a *special* argument. That argument is referred to as a
receiver. You can recognize them syntactically in the docs as:

```func(receiver receiverType) funcName(arg argType) returnType {}```

But what is a receiver? A receiver is essentially just a type that can directly
call the function. For instance, structs can have methods.

Within code, you can recognize methods as:

```
package main

import (
    "fmt"
)

type Bird struct {
    name string
}

// declare method
func(b Bird) Fly() {
    fmt.Println(b.name, "is flying...")
}

func main() {
    b := Bird{"Raven"}

    // call method
    b.Fly()         // Raven is flying...
}
```

A good instance where the receiver *isn't* a struct is specifically when we're
talking about the Time package.

The most essential thing to remember (in at least this example, and if you are
completely new to _methods_), is that they are called differently than functions
syntactically -- instead of `Pkg.Func(arg)`, it's `value.Method(arg)`, where
value is something you've already defined somewhere; in the case about, the
value is the variable `b`.

According to Golang's documentation, "receivers allow us to write function calls
in an OOP manner. That means whenever an object of some type is created that
type can call the function from itself".

In summary, I hate this; but it's probably useful.
