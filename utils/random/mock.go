package random

type Mock struct {
	float64MockStack []float64
}

const (
	floatValue float64 = 0.1
)

// NewRandomMock returns a new instance of the random mock object
func NewRandomMock() *Mock {
	return &Mock{
		float64MockStack: []float64{},
	}
}

// PatchFloat64 Returns a random float value mocked
func (mock *Mock) PatchFloat64(inputFloat64 float64) {
	mock.float64MockStack = append(mock.float64MockStack, inputFloat64)
}

// Float64 Returns a random float value
func (mock *Mock) Float64() float64 {
	float := floatValue
	if len(mock.float64MockStack) > 0 {
		float = mock.float64MockStack[0]
		// dequeues
		mock.float64MockStack = mock.float64MockStack[1:]
	}

	return float
}
