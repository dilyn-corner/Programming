/*
Package tournament provides a function which takes a well-defined sort of string
describing the outcomes of a game and adds that information to a table. That
table of outcomes is then sorted by number of points in descending order.
*/
package tournament

import (
	"encoding/csv"
	"io"
)

type outcome int

const (
	loss outcome = iota
	draw
	win
)

type inputEntry struct {
	teams    [2]string
	outcomes [2]outcome
}

type teamResult struct {
	team   string
	played int
	wins   int
	draws  int
	losses int
	points int
}

// Tally takes input strings describing the outcomes of games and returns a
// table of those outcomes, sorted by total points earned.
func Tally(reader io.Reader, writer io.Writer) error {
	var entries []inputEntry
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields
	for {
		record, err := csvReader.Read()
		if err != nil {
			return err
		}
		if len(record) == 3 {
			t1, t2 := record[0], record[1]
			teams := [2]string{t1, t2}
			var outcomes [2]outcome
			switch record[2] {
			case "win":
				outcomes = [2]outcome{win, loss}
			case "loss":
				outcomes = [2]outcome{loss, win}
			case "draw":
				outcomes = [2]outcome{draw, draw}
			default:
				return nil
			}
			entries = append(entries, inputEntry{teams: teams, outcomes: outcomes})
		} else {
			return nil
		}
	}
}
