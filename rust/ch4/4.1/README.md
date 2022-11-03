## What Is Ownership?

Ownership is a set of rules that govern how a Rust program manages memory. Some
languages have garbage collection which regularly checks for no-longer used
memory to free. Other languages, the memory must be explicitly allocated and
freed. Rust instead manages memory through a system of ownership with a set of
rules checked at compile time. If any rules are violated, the program will not
compile.

We will work through understanding this concept of ownership with a common data
structure: strings.

But first: what are the stack and the heap?

The stack and the heap are parts of memory available to your code to use at
runtime, but they are structured in different ways.

The stack stores values in the order it gets them and removes the values in the
opposite order (FILO). Adding data is called *pushing onto the stack*, and
removing data is called *popping off the stack*. All data stored on the stack
must have a known, fixed size. Data with an unknown size at compile tiem or a
size that might change must be stored on the heap instead.

The heap is less organized than the stack (get it?). When some data is put on
the heap, the memory allocator finds an empty spot in the heap big enough, marks
it as being in use, and returns a *pointer*, an address to that location in
memory. This process is called *allocating on the heap* and is sometimes
abbreviated as *allocating*. Because the pointer is a known and fixed size, that
itself can be put on the stack. But when you wish to use the actual data the
pointer *ahem* points at, you must follow the pointer.

Pushing to the stack is faster than allocating on the heap because the allocator
never has to search for a place; the location is always "on top".

Accessing data in the heap is slower than on the stack because you have to
follow a pointer to find it.

When your code calls a function, the values passed into the function and the
function's local variables get pushed onto the stack. When the function is over,
those values get popped off the stack.

Managing the heap is a problem addressed by ownership. Once you grasp ownership,
the stack and the heap will be things you rarely have to think about and juggle.

### Ownership Rules

Let's examine the rules of ownership:

    * Each value in Rust has an *owner*

    * There can only be one owner at a time

    * When the owner goes out of scope, the value will be dropped

### Variable Scope

As a first example of ownership, we'll discuss scope. A scope is the range
within a program for which an item is valid. For instance,

```
let s = "hello";
```

The variable `s` refers to a string literal. The variable `s` is valid from the
point at which it is declared until the end of the current scope. For instance,

```
{                           // s is not valid here, it's yet to be declared
    let s = "hello";        // s is valid from this point forward

    // do some stuff with s
}                           // this scope is now over, and s is no longer valid
```

In other words, there are two important points in time here:

    * When `s` comes *into scope*, it is valid

    * It remains valid until it goes *out of scope*

### The `String` type

In order to properly illustrate ownership, we need a more complicated data type
than ones previously discussed. Those ones all have a fixed and definite size,
can easily be stored on and popped off the stack, and can easily be copied to
make new instances if another part of code needs that same value in a different
scope. But we need to examine things which would be stored on the heap, and the
`String` type is a great example.

We've already seen string literals, which are just hard-coded strings. But these
may not be useful (hint: they probably aren't useful in general). One reason is
because they're immutable. Another is that not every string can be known when
writing the code; what if we want to take user input and store it? For these
situations, Rust has a second string type, `String`. This type manages data
allocated on the heap and as such is able to store an amount of text that is
unknown to us at compile time. You can create a `String` from a string literal
using the `from` function, like so:

```
let s = String::from("hello");
```

The double colon `::` operator allows us to namespace this particular `from`
function under the `String` type rather than using some sort of name like
`string_from`. We will discuss this *method* syntax more in [Chapter
5](../../ch5), and we will talk about namespacing with modules in [Chapter
7](../../ch7).

This sort of string *can* be mutated:

```
    let mut s = String::from("hello");

    s.push_str(", world!"); // push_str() appends a literal to a String

    println!("{}", s); // This will print 'hello, world!"
```

The difference between a string literal and a `String` is in how these types
deal with memory.

### Memory and Allocation

The `String` type requires that we allocate an amount of memory on the heap to
hold the contents. This means:

    * The memory must be requested from the memory or allocator at runtime

    * We need a way of returning this memory to the allocator when we're done
      with our `String`

The first part is done by us. When calling `String::from`, the implementation
requests the memory it needs. However, the second part is different. The memory
is automatically returned once the variable that owns it goes out of scope.

When a memory goes out of scope, Rust calls a special function called `drop`,
and it is where the author of (something like) `String` can put the code to
return the memory. Rust calls `drop` automatically at the closing curly bracket.

This pattern has "a profound impact on the way Rust code is written". Let's
explore some of those situations below!

#### Ways Variables and Data Interact: Move

Multiple variables can interact with the same data in different ways in Rust.
Let's look at an example using an integer:

```
    let x = 5;
    let y = x;
```

In this example, the value of `5` is bound to `x`, and then a copy of the value
in `x` is bound to `y`. Thus, we have `x`, `y`, both bound to `5`. What happens
if we do this with `String`?

```
    let s1 = String::from("hello");
    let s2 = s1;
```

This looks very similar, so you'll be forgiven for thinking it behaves the same.

A `String` is made up of three parts:

    * a pointer to the memory holding the contents of the string

    * a length

    * a capacity

This group of data is stored *on the stack*. The pointer's contents are stored
*on the heap*. That pointer points to a an (index, value) pair, containing the
contents of the string `s1` contains; `(0, h), (1, e), (2, l), (3, l), (4, o)`.
The length of `s1` is how much memory (in bytes) the contents of the `String` is
currently using. The capacity is the total amount of bytes that the `String` has
received from the allocator. The difference between the length and the capacity
matters, but not in this context.

When we assign `s1` to `s2`, the `String` data is copied, which means that we
copy the pointer, the length, and the capacity that are on the stack. We do
*not* copy the data on the heap that the pointer refers to -- `s2` already has a
pointer to this data, after all.

The key point here for Rust is this: when a variable in Rust goes out of scope,
the `drop` function is called and the heap memory for that variable is freed.
But because the pointer is copied between `s1` and `s2`, it's possible that when
both of these variables go out of scope, the same block of memory is freed. This
is called a double free error. Freeing twice can result in memory corruption,
which is bad news bears.

To ensure memory safety, after the line `let s2 = s1`, Rust considers `s1` as
*no longer valid*. Thus, Rust does not need to free anything when `s1` goes out
of scope - it need only free `s2`. You can test this for yourself! Code as
follows won't compile:

```
    let s1 = String::from("hello");
    let s2 = s1;

    println!("{}, world!", s1);
```

```
error[E0382]: borrow of moved value: `s1`
 --> src/main.rs:5:28
  |
2 |     let s1 = String::from("hello");
  |         -- move occurs because `s1` has type `String`, which does not implement the `Copy` trait
3 |     let s2 = s1;
  |              -- value moved here
4 |
5 |     println!("{}, world!", s1);
  |                            ^^ value borrowed here after move
  |
  = note: this error originates in the macro `$crate::format_args_nl` (in Nightly builds, run with -Z macro-backtrace for more info)

For more information about this error, try `rustc --explain E0382`.
error: could not compile `ownership` due to previous error
```

In Rust, this is known as a *move* (as opposed to shallow or deep copies, if
you're familiar with those). Here, we would say that `s1` was *moved* to `s2`.
Because `s1` is no longer valid when `s2` is declared, once `s2` goes out of
scope, `s2` is safely freed, and we no longer run into the problem of double
freeing.

#### Ways Variables and Data Interact: Clone

If instead we actually wanted to create a deep copy of the heap content of a
`String`, we would use a method called `clone`:

```
    let s1 = String::from("hello");
    let s2 = s1.clone();

    println!("s1 = {}, s2 = {}", s1, s2);
```

This works exactly as the naive Rust programmer may expect, and compiles without
issue. In this instance, the heap data IS copied. Thus, when you see a call to
`clone`, you know that some arbitrary code is being executed and that code may
be expensive. It's an indicator that something different is going on.

#### Stack-Only Data: Copy

There's one more wrinkle we haven't discussed. The below code uses integers
works and is valid:

```
    let x = 5;
    let y = x;

    println!("x = {}, y = {}", x, y);
```

Why does this work?

The reason is because here we are talking about integers (in the case of Rust,
an `i32` -- this is, remember, the default type for integers). Because the size
of the variable in question is known at compile time, there's no good reason we
shouldn't just copy the contents of `x` into `y` while building the program, and
no good reason we wouldn't want `x` to still be valid after `y` is declared.

Rust has a special annotation called the `Copy` trait that we can place on types
that are stored on the stack, as integers are. If a type implementation gets the
`Copy` trait, variables that use it do not move, but rather are trivially
copied, making them still valid after assignment to another variable. We will
learn more about traits in [Chapter 10](../../ch10).

If any of the types we use declare the `Drop` trait, Rust will not allow us to
annotate a type with `Copy` - instead, we would get a compile time error. We
will learn in [Appendix C](../../ch21/C) how to add the `Copy` annotation to a
type to implement the trait.

What types implement the `Copy` trait? You can certainly check the documentation
(and you should!), but as a general rule, any group of simple scalar values can
implement `Copy`, and nothing that requires allocation or is some form of
resource can implement `Copy`. Here are some types implementing `Copy`:

    * All the integer types, such as `u32`

    * The boolean type, `bool` ({`true`|`false`})

    * All the floating point types, such as `f64`

    * The character type, `char`

    * Tuples, if they only contain types that also implement `Copy`

### Ownership and Functions

Passing a variable to a function works similarly to assigning a value to a
variable in this context. Passing a variable to a function will move or copy,
just as assignment does:

```
fn main() {
    let s = String::from("hello");  // s comes into scope

    takes_ownership(s);             // s's value moves into the function...
                                    // ... and so is no longer valid here

    let x = 5;                      // x comes into scope

    makes_copy(x);                  // x would move into the function,
                                    // but i32 is Copy, so it's okay to still
                                    // use x afterward

} // Here, x goes out of scope, then s. But because s's value was moved, nothing
  // special happens.

fn takes_ownership(some_string: String) { // some_string comes into scope
    println!("{}", some_string);
} // Here, some_string goes out of scope and drop is called. The backing
  // memory is freed.

fn makes_copy(some_integer: i32) { // some_integer comes into scope
    println!("{}", some_integer);
} // Here, some_integer goes out of scope. Nothing special happens.
```

If we trie dto use `s` after the call to `takes_ownership`, Rust woudl throw a
compile-time error. These static checks act to protect us from mistakes. Try
adding code to `main` that uses `s` and `x` to see where you can use them and
where the ownership rules prevent you from doing so.

## Return Values and Scope

Returning values can also transfer ownership. Below is an example:

```
fn main() {
    let s1 = gives_ownership();        // gives_ownership moves its return
                                       // value into s1

    let s2 = String::from("hello");    // s2 comes into scope

    let s3 = takes_and_gives_back(s2); // s2 is moved into takes_and_gives_back,
                                       // which also moves its return value into s3
} // Here, s3 goes out of scope and is dropped. s2 was moved, so nothing
  // happens. s1 goes out of scope and is dropped.

fn gives_ownership() -> String {             // gives_ownership will move its
                                             // return value into the function
                                             // that calls it

    let some_string = String::from("yours"); // some_string comes into scope

    some_string                              // some_string is returned and
                                             // moves out to the calling
                                             // function
}

// This function takes a String and returns one
fn takes_and_gives_back(a_string: String) -> String { // a_string comes into
                                                      // scope

    a_string // a_string is returned and moves out to the calling function
}
```

Variable ownership follows the same pattern every time: assigning a value to
another variable moves it. When a variable that includes data on the heap goes
out of scope, the value will be cleaned up by `drop` unless ownership of the
data has been moved to another variable.

But taking and returning ownership can get a bit tedious. What if we want to let
a function use a value but not take ownership?

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

But this is too much ceremony and a lot of work for a concept that should be
common. Luckily for us, Rust ahs a feature for using a value without
transferring ownership. This is called *references*. See the next section!
