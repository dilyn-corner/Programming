## Operators

### Arithmetic

| Symbol | Operation |
| --- | --- |
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `%` | Modulus |

### Comparison

| Symbol | Operation |
| --- | --- |
| `==` | Equality | 
| `!=` | Inequality |
| `<` | Less than |
| `<=` | Less than or equal |
| `>` | Greather than |
| `>=` | Greather than or equal |

### Logical

Note that these are evaluated conditionally!

Conjunction is perhaps similar to logical implication:

p && q -> if p then q else false

p || q -> if p then true else q

| Symbol | Operation |
| --- | --- |
| `&&` | Conjunction |
| `\|\|` | Disjunction |
| `!` | Negation |

### Address

| Symbol | Operation |
| --- | --- |
| `*x` | Variable of type T pointed to by x |
| `&x` | Address operation generating a pointer of type *T to x |
| `var x *T` | x contains the memory address of a type *T |

Rephrasing some of these:

`var x *T` -> x is of type "pointer to T (*T)". x holds the memory address of
some type T.

```
var y int
y = 2
x = &y    -> x contains the memory address of y
```

```
var b int
b = *x    -> *x fetches the value stored at the memory address of x. This is called dereferencing.
```

In short,

to create a pointer -> `*T`

to assign a pointer to a value -> `&v, v some T`

to act on the assigned pointer -> `*x = ...`


[For more](https://go.dev/ref/spec#Operators)
