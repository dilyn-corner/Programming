### Slice Me Not

Slices are dynamically resizable arrays. They are also much more common than
arrays in Go.

A slice is a type []T with elements of type T.
A slice can be formed by specifying a lower and upper index:
```a[low : high]```
This selects a half-open range; [low, high). By example, `a[1:4]` is a slice which
includes elements one through three of `a`.

Importantly, a slice doesn't store any data; it merely acts as a description of
a section of some underlying array. Thus, changing the elements in a slice
modifies the array it is a slice of.

A large example:
```
package main

import "fmt"

func main() {
                                 // In order to specify an initial slice size
    s := make([]string, 3)       // use the make builtin to create an empty
    fmt.Println("empty:", s)     // slice of length 3

    s[0] = "a"                   // Values can be set just like with arrays
    s[1] = "b"                   // Values can be gotten just like with arrays
    s[2] = "c"
    fmt.Println("set:", s)       // set: [a b c]
    fmt.Println("get:", s[2])    // get: c

    fmt.Println("len:", len(s))  // len: 3

    s = append(s, "d")           // Append an element to s, increasing its size
    s = append(s, "e", "f")      // Append multiple elements -- a useful builtin
    fmt.Println("appended:", s)  // appended: [a b c d e f]

    c := make([]string, len(s))  // Create an array of strings with the length s
    copy(c, s)                   // Copy the elements of s into c
    fmt.Println("copy:", c)      // copy: [a b c d e f]

    l := s[2:5]                  // l is a slice of s, [2, 5)
    fmt.Println("sl1:", l)       // sl1: [c d e]

    l = s[:5]                    // l is a slice of s, [0, 5)
    fmt.Println("sl2:", l)       // slt: [a b c d e]

    l = s[2:]                    // l is a slice of s, [2, len(s))
    fmt.Println("sl3:", l)       // sl3: [c d e f]

    t := []string{"g", "h", "i"} // t, a slice with three specified elements
    fmt.Println("dcl:", t)       // dcl: [g h i]

    twoD := make([][]int, 3)            // Slices can be composed into
    for i := 0; i < 3; i++ {            // multi-dimensional data structures.
        innerLen := i + 1               // The length of the inner slices can
        twoD[i] = make([]int, innerLen) // vary, unlike with their
        for j := 0; j < innerLen; j++ { // multi-dimensional array counterparts.
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD) // 2d: [[0] [1 2] [2 3 4]]
}
```

Ultimately: remember that arrays are different types than slices, though they
can be thought of in similar ways.



[Slices, an exploration](https://go.dev/blog/slices-intro)
[Cute tricks](https://github.com/golang/go/wiki/SliceTricks)
