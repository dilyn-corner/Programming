## Hello, World!

The canonical first program.

The standard filename for Rust source code is `.rs`. Convention for filenames
with spaces is to use an `_` instead of ` ` or `-`.

Create a file named `main.rs` with contents:

```
fn main() {
    println!("Hello, world!");
}
```

```
rustc main.rs
./main
```

You should be greeted by a `"Hello, world!"`.

Let's explore this small bit of syntax.

`fn main()` defines a function named `main`. `main`, like in other languages, is
a special function name - it is the first code run in any Rust executable. This
particular snippet indicates that `main` takes no parameters and returns
nothing. As with other languages, the body of a function is wrapped in `{ }`.

`println!("Hello, world!");` is the meat of our `main` function. What's
important to note here is that `println!` is a Rust *macro*. This is different
from what you might normally expect, that being calling the *function*
`println`. Rust macros are explored more in `../ch19/README`. The point here is
that macros may sometimes behave differently from their function counterparts,
so you'll want to make sure your familiar with those differences when using
them!

`"Hello, world!"` is a simple string passed as an argument to the `println!`
macro.

Most lines in Rust are ended with a `;` as in other languages.
