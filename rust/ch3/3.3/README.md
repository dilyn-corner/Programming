## Functions

Functions are all over the place in Rust code. The `main` function is a special
function, but you'll probably leverage many, many functions in your coding
activities. Key here is the `fn` keyword used in every new function declaration.

Rust code, as a matter of style, users snake_case for function and variable
names. For example:

```
fn main() {
    println!("Hello, world!");

    another_function();
}

fn another_function() {
    println!("Another function.");
}
```

Note that the order of function declaration does not matter to Rust; as long as
the functions are in the same scope, they can be freely called by functions
which appear earlier or later in a code base.

Let's explore this further by creating a new project named functions. Add the
`another_function` code block from above to the boilerplate created by Cargo and
run the code with `cargo run`. You should see the below output:

```
   Compiling functions v0.1.0
(/mnt/share/git/programming/rust/ch3/3.3/functions)
    Finished dev [unoptimized + debuginfo] target(s) in 0.11s
     Running `target/debug/functions`
Hello, world!
Another function.
```

As you can see, the code is executed in the order as they appear in `main`. Feel
free to play around to convince yourself; move the function definition or alter
the order in `main`.

### Parameters

Functions can take *parameters*, special variables that are part of a functions
"signature". You can provide concrete values for those parameters.
Mathematicians might be inclined to call these things *arguments*, but what do
they even know about functions. Pft.

We can extend our example by adding a parameter:

```
fn main() {
    another_function(5);
}

fn another_function(x: i32) {
    println!("The value of x is: {x}");
}
```

If you run this program, you will get the following output:

```
    Finished dev [unoptimized + debuginfo] target(s) in 1.21s
     Running `target/debug/functions`
The value of x is: 5
```

`another_function` takes a single parameter `x`. This paramter is specified as
having a type `i32`. By passing `5`, a valid argument for `another_function`,
the `println!` macro places `5` in the place you would expect.

In function signatures, you *must* declare the type of each parameter. This is a
deliberate design decision. When specifying multiple parameters, delimit them
using a comma:

```
fn main() {
    print_labeled_measurement(5, 'h');
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}
```

In this example, our function `print_labeled_measurement` takes two arguments.
The first is named `value` and is a signed 32-bit integer `i32`. The second is
named `unit_label` and is a `char`. The function them simply prints these
values. Try running this code:

```
    Finished dev [unoptimized + debuginfo] target(s) in 0.31s
     Running `target/debug/functions`
The measurement is: 5h
```

The function prints what was expected.

### Statements and Expressions

Functions are usually made up of a series of statements and ending in an
expression. So far, we haven't included examples of expressions, but we've seen
expressions which are part of statements. Rust is an expression-based language,
so this is an important distinction to draw.

*Statements* are instructions that perform some action and do not return a
value. *Expressions* evaluate to a resulting value. For example...

We've used statements as expressions before. Creating a variable and assigning a
value to it with the `let` keyword is a statement e.g., `let y = 6;` is a
statement.

```
fn main() {
    let y = 6;
}
```

Function definitions are also statements; the entire preceding example is in
itself a statement.

Statements don't return values. Therefore, you can't assign a `let` statement to
another variable. For instance, the below code returns the following error:

```
fn main() {
    let x = (let y = 6);
}
```

```
error: expected expression, found statement (`let`)
 --> src/main.rs:2:14
  |
2 |     let x = (let y = 6);
  |              ^^^^^^^^^
  |
  = note: variable declaration using `let` is a statement

error[E0658]: `let` expressions in this position are unstable
 --> src/main.rs:2:14
  |
2 |     let x = (let y = 6);
  |              ^^^^^^^^^
  |
  = note: see issue #53667 <https://github.com/rust-lang/rust/issues/53667> for more information

warning: unnecessary parentheses around assigned value
 --> src/main.rs:2:13
  |
2 |     let x = (let y = 6);
  |             ^         ^
  |
  = note: `#[warn(unused_parens)]` on by default
help: remove these parentheses
  |
2 -     let x = (let y = 6);
2 +     let x = let y = 6;
  |

For more information about this error, try `rustc --explain E0658`.
warning: `functions` (bin "functions") generated 1 warning
error: could not compile `functions` due to 2 previous errors; 1 warning emitted
```

The specified statement `let y = 6` doesn't actually *return a value*, so there
is nothing for `x` to bind to. In languages like C and ruby, one would be able
to write something like `x = y = 6` and assign `6` to both `x` and `y`. Not so
in Rust.

The *statement* `let y = 6` contains an *expression* which evaluates to `6`
(namely, the "`6`"). Calling a function is an expression, calling a macro is an
expression, a new scope block created with curly brackets is an expression.

```
fn main() {
    let y = {
        let x = 3;
        x + 1
    };

    println!("The value of y is: {y}");
}
```

The expression

```
{
    let x = 3;
    x + 1
}
```

is a block that evaluates to `4`. That value is then bound to `y` as part of the
`let` statement. Of note is the fact that `x + 1` does not have a semi-colon at
the end; expressions do not have semicolon endings. If you add a semicolon to
the end of an expression, it becomes a statement, and thus not return a value.
Keep that in mind (this seems damningly dangerous).

### Functions with Return Values

Functions can return values to code which calls it! You don't have to name
return values, but you do have to declare their type. We do this syntactically
after a `->`. In Rust, the return value of the function is synonymous with the
value of the final expression in the block of the body of a function. You can
return *early* from a function by using the `return` keyword and specifying a
value, but most functions simply process and then return the last expression
*implicitly*. By example:

```
fn five() -> i32 {
    5
}

fn main() {
    let x = five();

    println!("The value of x is: {x}");
}
```

There are no function calls, there are no macros, nor are there `let` statements
in the `five` function. It's simply `5`. This is a perfectly valid function in
Rust. Note that `five`'s return value is specified: `-> i32`.

The line `let x = five();` means that `x` is a variable initialized with the
return value of `five()`, which is `5`. This is identical to `let x = 5;`.

Importantly, the function `five()` takes no arguments and only specifies a
return value, and the (lonely) body consists in `5`, with no semicolon. This is
because it is an *expression* whose value *we want to return*.

Another example:

```
fn main() {
    let x = plus_one(5);

    println!("The value of x is: {x}");
}

fn plus_one(x: i32) -> i32 {
    x + 1
}
```

Running this code results in `The value of x is: 6`. If we placed a semicolon at
the end of the line `x + 1`, we will get an error. This is because we changed
the line from being an *expression* to being a *statement*:

```
fn main() {
    let x = plus_one(5);

    println!("The value of x is: {x}");
}

fn plus_one(x: i32) -> i32 {
    x + 1;
}
```

```
error[E0308]: mismatched types
 --> src/main.rs:7:24
  |
7 | fn plus_one(x: i32) -> i32 {
  |    --------            ^^^ expected `i32`, found `()`
  |    |
  |    implicitly returns `()` as its body has no tail or `return` expression
8 |     x + 1;
  |          - help: remove this semicolon

For more information about this error, try `rustc --explain E0308`.
error: could not compile `functions` due to previous error
```

Note in the error message the `()` and recall what we said in the earlier
section about tuples. `()` is the unit type. This means that nothing is
returned, but we anticipate a return (because of the `-> i32`). Rust provides a
message suggesting that we remove the semicolon.
