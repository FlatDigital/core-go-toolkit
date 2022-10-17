package datetime

import "time"

type (
	// Clock defines a clock functionality
	Clock interface {
		Now() time.Time
	}

	clock struct {
		Clock
	}
)

// NewClock returns a new instance of clock object
func NewClock() Clock {
	return &clock{}
}

// Now Returns a timestamp
func (clock *clock) Now() time.Time {
	return time.Now().In(time.FixedZone("UTC", 0))
}
