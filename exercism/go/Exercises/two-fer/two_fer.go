/*
Package twofer returns a simple string when
given a name or when a name is omitted
*/
package twofer

import "fmt"

// ShareWith returns a name or 'you' in a simple string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
