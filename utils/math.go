package utils

import (
	"fmt"
	"math"
	"strconv"
)

// RoundingMode type for rounding mode
type RoundingMode float64

const (
	// RoundingModeHalfUp half up rounding mode
	RoundingModeHalfUp RoundingMode = .5
	// RoundingModeHalfDown half down rounding mode
	RoundingModeHalfDown RoundingMode = .5 + math.SmallestNonzeroFloat64
)

// defaultRoundingMode default rounding mode
const defaultRoundingMode = RoundingModeHalfUp

// Round round a float64 using the default rounding mode
func Round(val float64, places int) float64 {
	return round(val, float64(defaultRoundingMode), places)
}

// RoundWithRoundingMode rounds a float64 using an especific rounding mode
func RoundWithRoundingMode(val float64, places int, roundingMode RoundingMode) float64 {
	return round(val, float64(roundingMode), places)
}

// rounds a float64 number
func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	var base float64 = 10
	pow := math.Pow(base, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return newVal
}

//

// RoundFloor rounds to floor for the digit of places
func RoundFloor(val float64, places int) (newVal float64) {
	var round float64
	var base float64 = 10
	pow := math.Pow(base, float64(places))
	digit := pow * val
	round = math.Floor(digit)
	newVal = round / pow
	return newVal
}

// RoundCeiling rounds to ceiling for the digit of places
func RoundCeiling(val float64, places int) (newVal float64) {
	var round float64
	var base float64 = 10
	pow := math.Pow(base, float64(places))
	digit := pow * val
	round = math.Ceil(digit)
	newVal = round / pow
	return newVal
}

// MinInt overloads math.Min for ints
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt overloads math.Max for ints
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// EqualFloatValues compare two float64 as string comparation and returns if they are equals
func EqualFloatValues(val1, val2 float64, precision int32) bool {
	valString1 := fmt.Sprintf("%."+strconv.Itoa(int(precision))+"f", val1)
	valString2 := fmt.Sprintf("%."+strconv.Itoa(int(precision))+"f", val2)
	return valString1 == valString2
}
