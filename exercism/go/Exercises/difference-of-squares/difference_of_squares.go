/*
Package diffsquares is an algorithm of dubious efficiency which finds the
difference between the square of the sums and the sum of the squares of the
first n natural numbers.
*/
package diffsquares

// SquareOfSum squares the sum of the first n natural numbers
func SquareOfSum(n int) int {
	var easySum int
	easySum = (n * (n + 1)) / 2

	return easySum * easySum
}

// SumOfSquares sums the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * ((2 * n) + 1)) / 6
}

// Difference finds the difference between SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
