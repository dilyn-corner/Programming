## Control Flow

Basic buildings blocks of code include things like only running certain blocks
if a condition is true, or running something repeatedly while a condition is
true. The most common constructs that let you *control* the *flow* (get it) or
Rust code are `if` expressions and loops.

### `if` Expressions

An `if` expression allows you to branch your code depending on conditions. You
provide a condition and then state, "if this condition is met, run this block of
code. If the condition is not met, do not run this block of code".

Create a new project named branches, with `branches/src/main.rs`:

```
fn main() {
    let number = 3;

    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }
}
```

Blocks of code associated with conditions in `if` expressions are sometimes
called *arms*, like in `match` expressions.

`else` statements are optional for `if` statements, and we opt to include one in
the above example. If you don't provide an `else`, the entire `if` block is
simply skipped. Running the above code:

```
   Compiling branches v0.1.0 (/mnt/share/git/programming/rust/ch3/3.5/branches)
    Finished dev [unoptimized + debuginfo] target(s) in 0.12s
     Running `target/debug/branches`
condition was true
```

If you changed the value of `number` to something like `7`, you would get the
following output:

```
    Finished dev [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
condition was false
```

This happens because the condition in our `if` instantiation is now `false`, and
so we evaluate our `else` block.

The condition in an `if...else` *must* be a bool (true|false). If the condition
is *not* a bool, we get an error. For instance, if we simply do `if number {`
instead of passing a condition on `number`:

```
error[E0308]: mismatched types
 --> src/main.rs:4:8
  |
4 |     if number {
  |        ^^^^^^ expected `bool`, found integer

For more information about this error, try `rustc --explain E0308`.
error: could not compile `branches` due to previous error
```

Rust does not convert nonbools to bools. You must always give a condition for
your `if`. if you wanted the code to run only if `number` were nonzero,

```
fn main() {
    let number = 3;

    if number != 0 {
        println!("number was something other than zero");
    }
}
```

And it will print what you expect.

#### Handling Multiple Conditions with `else if`

You can use chained `if` and `else` conditions by using an `else if`:

```
fn main() {
    let number = 6;

    if number % mod 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }
}
```

This program has four possible paths, and if you run it:

```
    Finished dev [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/branches`
number is divisible by 3
```

It will check each `if` in sequence (much like how it would check the arms of a
`match`) and if the condition holds true, that code is executed and the `if`
block is exited. Style-wise, too many `if else` blocks clutter your code. You
might want to look into using `match` instead to refactor your code base. Refer
to [Chapter 6](../../ch6) for useful information on `match`.

#### using `if` in a `let` statement

As `if` is an expression, it can be used in a `let` statement:

```
fn main() {
    let condition = true;
    let number = if condition { 5 } else { 6 };

    println!("The value of number is: {number}");
}
```

The `number` variable will be bound to a value based on the outcome of the `if`.
In running, you will see that `5` is printed.

Note that if we changed from `6`, an `i32`, to `"six"`, a `string`,

```
error[E0308]: `if` and `else` have incompatible types
 --> src/main.rs:4:44
  |
4 |     let number = if condition { 5 } else { "six" };
  |                                 -          ^^^^^ expected integer, found
`&str`
  |                                 |
  |                                 expected because of this

For more information about this error, try `rustc --explain E0308`.
error: could not compile `branches` due to previous error
```

Thus, the code in our `if` and `else` blocks must end with the same type.
Essentially, `number` cannot conditionally be an `i32` *or* a `string`, it must
be known at compile-time if it is one or the other.

### Repetition with Loops

Often, you will want to execute a block of code more than once. Rather than
copy-pasting a potentially arbitrarily unknown number of times, we can leverage
an (potentially) infinite loop. Rust supports three kinds of loops: `loop`,
`while`, and `for`.

#### Repeating Code with `loop`

`loop` tells Rust to execute a block of code over and over forever until it is
*explicitly told* to stop:

```
fn main() {
    loop {
        println!("again!");
    }
}
```

When you run this code, `again!` will be printed. Forever. Good luck.

Fortunately, Rust includes a way to escape a `loop`. You can do this by using
`break` within a `loop` to inform the program under what conditions it can exit
the loop. You can also use `continue`, and in a loop this tells the program to
skip any *remaining* code in the iteration of the loop and to go to the next
one.

#### Returning Values from Loops

A good use of a `loop` is to retry some operation which can potentially fail,
like checking whether a thread has completed its job. You could also need to
pass on the result of that operation to some code outside of the loop. To do
this, you add the value to return after the `break` expression used to stop the
`loop`:

```
fn main() {
    let mut counter = 0;

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };

    println!("The result is {result}");
}
```

#### Loop Labels to Disambiguate Between Multiple Loops

When it comes to nested loops, `break` and `continue` apply to the innermost
loop at that point. You can optionally label your loops and then use those
labels with `break` or `continue` to apply them to the labelled loop instead of
the innermost loop. Loop labels *must* begin with a single quote:

```
fn main() {
    let mut count = 0;
    'counting_up: loop {
        println!("count = {count}");
        let mut remaining = 10;

        loop {
            println!("remaining = {remaining}");
            if remaining == 9 {
                break;
            }
            if count == 2 {
                break 'counting_up;
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {count}");
}
```

The first loop is labelled `'counting_up`. The inner loop is unlabelled. The
first `break` will exit the inner loop (the unlabelled `loop`), and the second
`break` explicitly `break`s a labelled loop (the outer one).

#### Conditional Loops with `while`

Frequently, you'll want to run a `loop` only given some condition being true.
When the condition ceases to be true, you would call `break` to stop. You could
of course implement this yourself using a combination of `loop`, `if`, `else`,
and `break`. But it's so common and useful that it has its own keyword: `while`.
A `while` loop will check a condition. If it is true, the code runs. Then it
checks again. This sequence continues until the condition evaluates to false, in
which case the loop `break`s:

```
fn main() {
    let mut number = 3;

    while number != 0 {
        println!("{number}!");

        number -= 1;
    }

    println!("LIFTOFF!!!");
}
```

#### Looping Through a Collection with `for`

You could opt to use `while` to loop over the elements of some collection such
as an array. For instance:

```
fn main() {
    let a = 10, 20, 30, 40, 50];
    let mut index = 0;

    while index < 5 {
        println!("the value is: {}", a[index]);

        index += 1;
    }
}
```

Here, the code simply counts through the elements of the array `a`. Because `a`
is of length five and indices start at zero, once `index` is five the condition
`index < 5` is false and the `while` loop breaks.

However, this approach can be error prone; the program could panic if the index
value or test condition are incorrect. For instance, if we modified `a` to have
fewer elements than five, but didn't update the condition of our `loop`, we
would cause a panic. It's also slow, because the compiler will add runtime code
to perform the check on the condition.

Instead, we can use a `for` loop and execute some code for each item in the
collection:

```
fn main() {
    let a = 10, 20, 30, 40, 50];

    for element in a {
        println!("the value is: {element}");
    }
}
```

When this code is run, it will yield the same result as in the previous example.
More importantly, code safety has increased -- the loop no longer depends on any
facts about `a` other than that it exists.

It is for these reasons that `for` loops are the most commonly used loops in
Rust. Even if you just wanted to loop through something for a known number of
iterations, you would probably want to use a `for` loop. The way to do this is
to use a `Range`. `Range` is provided by the standard library, and generates all
numbers in sequence starting from *one* (sigh) and continuing to an upper-bound
(exclusively). Here's an example using a countdown (using a method `rev` we've
yet to talk about):

```
fn main() {
    for number in (1..4).rev() {
        println!("{number}!");
    }
    println!("LIFTOFF!!!");
}
```

The code is a fair bit nicer, eh?

Here are some examples to try your hand at!

    * Convert temperatures between Fahrenheit and Celsius
    * Generate the nth Fibonacci number
    * Print the lyrics to the Christmas carol “The Twelve Days of Christmas,” taking
        advantage of the repetition in the song
