## Variables and Mutability

Variables in Rust are immutable by default. This is to "motivate you" to write
code in a way that takes advantage of safety and concurrency in Rust. However,
you can still choose to make variables mutable. Here, we will explore why
immutability is the default and why sometimes you might choose otherwise.

Create a new project called `variables` using `cargo new variables`, and replace
the code in `variables/src/main.rs` with the following. It won't compile yet.

```
fn main() {
    let x = 5;
    println!("The value of x is: {x}");
    x = 6;
    println!("The value of x is: {x}");
}
```

If you run `cargo run` you should receive the following message:
```
   Compiling variables v0.1.0 (/mnt/share/git/programming/rust/ch3/3.1/variables)
error[E0384]: cannot assign twice to immutable variable `x`
 --> src/main.rs:4:5
  |
2 |     let x = 5;
  |         -
  |         |
  |         first assignment to `x`
  |         help: consider making this binding mutable: `mut x`
3 |     println!("The value of x is: {x}");
4 |     x = 6;
  |     ^^^^^ cannot assign twice to immutable variable

For more information about this error, try `rustc --explain E0384`.
error: could not compile `variables` due to previous error
```

The error message indicates that you `cannot assign twice to immutable variable
'x'`, as we tried to assign a second value to the immutable `x` variable. It
even offers a helpful suggestion!

> It’s important that we get compile-time errors when we attempt to change a
> value that’s designated as immutable because this very situation can lead to
> bugs. If one part of our code operates on the assumption that a value will
> never change and another part of our code changes that value, it’s possible
> that the first part of the code won’t do what it was designed to do. The cause
> of this kind of bug can be difficult to track down after the fact, especially
> when the second piece of code changes the value only sometimes. The Rust
> compiler guarantees that when you state a value won’t change, it really won’t
> change, so you don’t have to keep track of it yourself. Your code is thus
> easier to reason through.

Mutability however makes code more convenient to write, and thus we can make a
variable mutable by adding `mut` before the variable name. Adding `mut` also
telegraphs to future code readers the intention of the variable: it is designed
with change in mind.

By changing the code in `variables/src/main.rs` from `let x = 5;` to `let mut x
= 5;`, we get code which compiles:

```
   Compiling variables v0.1.0 (/mnt/share/git/programming/rust/ch3/3.1/variables)
    Finished dev [unoptimized + debuginfo] target(s) in 0.11s
     Running `target/debug/variables`
The value of x is: 5
The value of x is: 6
```

### Constants

Like immutable variables, *constants* are values which are bound to a name and
not allowed to change. However, constants and variables have a few differences.

First, constants *cannot be made* immutable. Constants are declared using
`const` instead of `let`, and the type of the value *must be* annotated.

Constants can be declared in any scope, including the global scope, which means
that they can be used across many parts of code in a project.

Finally, constants must be set to a constant expression (get it?), not the
result of some computation that can only be known at runtime.

An example of a constant:

```
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
```

The constant's name here is `THREE_HOURS_IN_SECONDS` and its value is set to
what you might expect: how long three hours is in seconds. The naming convention
in Rust for constants is all uppercase with underscores separating words.
Because the compiler can compute some operations at compile-time, the constant's
value can be demonstrated in a more readily understood and verifiable manner
than just using the value 10,800.

Constants are valid for the entire duration of a program within the scope they
are declared in. Thus, constants can be useful for singular functions, or for
many different functions to refer to and use. The usage of constants means that,
if the value needed to be changed (god forbid a minute were no longer sixty
seconds), only one place in the code would need to updated instead of many.

### Shadowing

You can leverage *shadowing* to declare a new variable with the same name as a
previous one. We say that the first variable is *shadowed* by the second, which
means that the second variable is what the compiler will see when you use the
name of the variable.  A variable can be shadowed by using the same variable's
name and repeating the use of the `let` keyword as follows:

```
fn main() {
    let x = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}");
    }

    println!("The value of x is: {x}");
}
```

Doing `cargo run`, we see:

```
   Compiling variables v0.1.0 (/mnt/share/git/programming/rust/ch3/3.1/variables)
    Finished dev [unoptimized + debuginfo] target(s) in 0.11s
     Running `target/debug/variables`
The value of x in the inner scope is: 12
The value of x is: 6
```

First, the program binds `x` to a value of 5. Then, a new variable `x` is
created by repeating the declaration of `x`, taking the originally defined value
and adding 1. Now the value of `x` is 6. *Then*, an inner scope is created with
`{ }`, and a third `let` shadows the previously shadowed `x` with a new value,
doubling it to 12. When the inner scope ends, the inner shadowing ends. `x`
returns to its previously shadowed version of 6, as expected.

Shadowing is different from mutability. If we tried to reassign the value of
this variable without using the `let` keyword, we would get a compile time
error. By using `let`, we can perform a few transformations on a value but have
the variable remain immutable afterwards.

Another important difference between shadowing and mutability is that because
we're effectively creating a new variable, we can leverage shadowing to modify
the type of the variable.

By way of example, suppose we ask a user to *show* how many spaces they want
between some text by *inputting space characters*. We can shadow that variable
to store that string of spaces as a number instead of as raw characters:

```
    let spaces = "   ";
    let spaces = spaces.len();
```

The first `spaces` variable is a string type and the second `spaces` variable is
a number type! Shadowing variables allows us to reuse variable names instead of
painstakingly appending differentiators to our variables a la `spaces_str` and
`spaces_num`. __However__, if we try to use `mut` for our first declaration of
spaces, we get a compile time error:

```
error[E0308]: mismatched types
 --> src/main.rs:3:14
  |
2 |     let mut spaces = "   ";
  |                      ----- expected due to this value
3 |     spaces = spaces.len();
  |              ^^^^^^^^^^^^ expected `&str`, found `usize`

For more information about this error, try `rustc --explain E0308`.
error: could not compile `variables` due to previous error
```
