### A Pointer or Two...

A pointer simply holds the memory address of some value.

The type `*T` is a pointer to a T value. The zero value of a pointer is nil:
```var p *int```

The `&` operator generates a pointer to its operand:
```
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value:

```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```
Doing this operations is known as "dereferencing" or "indirecting".


A more comprehensive example:
```
package main

import "fmt"

func zeroval(ival int) { // Just some regular function. 
     ival = 0            // Arguments are passed to it as values
}

func zeroptr(iptr *int) { // Takes an *int parameter - it takes an int pointer
    *iptr = 0             // The pointer is then dereferenced by the * operator
}                         // Because we are assigning a value to a dereferenced
                          // pointer, the actual value at that address is modified

func main() {
    i := 1
    fmt.Println("initial:", i) // 1

    zeroval(i)
    fmt.Println("zeroval:", i) // 1 (the i passed to zeroval() is not modified by zeroval()

    zeroptr(&i)
    fmt.Println("zeroptr:", i) // 0 (the &i passed to zeroval() IS modified by zeroval()
                               // In other words, the value assigned to the
                               // memory address at &i is modified by zeroval(),
                               // and then we print that value

    fmt.Println("pointer:", &i) // 0x42131100 <- the memory address
}



[This might be useful](https://www.digitalocean.com/community/conceptual_articles/understanding-pointers-in-go).
