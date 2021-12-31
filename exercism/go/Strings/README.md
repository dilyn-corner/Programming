A `string` in Go is an immutable sequence of bytes, which don't necessarily have
to represent characters.

A string literal is defined between double quotes:
```const name = "Jane"```

Some values need to be escaped:

| Value | | Description |
| --- | --- |
| `\a` | Alert or bell |
| `\b` | Backspace |
| `\\` | Backslash |
| `\t` | Horizontal tab |
| `\n` | Line feed or newline |
| `\f` | Form feed |
| `\r` | Carriage return |
| `\v` | Vertical tab |
| `\'` | Single quote |
| `\"` | Double quote |

```const daltons = "Joe\nWilliam\nJack\nAverell"```

Raw string literals use backtics (\`) as their delimiter instead of double
quotes and are interpreted literally, meaning that there is no need to escape
characters or newlines:
```
const daltons = `Joe
William
Jack
Averell`
```
The strings package contains many useful functions to work on strings:

| Function | Purpose |
| --- | --- |
| `ToLower` | Convert the string to lower case |
| `ToUpper` | Convert the string to upper case |
| `TrimSpace` | Trim leading and trailing whitespace |
| `Index` | Find the index of the first instance of a substring within a string |
| `Replace` | Replace one occurence of a substring in a string |
| `ReplaceAll` | Replace all occurrences of a substring in a string |
| `Split` | Split a string into parts based on a separator |
| `HasSuffix` | Check if a string ends with a specific substring |
| `Count` | Count the number of occurrences of a substring within a string |

```
import "strings"

strings.ToLower("Gopher")    // Output: "gopher"
strings.Index("Apple", "le") // Output: 3
strings.Count("test", "t")   // Output: 2
