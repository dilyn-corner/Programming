## Installation

The recommended way of installing Rust is through using `rustup`, a command line
tool for managing Rust versions and its associated tools. It's up to you if this
is how you want to manage Rust; it's usually packaged across many distributions
through your usual package manager.

Doing it this way will place all of the tooling in `$HOME/.cargo/bin`, so you'll
want to make sure this is added to your `$PATH` variable.

```
curl --proto '=https' --tlsv1.3 https://sh.rustup.rs -sSf | sh
export PATH=$HOME/.cargo/bin:$PATH
```

This installation process also includes Rust's documentation, which can be read
offline - `rustup doc` will open the local documentation in your web browser.

You can now use `rustup update` to update these artifacts, and even `rustup self
uninstall` to have `rustup` uninstall Rust and itself!
