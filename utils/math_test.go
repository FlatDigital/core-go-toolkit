package utils_test

import (
	"fmt"
	"math"
	"testing"

	utils "github.com/FlatDigital/core-go-toolkit/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Round(t *testing.T) {
	ass := assert.New(t)
	var base float64 = 10
	for i := 0; i < 9; i++ {
		f := float64(i) / base

		// default / half up
		if f >= .5 {
			ass.Equal(float64(1), utils.Round(f, 0))
			ass.Equal(float64(1), utils.RoundWithRoundingMode(f, 0, utils.RoundingModeHalfUp))
		} else {
			ass.Equal(float64(0), utils.Round(f, 0))
			ass.Equal(float64(0), utils.RoundWithRoundingMode(f, 0, utils.RoundingModeHalfUp))
		}

		// half down
		if f == 0 {
			ass.Equal(float64(-1), utils.RoundWithRoundingMode(f-math.SmallestNonzeroFloat64, 0, utils.RoundingModeHalfDown))
		} else if f >= .5 {
			ass.Equal(float64(1), utils.RoundWithRoundingMode(f-math.SmallestNonzeroFloat64, 0, utils.RoundingModeHalfDown))
		} else {
			ass.Equal(float64(0), utils.RoundWithRoundingMode(f-math.SmallestNonzeroFloat64, 0, utils.RoundingModeHalfDown))
		}

	}
}

func Test_RoundFloorCeiling(t *testing.T) {
	ass := assert.New(t)
	var divider float64 = 12.34

	f := float64(1) / divider
	ass.Equal(float64(0.08), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(2) / divider
	ass.Equal(float64(0.16), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(3) / divider
	// 0.243112 --> 0.24
	ass.Equal(float64(0.24), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(4) / divider
	// 0.324149 --> 0.32
	ass.Equal(float64(0.32), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(5) / divider
	// 0.405186 --> 0.4
	ass.Equal(float64(0.4), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(6) / divider
	// 0.486224 --> 0.48
	ass.Equal(float64(0.48), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(7) / divider
	// 0.567261 --> 0.56
	ass.Equal(float64(0.56), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(8) / divider
	// 0.648298 --> 0.64
	ass.Equal(float64(0.64), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	f = float64(9) / divider
	// 0.729335 --> 0.72
	ass.Equal(float64(0.72), utils.RoundFloor(f, 2), fmt.Sprintf("%f", f))

	// other cases

	// 0.547593 --> 0.54
	ass.Equal(float64(0.54), utils.RoundFloor(0.547593, 2))
	ass.Equal(float64(0.55), utils.RoundCeiling(0.547593, 2))

	// -0.547593 --> -0.54
	ass.Equal(float64(-0.55), utils.RoundFloor(-0.547593, 2))
	ass.Equal(float64(-0.54), utils.RoundCeiling(-0.547593, 2))

	// -0.0222
	ass.Equal(float64(-0.03), utils.RoundFloor(-0.0222, 2))
	ass.Equal(float64(-0.02), utils.RoundCeiling(-0.0222, 2))
}

func Test_EqualFloat_Equals(t *testing.T) {
	assertions := assert.New(t)

	assertions.True(utils.EqualFloatValues(46781.01234567890123, 46781.01234567890123, 14))
	assertions.True(utils.EqualFloatValues(46781.01234567890123, 46781.01234567890123, 4))
	assertions.True(utils.EqualFloatValues(21.012345678901234, 21.012345678901235, 13))
}

func Test_EqualFloat_NOT_Equals(t *testing.T) {
	assertions := assert.New(t)

	assertions.False(utils.EqualFloatValues(21.012345678901234, 21.012345678901236, 14))
	assertions.False(utils.EqualFloatValues(46781.01234557890123, 46781.01234567890123, 14))
	assertions.False(utils.EqualFloatValues(46782.01234567890123, 46781.01234567890123, 14))
}

func Test_MinInt(t *testing.T) {
	ass := assert.New(t)

	min := utils.MinInt(3, 5)
	ass.Equal(min, 3)
}

func Test_MaxInt(t *testing.T) {
	ass := assert.New(t)

	max := utils.MaxInt(3, 5)
	ass.Equal(max, 5)
}
