/*
Package luhn is an implementation of the Luhn Algorithm
*/
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid determines whether or not some ID value is verified by the Luhn
// Algorithm
func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")
	idNumbers := make([]int, len(id))
	var idSum int

	if len(id) < 2 {
		return false
	}

	for _, digit := range id {
		if unicode.IsNumber(digit) == false {
			return false
		}
	}

	for i := 0; i < len(id); i++ {
		i, _ := strconv.Atoi(id[i : i+1])
		idNumbers = append(idNumbers, i)
	}

	for i := len(idNumbers) - 2; i >= 0; i -= 2 {
		if (idNumbers[i] * 2) > 9 {
			idNumbers[i] = (idNumbers[i] * 2) - 9
		} else {
			idNumbers[i] = idNumbers[i] * 2
		}
	}

	for i := 0; i < len(idNumbers); i++ {
		idSum += idNumbers[i]
	}

	return idSum%10 == 0
}
