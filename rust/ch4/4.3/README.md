## The Slice Type

Slices let you reference a contiguous sequence of elements in a collection, as
opposed to referencing the entire collection. A slice is a *kind* of reference,
so it does not have ownership.

Here's a fun small problem: write a function that takes a string of words
separated by spaces and return the first word it finds in that string. If the
function doesn't find a space in the string, the whole string must be one word,
so the entire string should be returned.

Let's solve this problem using slices!

But first, let's try to get at the essence of the solution without appealing to
slices:

```
fn first_word(s: &String) -> ?
```

The `first_word` function has a `&String` as a parameter. We don't want
ownership, so this is fine. But what should the return be? We don't have a way
of talking about a *part* of a string. But, we could return the index of the end
of the word, indicated by a space...

```
fn first_word(s: &String) -> usize {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return i;
        }
    }

    s.len()
}
```

Because we want to parse the elements of the `String` to identify the space, we
convert the `String` to an array of bytes using the `as_bytes` method. Next, we
create an iterator over the array of bytes using the `iter` method. (Iterators
will be discussed more in [../../ch13](Chapter 13)). For now, just know that
`iter` is a method that returns each element in a collection and that
`enumerate` wraps the result of `iter` and returns each element as part of a
tuple instead. The first element of the tuple returned from `enumerate` is the
index, and the second element is a reference to the element.

Because `enumerate` returns a tuple, we can use patters to destructure that
tuple. In the `for` loop, we specify a pattern that has `i` for the index in the
tuple and `&item` for the single byte in the tuple. Because we get a reference
to the element from `.iter().enumerate()`, we use `&` in the pattern.

Inside the loop we merely check for the byte that represents the literal byte
syntax of ' ' (`b' '`). If we find a space, the position `i` is returned,
otherwise we return the length of the string using `s.len()`.

This can get very complicated. For instance, what happens when the string `s`
gets cleared by `s.clear()`? The length that we identify will still be the
length of the word, but the word is now gone. Not to mention if we wanted to do
a similar thing to the second word in the string; now we would have to juggle
the index ending the first word and the index ending the second word! It's all
very gross.

But Rust yet again comes to the rescue: string slices.

### String Slices

A *string slice* is a reference to a part of a `String`, and it looks like:

```
    let s = String::("hello world");

    let hello = &s[0..5];
    let world = &s[6..11];
```

In this case, `hello` is now just a reference to a portion of the `String` `s`,
and that portion is specified in the `[0..5]`. Slices are created using a range
within brackets by specifying `[starting_index..ending_index]`, where
`starting_index` is the first position in the slice and `ending_index` is one
more than the last position in the slice. So in the case of `let world =
&s[6..11];`, `world` would be a slice that contains a pointer to the byte at
index 6 of `s` with a length value of 5.

The range syntax `..` allows us to drop the first index if it would be 0. So

```
let s = String::from("hello");

let slice = &s[0..2];
let slice = &s[..2];
```

are equivalent statements.  Likewise, if your slice includes the last byte of
the `String`, you can drop the trailing number:

```
let s = String::from("hello");

let len = s.len();

let slice = &s[0..len];
let slice = &s[0..];
```

and both `slice`s are equal.

If you drop both the starting index and the final index (`[..]`), then the whole
string is taken.

So given all of this, let's rewrite `first_word` to return a slice. The type
that signifies "string slice" is written as `&str`:

```
fn first_word(s: &String) _> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```

Here we parse the string in the same way, except when we find a `b' '` in the
string we return a string slice using the start of the string and the index of
the space as the starting and ending indices. Now when we call `first_word`, we
get back a single value that is tied to the underlying data. The value is made
up of a reference to the starting point of the slice and the number of elements
in the slice.

Returning a slice works just as well for a theoretical `second_word` function.
If we used `s.clear()` using slices, we would get a compile-time error!

#### String Literals Are Slices

Recall that we talked about string literals being stored inside the binary. Now
that we know about slices, we can properly understand string literals:

```
let s = "Hello, world!";
```

The type of `s` is `&str`. It is a slice pointing to that specific point of the
binary. This is also why string literals are immutable; `&str` is an immutable
reference!

#### String Slices as Parameters

Knowing that you can take slices of literals and `String` values leads us to one
more optimization on `first_word` we can make:

```
fn first_word(s: &str) -> &str {
```

We switched `s` from being `&String` to being `&str`. This means that if we have
a string slice, we can pass that directly. If we have a `String`, we can pass a
slice of the `String` or a reference to the `String`. This flexibility takes
advantage of something called deref coercions, a feature covered in
[../../ch15](Chapter 15).

### Other Slices

String slices are specific to strings. But there is a more general slice type
too. Consider the array:

```
let a = [1, 2, 3, 4, 5];
```

Just as we might want to refer to a part of a string, we might want to refer to
part of an array. We'd do so like this:

```
let a = [1, 2, 3, 4, 5];

let slice = &a[1..3];

assert_eq!(slice, &[2, 3]);
```

This slice as the type `&[i32]`. It works the same way as string slices do, by
storing a reference to the first element and a length.
