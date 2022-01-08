### Semantically...

Short declarations can precede if-statements:
```x := 4; if x < 0 { ... }```

Conditions can be pretty wild:
```
if condition1 {
    ...
} else {
    ...
}

if condition 2 {
    ...
}
Func(foo)
```

Switch statements might be more pragmatic:
```
switch someVariable {
    case foo:
        return bar
    case baz:
        return foobar
}
```

Switches might include some if-style logic:
```
switch {
    case foo < baz:
        return "uhoh"
    case foo > baz:
        return "phew"
}
```

Switches can also include a `default:` case which will be what happens in case
no other case holds.


Conditionals are great for things like resolving errors! For instance,

```
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

In this example, `f`, `d`, and `err` are declared. Then, immediately, `err` is checked. If
`err` isn't nil (its zero value), then we just immediately are done -- we've
failed, and you should know about it.

Note that `err` is merely _redefined_ in the `d, err ...` short declaration. It
isn't being declared again; instead, it is being _redefined_. Thus, it's legal.

```
More on this:
In a := declaration a variable v may appear even if it has already been declared, provided:

    1) this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
    2) the corresponding value in the initialization is assignable to v, and
    3) there is at least one other variable that is created by the declaration.
```


[Valuable](https://go.dev/doc/effective_go#if)
