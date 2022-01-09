/*
Package hamming calculates the Hamming distance
between two equally sized DNA strands
*/
package hamming

import "errors"

// Distance returns the number of character differences between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands are of different length")
	}

	var count int

	for i := 0; i < len(a); i++ {
		switch {
		case a[i] != b[i]:
			count++
		}
	}
	return count, nil
}
