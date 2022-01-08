### (To Be Or Not To Be) For Tis Nobler in the Func to Loop...

`for` is an excellent type of loop. Here are use-cases:

```
package main

import "fmt"

func main() {

    i := 1                      // Short declaration
    for i <= 3 {                // Simple condition
        fmt.Println(i)          // Print, increment...
        i = i + 1               // continue
    }

    for j: = 7; j <= 9; j++ {   // Short delcaration, condition, rule
        fmt.Println(j)          // Action
    }

    for {                       // Infinite loop
        fmt.Println("loop")
        break                   // One of the only ways out!
    }

    for n := 0; n <= 5, n++ {   // Normal
        if n%2 == 0 {
            continue            // Continue! So powerful.
        }
        fmt.Println(n)
    }
}
```


For an advanced version of `for`, checkout [Methods](/exercism/go/Methods/README.md)!
