package robotname

import (
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {

	if r.name > "" {
		return r.name, nil
	}

	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	letter1 := string(letters[rand.Intn(26)])
	letter2 := string(letters[rand.Intn(26)])
	randNum := rand.Intn(1000)

	r.name = fmt.Sprintf("%s%s%03d", letter1, letter2, randNum)

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}
