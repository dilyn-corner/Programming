### The Err of our Funcs

Errors are supremely useful. So useful, in fact, that Go has a builtin type for
errors: 
```
type error interface {
    Error() string
}
```

Frequently, functions include an `err` in their return values:
```func Func(foo type) (bar type, err error) { ... }```

So, for instance...
```
f, err := func.Func(foo)
if err != nil {
    log.Fatal(err)
}
// Do something meaningful because we avoided certain disaster
```

`log.Fatal()` is a specific function which prints an error message and stops.

If you achieve some undesired condition, you can return an error in useful ways:

```
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // Some implementation details
}
```

We a priori know something horrendous would occur in real space if we allowed
Sqrt() to accept negative values. And because we know this a priori, we can
include such a case in our function's return values AND check for it. Genius.


You can find a lot more interesting details [here](https://go.dev/blog/error-handling-and-go)

Also some insights can be had [here](https://github.com/golang/go/wiki/Errors)
