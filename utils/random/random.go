package random

import "math/rand"

type (
	// Random defines the random functionality
	Random interface {
		Float64() float64
	}

	random struct{}
)

// NewRandom returns a new instance of the random object
func NewRandom() Random {
	return &random{}
}

// Now Returns a timestamp
func (random *random) Float64() float64 {
	return rand.Float64()
}
