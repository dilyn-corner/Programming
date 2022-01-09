/*
Package scrabble offers a quick way of calculating the point value of 
any given word in the popular game Scrabble.
It does not implement any special or extravagant features (triple word bonus,
for instance), but it at least counts letter values!
*/
package scrabble

import "strings"

// Score returns the sums of the values of the letters in word
func Score(word string) int {

	capWord := strings.ToUpper(word)
	var count int
	for i := 0; i < len(word); i++ {
		switch string(capWord[i]) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			count++
		case "D", "G":
			count += 2
		case "B", "C", "M", "P":
			count += 3
		case "F", "H", "V", "W", "Y":
			count += 4
		case "K":
			count += 5
		case "J", "X":
			count += 8
		case "Q", "Z":
			count += 10
		}
	}
	return count
}

// I have a really cute idea of doing something like in the chessboard exercise.
// Instead of just hard-coding certain values into the outcomes of a switch {},
// We can take advantage of arrays and maps.
// It's... Hairy... The implementation details are foggy.

// type Letters []string
// type Scores map[int]Letters
//
// oneLetters   := Letters{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
// twoLetters   := Letters{"D", "G"}
// threeLetters := Letters{"B", "C", "M", "P"}
// fourLetters  := Letters{"F", "H", "V", "W", "Y"}
// fiveLetters  := Letters{"K"}
// eightLetters := Letters{"J", "X"}
// tenLetters   := Letters{"Q", "Z"}
//
// var scoreMap Scores
// scoreMap[1]  = oneLetters
// scoreMap[2]  = twoLetters
// scoreMap[3]  = threeLetters
// scoreMap[4]  = fourLetters
// scoreMap[5]  = fiveLetters
// scoreMap[8]  = eightLetters
// scoreMap[10] = tenLetters
//
// var i, count int
// for i, l := range scoreMap {
// 	for j := 0; j < len(word); j++ {
// 		switch string(word[j]) {
// 			case l:
// 				count += i
// 			}
// 		}
// 	}
// return count
