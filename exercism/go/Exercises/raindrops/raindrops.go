/*
Package raindrops returns a simple string depending on which
of three specific numbers any given number is divisible by.
*/
package raindrops

import "strconv"

// Convert returns a particular string depending on the factors of number
func Convert(number int) string {
	var badFizz string
	if number%3 == 0 {
		badFizz += "Pling"
	}

	if number%5 == 0 {
		badFizz += "Plang"
	}

	if number%7 == 0 {
		badFizz += "Plong"
	}

	if badFizz == "" {
		return strconv.Itoa(number)
	}

	return badFizz
}
