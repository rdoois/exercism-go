package clock

import (
	"fmt"
)

// Clock is a type
type Clock struct {
	h int
	m int
}

var minPerHour = 60
var hourPerDay = 24

func norm(h int, m int) (int, int) {
	for m >= minPerHour {
		m -= minPerHour
		h++
	}
	for h >= hourPerDay {
		h -= hourPerDay
	}
	for m < 0 {
		m += minPerHour
		h--
	}
	for h < 0 {
		h += hourPerDay
	}
	return h, m
}

// New is a constructor
func New(h int, m int) Clock {
	h, m = norm(h, m)
	return Clock{h, m}
}

// Add add min
func (c Clock) Add(min int) Clock {
	return New(c.h, c.m+min)
}

// Subtract substract min
func (c Clock) Subtract(min int) Clock {
	return New(c.h, c.m-min)
}
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
