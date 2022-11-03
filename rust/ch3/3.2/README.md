## Data Types

All values in Rust have a certain *data type*. Here, we will look at two data
type subsets: scalar and compound.

Remember that because Rust is a statically typed language, all variable types
must be known at compile time. Usually the type can be inferred by the compiler,
but in instances where there are many possible interpretations of a type (for
instance, the numerous kinds of numerical types available), we have to add a
type annotation, as in

```
let guess: u32 = "42".parse().expect("Not a number!");
```

If we don't add the `: u32` type annotation here, Rust displays the following:

```
error[E0282]: type annotations needed
 --> src/main.rs:2:9
  |
2 |     let guess = "42".parse().expect("Not a number!");
  |         ^^^^^ consider giving `guess` a type

For more information about this error, try `rustc --explain E0282`.
error: could not compile `no_type_annotations` due to previous error
```

### Scalar Types

A scalar type represents a single value. There are four primary scalar types in
Rust: integers, floating-points, booleans, and characters.

#### Integer Types

Here is a table of the built-in integer types in Rust:

|Length |Signed |Unsigned|
|-------|-------|--------|
|8-bit  |`i8`   |`u8`    |
|16-bit |`i16`  |`u16`   |
|32-bit |`i32`  |`u32`   |
|64-bit |`i64`  |`u64`   |
|128-bit|`i128` |`u128`  |
|arch   |`isize`|`usize` |

Key here are two things:

1) Signed integers can be negative (hence the "signed")

2) The numerical value of the integer is how many bits the integer can occupy

Signed and unsigned integers of the same size can store the same number of
numbers, though the signed numbers cover [-(2^(n-1), 2^n-1) whereas the unsigned
integers cover [0, 2^n). Thus, a `u8` can store numbers [0, 256), and an `i8`
can store [-128,128).

In the case of `{i,u}size`, the size depends on the architecture at runtime. So
for i386, it's 32 bits; for x86_64, it's 64 bits.

You can also represent integer literals of various forms. Usefully, `_` can be
used as a delimiter for literals to make their readability easier without
changing their value in the code.

|Number literal  |Example    |
|----------------|-----------|
|Decimal         |98_222     |
|Hex             |0xff       |
|Octal           |0o77       |
|Binary          |0b1111_0000|
|Byte (`u8` only)|b'A'       |

By design, Rust defaults the type of an integer if left unspecified to `i32`.
This is a sane default and is usually a good choice. The main concern is with
`{i,u}size`, and these would be used when indexing some sort of collection.

Here is a useful tidbit about integer overflow:

> Let’s say you have a variable of type `u8` that can hold values between 0
> and 255. If you try to change the variable to a value outside of that range,
> such as 256, integer overflow will occur, which can result in one of two
> behaviors. When you’re compiling in debug mode, Rust includes checks for
> integer overflow that cause your program to panic at runtime if this
> behavior occurs. Rust uses the term panicking when a program exits with
> an error; we’ll discuss panics in more depth in the “Unrecoverable Errors
> with `panic!`” section in [Chapter 9](../../ch9).

> When you’re compiling in release mode with the `--release flag`, Rust does
> not include checks for integer overflow that cause panics. Instead, if
> overflow occurs, Rust performs two’s complement wrapping. In short, values
> greater than the maximum value the type can hold “wrap around” to the
> minimum of the values the type can hold. In the case of a `u8`, the value 256
> becomes 0, the value 257 becomes 1, and so on. The program won’t panic, but
> the variable will have a value that probably isn’t what you were expecting
> it to have. Relying on integer overflow’s wrapping behavior is considered
> an error.
>
> To explicitly handle the possibility of overflow, you can use these families
> of methods provided by the standard library for primitive numeric types:
>
> Wrap in all modes with the `wrapping_*` methods, such as `wrapping_add`
> Return the `None` value if there is overflow with the `checked_*` methods
> Return the value and a boolean indicating whether there was overflow with the
> `overflowing_*` methods
> Saturate at the value’s minimum or maximum values with `saturating_*` methods

#### Floating-Point Types

Rust also includes two primitive types for floating-point numbers, `f32` and
`f64`. The default in Rust is `f64`, because in modern systems it is roughly the
same speec as `f32`, but offers greater precision.

An example which shows floating points in action:

```
fn main() {
    let x = 2.0;      // f64
    let y: f32 = 3.0; // f32
}
```

Floating-points are represented in Rust in accordance with IEEE-754. `f32` is a
single-precision float; `f64` is double-precision.

#### Numeric Operations

Standard numerical operations are all supported by Rust. Integer division always
rounds down.

```
fn main() {
    let sum = 5 + 10             // addition
    let difference = 95.5 - 4.3; // subtraction
    let product = 4 * 30;        // multiplication
    let quotient = 56.7 / 32.2;  // division
    let floored = 2 / 3;         // integer division (0)
    let remainder = 43 % 5       // modular; get remainder
}
```

#### The Boolean Type

Booleans in Rust, as in other languages, have two possible values: `true` and
`false`. Booleans are a single byte in size. In Rust, the type is specified by
`bool`. For example,

```
fn main() {
    let t = true;
    let f: bool = false; // explicit type annotation
}
```

Booleans are primarily used in things such as conditionals. See [section
3.5](../3.5/README.md) for more.

#### The Character Type

The `char` type is Rust's most primitive alphabetic type. By example,

```
fn main() {
    let c = 'z'
    let z: char = 'ℤ'; // explicit type annotation
    let heart_eyed_cat = '😻';
}
```

Note how `char` literals are specified with single quotes. This is opposed to
string literals, which use double quotes. A small amount of pedantry is to be
expected in any language.

The `char` type is four bytes in size and represents a Unice Scalar Value, which
extends their possible representations beyond just ASCII. Everything from
characters in different languages, accents, even zero-width spaces, are valid
`char` values in Rust. Unicode Scalar Values range from `U+0000` to `U+D7FF` and
`U+E0000` to `U+10FFFF`. Of course, Unicode is a mess, and so the intuition of
what a `char` is doesn't usually align with what a `char` *actually* is. Have
fun.

### Compound Types

In contrast to scalar types, compound types can group multiple values into one
type. Rust has two primitive compound types: tuples and arrays.

#### The Tuple Type

You can think of tuples as being like vectors in mathematics. They are a way of
grouping together a number of values with a variety of types into a single
compound type.  Importantly (and much like mathematical vectors), tuples have a
fixed length. Once declared, they cannot grow or shrink.

Tuples are created by writing a comma-separated list of values in parentheses
(shock). Each position within a tuple has a type, and the each position's type
need not be the same as any other. The type annotation is made explicit below
for the saliency of the example:

```
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);
}
```

To get the individual values out of this tuple, we can use pattern matching to
tear it apart, as such:

```
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let (x, y, z) = tup;

    println!("The value of y is: {y}");
}
```

First, a tuple is created and bound to the variable `tup`. Then, a mattern is
used with `let` to take `tup` and turn it into three *separate variables* `x`,
`y`, and `z`. This is referred to in Rust as *destructuring*, because it breaks
up a tuple into multiple parts. Finally, we use pattern matching with the
`println` macro to extract the value of `y`, which is `6.4`.

The elements of a tuple can also be extracted directly using the `.` notation:

```
fn main() {
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let five_hundred = x.0;
    let six_point_four = x.1;
    let one = x.2;
}
```

Of course, as is tradition, our indices with tuples start at zero.

An empty tuple is called the unit tuple. The value and its corresponding type
are both written `()` and it represents an empty value or an empty return type.
Expressions will implictly return the unit value if they don't return any other
value.

#### The Array Type

An alternative to tuples is the array. In contrast to tuples, the values in an
array must all be of the same type. In Rust, much like tuples, arrays have a
fixed length.

Arrays are instantiated with comma separated lists in square brackets:

```
fn main() {
    let a = [1, 2, 3, 4, 5];
}
```

Arrays are useful in cases where you want data allocated on the stack instead of
[the heap](../../ch4), or when you want to ensure you always have a fixed number of
elements. Arrays aren't as flexible as the vector type, however. A vector is a
similar collection type provided by the standard library that *is* allowed to
grow and shrink in size (so don't confuse vectors and tuples!). If you can't
decide between an array and a vector, you should probably use a vector. [Chapter
8](../../ch8) discusses vectors in more detail.

Arrays are useful when you know the number of elements will not need to change.
For instance, if you wanted to refer to the months in a year:

```
let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];
```

You can also declare an array like so:

```
let a: [i32; 5] = [1, 2, 3, 4, 5];
```

This declares an array with five elements where those elements are `i32`
integers. Because types can be inferred by the compiler, you can skip the type
declaration and instead set a value for every element of the array by doing:

```
let a: [3; 5];
```

Which will declare a five element array where each element is `3`, an `i32`.

#### Accessing Array Elements

You can access the elements of an array like you can with a tuple by using
indexing:

```
fn main() {
    let a = [1, 2, 3, 4, 5];

    let first = a[0];
    let second = a[1];
}
```

Thus, `first` is `1`, `second` is `2`, and continue as desired.

#### Invalid Array Element Access

As arrays have a fixed size, you can't access something like `a[6]` when `a` is
only five elements long. But what happens if you try to access an element of an
array that isn't actually in that array?

```
use std::io;

fn main() {
    let a = [1, 2, 3, 4, 5];

    println!("Please enter an array index.");

    let mut index = String::new();

    io::stdin()
        .read_line(&mut index)
        .expect("Failed to read line");

    let index: usize = index
        .trim()
        .parse()
        .expect("Index entered was not a number");

    let element = a[index];

    println!("The value of the element at index {index} is: {element}");
}
```

Thsi code will successfully compile -- the problem is the user. If you run the
code with `cargo run` and enter 0, 1, 2, 3, or 4, the program will correctly
operate and print the corresponding value contained in the array `a`. If instead
you pass a number which extends beyond the length of this array (like 10),
the disingenuous user will be greeted with:

```
thread 'main' panicked at 'index out of bounds: the len is 5 but the index is 10', src/main.rs:19:19
note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
```

This results in a runtime error. Rust will check that the index specified is
less than the array length. Because the index is greater than or equal to the
length (remember, index starts at 0), Rust will panic.

This is an example of Rust's *memory safety*. In other, nonmemory safe
languages, providing a value like `10` to the program may actually result in
accessing invalid memory regions. Rust protects you and users from this sort of
invalid action by immediately exiting instead of fucking around.

[Chapter 9](../../ch9) discusses more on error handling in Rust and how to write
readable, safe code that won't e.g. panic, or allow invalid memory access.
