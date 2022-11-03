## Hello, Cargo!

Cargo is Rust's build system and its package manager. It's used to handle many
common tasks for you, like building code and downloading the necessary
dependencies. Cargo becomes very useful for larger projects than the simple ones
we are building right now.

Let's create a similar project as in `../1.2/README.md`, but using Cargo.

```
cargo new hello_cargo
cd hello_cargo
```

Looking at the contents of this new directly created by Cargo, we see...

```
hello_cargo/
├── Cargo.toml
└── src
    └── main.rs

1 directory, 2 files
```

Cargo does some boilerplate work for us here in the `hello_cargo/Cargo.toml` and
`hello_cargo/src/main.rs` files. It also will create a `.git` and `.gitignore`
file for us if we aren't already in a directory using `git`. You can force Cargo
to create these files by using `cargo new --vcs=git`, and you can even use
another VCS if you prefer.

Let's examine these files, starting with `hello_cargo/Cargo.toml`:

```
[package]
name = "hello_cargo"
version = "0.1.0"
edition = "2021"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
```

`[package]` defines some metadata about the project we've created. `edition`
will be spoken more of in Appendix E.

The `[dependencies]` line is where we would list any relevant dependencies for
our project. These dependencies are known as crates.

In our `hello_cargo/src/main.rs` we find:

```
fn main() {
    println!("Hello, world!");
}
```

Cargo has generated a simple "Hello, world!" program for us! Note that Cargo
expects our `.rs` files to live in the `src/` tree, leaving the top-level
directory for project metadata files like READMEs, licensing information,
configuration files, and other things not pertient for compiling the code. You
would also find a directory named `tests/` in the top-level directory if we were
making tests for our code.

Now let's use cargo to build our project! Running `cargo build`, we see:

```
   Compiling hello_cargo v0.1.0
(/mnt/share/git/programming/rust/ch1/1.3/hello_cargo)
    Finished dev [unoptimized + debuginfo] target(s) in 0.25s
```

The output of Cargo's build process places our artifacts by default in
`target/debug/hello_cargo`, and we can run it in the usual way:

```
./target/debug/hello_cargo
```

Cargo will also create a new file named `Cargo.lock` when it runs a build. This
will lock the dependencies to a particular version to ensure no upstream changes
sneak up on you.

You can also use `cargo run` to build and run the code all in one. Note that
much like something like `make`, Cargo supports something like incremental
building; because we've already built the code using `cargo build` and haven't
changed any of the `src/` files, Cargo will skip rerunning the build and simply
execute the executable!

Finally, Cargo also supports checking that the code will compile but will skip
generating an executable. You can do this sanity checking with `cargo check`.
You can use this to save time during development work, skipping potentially
lengthy builds if all you want to check is that it *does* build.

When you finish your project you may want to strip the debug symbols from your
binaries; you can do this with the `--release` flag: `cargo build --release`.
The resulting executable will be in `target/release/` instead.

As a bonus, you can use `cargo clean` to remove the artifacts Cargo creates
during the build process (though this will leave behind the `Cargo.lock` file).
