### The Methods to my Madness

Go does not have classes. For our hubris, the gods have punished us with
methods.

Methods have a special argument called a _receiver_. The receiver appears in its
own argument list beween the `func` keyword and the method's name.

Compare:
```
func Abs(v Vertex) float64 { ... }    // Abs() as some function
func (v Vertex) Abs() float64 { ... } // Abs() as some method
```

Methods are just functions with a receiver argument. The difference:
```
pkg.Function(arg)
obj.Method(arg)
```

A more involved and advanced exploration:
```
package main

import "fmt"

type rect struct {      // Define the rect struct
    width, height int   // It contains the dimensions of a rectangle
}

func (r *rect) area() int {     // The area() method has a receiver type of *rect
    return r.width * r.height
}

func (r rect) perim() int {     // Methods can be for either *pointers or values
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}     // define some rectangle r...

    fmt.Println("area: ", r.area())     // Call the two methods on our struct
    fmt.Println("perim: ", r.perim())

    rp := &r                            // Conveniently, Go handles conversions
    fmt.Println("area: ", rp.area())    // between pointers and values for method calls
    fmt.Println("perim: ", rp.perim())
// You may want to use a pointer receiver type to avoid copying on method calls!
// Or if you want to allow the method to actually change the receiver struct
}
```

[A useful tour](https://go.dev/tour/methods/1).
[Perhaps some more information](https://www.callicoder.com/golang-methods-tutorial/)
