### The Keys to my Values

Maps are relatively complicated data types. It's specifically because of what
sorts of items can get rammed into them. Fundamentally, maps are just `key, value`
pairs. You might better know them as dictionaries.

Here are some fundamental facts about maps worth documenting:

```
package main

import "fmt"

func main() {

    m := make(map[string]int) // alternatively m := map[string]int{}

    m["k1"] = v1 // set the k1, v1 pair
    m["k2"] = v2 // set the k2, v2 pair

    fmt.Println("map:", m) // print the map:
                           // map: map[k1:v1 k2:v2]

    u1 := m["k1"]           // assign u1 to the value of the k1 key's value
    fmt.Println("v1: ", u1) // Print the value assigned to u1:
                            // v1: v1

    fmt.Println("len:", len(m)) // Print the length of the map m using the len() builtin
                                // len: 2
                                // len(m) is the number of key, value pairs

    delete(m, "k2")         // The delete() builtin does what it says on the tin:
    fmt.Println("map:", m)  // Print the freshly modified map:
                            // map: map[k1:v1]

    _, present := m["k2"]               // We omit the first return value here
                                        // because we aren't concerned with it
                                        // the second return value stores
                                        // whether or not there is a k2 key in 
                                        // the map m which has a value.
                                        // Because we deleted the k2 key...
    fmt.Println("is present:", present) // present: false
                                        // If we hadn't omitted the first return
                                        // value, we would have been returned v2, true

    n := map[string]int{"foo": 1, "bar": 2} // A shortcut for map
                                            // creation & initialization
```

Please go to [Iterations](/exercism/go/Iterations/README.md) for more fuckery
with maps.
