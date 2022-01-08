### Struct by Lightning

A struct is a type! A struct type is nothing but a schema containing the
blueprint of a data a structure will hold. For instance:

```
type structName struct {
    field1 fieldType1
    field2 fieldType2
}
```

A struct collects data, organized into fields and their associated types. As an
example, we can say that `ross` is a type of `Employee` (struct type) which has
`firstName`, `lastName`, `salary`, and `fullTime` properties (structure fields):

Two alternative ways of writing our Employee struct:
```
type Employee struct {          |      type Employee struct { 
    firstName string            |          firstName, lastName string
    lastname string             |          salary int
    salary int                  |          fullTime bool
    fullTime bool               |
```

`ross`, in this example, is our struct _proper_. `Employee` is merely a *struct
type*. We can create our `ross` struct and fill in the values in two ways:

```
var ross Employee               |       ross := Employee {
ross.firstName = "ross",        |           firstName: "ross",
ross.LastName = "Bing",         |           lastName: "Bing",
ross.salary = 1200,             |           salary: 1200,
ross.fullTime = true,           |           fullTime: true,
                                |       }
```

Note that you can also omit the values of some fields when creating your struct.

We can also have pointers to structs!

```s := &StructType{...}```

We can do this with our `ross` struct:
```
ross := &Employee {
    firstName: "ross",
    lastName: "Bing",
    salary: 1200,
    fullTime: true,
}
```

Because `ross` is a pointer, we have to dereference it to get the values back
out... Specifically, to get the `firstName` value, we must use
`(*ross).firstName`. We use parentheses here in order to disambiguate -
`*ross.firstName` could mean `(*ross).firstName` OR `*(ross.firstName)`.

... However! Go is smart and, just like with maps, we need not dereference the
fields of a struct pointer before accessing the values!

```
ross := &Employee {
    firstName: "ross",
    lastName: "Bing",
    salary: 1200,
    fullTime: true,
}

fmt.Println("firstName", ross.firstName) // firstName ross
```

[For a lot more information](https://medium.com/rungo/structures-in-go-76377cc106a2) (and a
large source of these examples)
