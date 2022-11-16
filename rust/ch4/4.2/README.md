## References and Borrowing

Look at the [example/src/main.rs](code snippet from the end of last time):

```
fn main() {
    let s1 = String::from("hello");

    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() returns the length of a String

    (s, length)
}
```

What's wrong here?

Specifically it's the tuple we create. We have to return the `String` to the
calling function so that we can still use the `String` after the call to
`calculate_length`, because the `String` was moved into `calculate_length`.
Instead of doing this, we could pass a reference *to* the `String` value. A
*reference* is like a pointer in that it's an address we can follow to access
the data stored at that address. Unlike a pointer, a reference is guaranteed to
point to a valid value of a particular type for the life of that reference.

Here's how we would rewrite the above [borrowing/src/main.rs](using a reference)
instead of changing ownership:

```
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
```

Notice how we don't even declare a tuple in this new code. We instead pass a
reference to `s1`, `&s1`, into `calculate_length`, and in its definition we take
`&String` rather than `String`. the `&` represents a *reference*, and they allow
you to refer to something without taking ownership of it.

**NOTE**: The opposite of referencing is derefencing. While referencing is done
with `&`, dereferencing is done with `*`. See (Chapter 8)[../../ch8] and
(Chapter 15)[../../ch15] for more information on dereferencing.

The `&s1` syntax lets us create a reference that *refers to* the value of `s1`
but doesn't own it. Because it doesn't own it,the value it points to will not be
dropped when the reference stops being used (goes out of scope). Because we are
passing a reference `&s` to the function `calculate_length()`, we have to let
the function know that it will be accepting a reference; we do this with `s:
&String`.

In short, the scope in which the variable `s` is valid is the same as any
function parameter's scope, but the value pointed to by the reference is not
dropped when `s` stops being used because `s` doesn't have ownership. When
functions have references as parameters instead of the actual values, we won't
need to return the values in order to give back ownership -- we never had
ownership in the first place.

The action of creating a reference is called *borrowing*. Importantly, a
[borrowing_wrongly/src/main.rs](borrowed value cannot be changed):

```
fn main() {
    let s = String::from("hello");

    change(&s);
}

fn change(some_string: &String) {
    some_string.push_str(", world");
}
```

If we attempt to compile this code, we get:

```
   Compiling borrowing_wrongly v0.1.0
(/mnt/share/git/programming/rust/ch4/4.2/borrowing_wrongly) error[E0596]: cannot borrow `*some_string` as mutable, as it is behind a `&` reference
 --> src/main.rs:8:9
  |
7 | fn change(some_string: &String) {
  |                        ------- help: consider changing this to be a mutable reference: `&mut String`
8 |         some_string.push_str(", world");
  |         ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ `some_string` is a `&` reference, so the data it refers to cannot be borrowed as mutable

For more information about this error, try `rustc --explain E0596`.
error: could not compile `borrowing_wrongly` due to previous error
```

In much the same way variables are immutable by default, so too are references.

### Mutable References

We can fix this code by allowing ourselves to mdoify a borrowed value with a few
small tweaks, ultimately resulting in a [borrowing_mutably/src/main.rs](*mutable* reference):

```
fn main() {
    let mut s = String::from("hello");

    change(&mut s);
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```

The first thing we do here is change `s` to be `mut`. From there, we create a
mutable reference with `&mut s`, passing this mutable reference to the `change`
function, and updating that function's operands to be on a mutable reference
with `&mut String`.

Mutable references have one big restriction: if you have a mutable reference to
a value, you can have no other references to that value:

```
    let mut s = String::from("hello");

    let r1 = &mut s;
    let r2 = &mut s;

    println!("{}, {},", r1, r2);
```

Building this results in the error:

```
error[E0499]: cannot borrow `s` as mutable more than once at a time
 --> src/main.rs:5:14
  |
4 |     let r1 = &mut s;
  |              ------ first mutable borrow occurs here
5 |     let r2 = &mut s;
  |              ^^^^^^ second mutable borrow occurs here
6 |
7 |     println!("{}, {}", r1, r2);
  |                        -- first borrow later used here

For more information about this error, try `rustc --explain E0499`.
error: could not compile `ownership` due to previous error
```

The error says it right there: we cannot borrow `s` multiple times because it is
mutable. The borrowing will last until the `println!`.

A big boon for this is avoiding data races at compile time. A data race is
similar to a race condition and happens with these three behaviors:

    1) two or more pointers access the same data at the same time,
    2) at least one of the pointers is being used to write to the data, and
    3) there is no mechanism being used to synchronize access to the data.

Data races cause undefined behavior and can be difficult to diagnose, so the
obvious choice is to exclude such possibilities as a design philosophy.

As always, you could nest your calls in `{ ... }` to create a new scope. Doing
this will allow, say, `r1` to fall out of scope and allow `r2` to use `&mut s`
as it has become available when `r1` leaves. So you *can* use a mutable
reference multiple times, just not simultaneously.

You also cannot borrow an immutable reference as a mutable one:

```
    let mut s = String::from("hello");

    let r1 = &s     // no problem
    let r2 = &s     // no problem
    let r3 = &mut s // BIG problem

    println!("{}, {}, and {}", r1, r2, r3);
```

This returns the below error:

```
error[E0502]: cannot borrow `s` as mutable because it is also borrowed as immutable
 --> src/main.rs:6:14
  |
4 |     let r1 = &s;     // no problem
  |              -- immutable borrow occurs here
5 |     let r2 = &s;     // no problem
6 |     let r3 = &mut s; // BIG PROBLEM
  |              ^^^^^^ mutable borrow occurs here
7 |
8 |     println!("{}, {}, and {}", r1, r2, r3);
  |                                -- immutable borrow later used here

For more information about this error, try `rustc --explain E0502`.
error: could not compile `ownership` due to previous error
```

This means that we can have multiple immutable references, but we can't mix
immutable and mutable references to the same object. Nice looking out, `rustc`.

Remember: a reference's scope starts from where it is introduced and continues
through the last time that reference is used. So we could do the following:

```
    let mut s = String::from("hello");

    let r1 = &s; // no problem
    let r2 = &s; // no problem
    println!("{} and {}", r1, r2);
    // variables r1 and r2 will not be used after this point

    let r3 = &mut s; // no problem
    println!("{}", r3);
```

And we would have no issues. The scope of `r1` and `r2` ends after the first
`println!` because that is where they are last used, and they are used before
the mutable reference `r3` is created. The scopes here do not overlap, and so
the code is fine. If you want to learn more about this super cool shit, it's
called Non-Lexical Lifetimes (or NLL). You can read more about it
[here](https://blog.rust-lang.org/2018/12/06/Rust-1.31-and-rust-2018.html#non-lexical-lifetimes).

### Dangling References

In some languages, you can erroneously end up with *dangling pointers*. These
are pointers that reference a location in memory that may have been given to
someone else. This can be done by, for instance, freeing some memory while
preserving a pointer to that memory. In Rust, the compiler guarantees that no
references will dangle. We can [dangling/src/main.rs](demonstrate this):
```
fn main() {
    let reference_to_nothing = dangle();
}

fn dangle() -> &String {
    let s = String::from("hello");

    &s
}
```

When we compile this, we are greeted by:

```
   Compiling dangling v0.1.0 (/mnt/share/git/programming/rust/ch4/4.2/dangling)
error[E0106]: missing lifetime specifier
 --> src/main.rs:5:16
  |
5 | fn dangle() -> &String {
  |                ^ expected named lifetime parameter
  |
  = help: this function's return type contains a borrowed value, but there is no value for it to be borrowed from
help: consider using the `'static` lifetime
  |
5 | fn dangle() -> &'static String {
  |                 +++++++

For more information about this error, try `rustc --explain E0106`.
error: could not compile `dangling` due to previous error
```

This touches on something we haven't yet discussed: lifetimes. Lifetimes will be
dealt with in [../../ch10](Chapter 10), but if you ignore that part of the
error, you can see the real reason why this code is problematic:

```
this function's return type contains a borrowed value, but there is no value for it to be borrowed from
```

In this case, our reference to `s` goes out of scope in `dangle`, but `dangle`
is attempting to return a reference to it (`-> &String`). This means that the
reference would be pointing to an invalid `String`, and so Rust stops us.

The solution is to simply return a `String` directly instead of a reference:

```
fn no_dangle() -> String {
    let s = String::from("hello");

    s
}
```

This works fine; ownership is moved out, and nothing is deallocated.
