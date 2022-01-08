### Be Aware

Go provides two comment styles:

```
/* 
    Block comments
*/

// Line comments
```

Line comments are the norm in Go. Block comments are mostly reserved for package
comments.

[The reason for this](https://pkg.go.dev/cmd/go#hdr-Show_documentation_for_package_or_symbol)

`godoc` is used for document generation. It does this by essentially parsing for
comments. For instance, 
```
// FuncName does a particular action and returns a specific output
func FuncName() type { ... }
```

The comment here would actually provide documentation for the function that
immediately follows it (note the lack of a blank line), scraped by `godoc`. A
similar thing exists for block style comments which act to describe what a
package is good for.

So in short, make sure your comments are good. Odds are, someone will be reading
them at some point.


[Specifically Relevant Read](https://go.dev/doc/effective_go#commentary)
