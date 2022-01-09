/*
Package isogram determines whether or not a word is an isogram:
https://en.wikipedia.org/wiki/Isogram
*/
package isogram

import "strings"

// IsIsogram checks if any letters in a word repeat
func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)
	for index, letter := range lowerWord {
		if letter != ' ' &&
			letter != '-' &&
			index < strings.LastIndex(lowerWord, string(letter)) {
			return false
		}
	}
	return true
}
