/*
Package clock implements a straightforward way of modifying the 24-hour time on
a clock, as well as comparing the times on two different clocks.
*/
package clock

import "fmt"

// Define the Clock type here.

type Clock struct {
	hours, minutes int
}

// New Creates a new clock, but stores the time on it in minutes
func New(h, m int) Clock {
	for m < 60 {
		m += 60
		h--
	}
	for m >= 60 {
		m -= 60
		h++
	}

	for h >= 24 {
		h -= 24
	}

	for h < 0 {
		h += 24
	}

	return Clock{h, m}
}

// Add increments the number of minutes on a given clock
func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

// Subtract decrements the number of minutes on a given clock
func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

// String returns the time on a clock face
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
