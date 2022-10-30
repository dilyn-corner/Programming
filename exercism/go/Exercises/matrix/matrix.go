/* Package matrix creates an arbitrary matrix from a given string and parses
   that matrix to find a given element, and can change the value of it.
*/
package matrix

import "strings"

// Define the Matrix type here.

type Matrix [][]int

// New converts a string into a matrix
func New(s string) (*Matrix, error) {
}

// Cols and Rows must return the results without affecting the matrix.

//Cols outputs the columns of a matrix m
func (m *Matrix) Cols() [][]int {
}

// Rows outputs the rows of a matrix m
func (m *Matrix) Rows() [][]int {
	row := make([][]int, len(*m))
	for i, j := range *m {
		row[i] = append([]int{}, j...)
	}
	return row
}

// Set modifies the elment if an mxn matrix
func (m *Matrix) Set(row, col, val int) bool {
}
