## Comments

When your code isn't self-documenting (shame on you), you should add comments to
help your poor unfortunate future readers.

You can do this by leveraging comments, which you've seen in these notes:

```
// hello, world
```

In Rust, the idiomatic comment style begins with two slashes and the comment
continues until the end of the line. For comments extending beyond a single
line, simply use more double slashes:

```
// So we’re doing something complicated here, long enough that we need
// multiple lines of comments to do it! Whew! Hopefully, this comment will
// explain what’s going on.
```

Comments can also happen in-line:

```
fn main() {
    let lucky_number = 7; // I’m feeling lucky today
}
```

But more often, they will exist on their own line:

```
fn main() {
    // I’m feeling lucky today
    let lucky_number = 7;
}
```

Rust also has another sort of comment, a documentation comment, which will be
discussed in [Chapter 14](../../ch14).
