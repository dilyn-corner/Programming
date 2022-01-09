/*
Package grains provides a loosely effective way of determining how many grains
it takes to impoverish a king
*/
package grains

import (
	"errors"
	"math"
)

// Square provides sane error messages and does some quick exponentiation
func Square(number int) (uint64, error) {
	if number < 0 {
		return 0, errors.New("squaring a negative")
	} else if number == 0 {
		return 0, errors.New("this isn't a square")
	} else if number > 64 {
		return 0, errors.New("too many squares")
	}

	return uint64(math.Exp2(float64(number - 1))), nil
}

// Total doesn't succumb to cheating with bit shifting
func Total() uint64 {
	var totalGrains uint64
	for i := 1; i < 65; i++ {
		totalGrains += uint64(math.Exp2(float64(i) - 1))
	}
	return totalGrains
}
