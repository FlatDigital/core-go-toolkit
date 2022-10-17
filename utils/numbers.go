package utils

import (
	"fmt"
	"math"
	"strconv"
)

// Uint64FromString Cast a String to an Uint64
// If the string is not number casteable a zero value is returned
func Uint64FromString(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// Int64FromString Cast a String to an Int64
// If the string is not number casteable a zero value is returned
func Int64FromString(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// Uint64FromInterface Cast a interface{} to an Uint64
// If the interface{} is not number casteable a zero value is returned
func Uint64FromInterface(i interface{}) (uint64, error) {
	var err error
	var ok bool
	var str string

	// uint
	uint64Value, ok := i.(uint64)
	if ok {
		return uint64Value, nil
	}

	// int
	int64Value, ok := i.(int64)
	if ok {
		if int64Value < int64(0) {
			return uint64(0), fmt.Errorf("can not parse int %d to a uint64", int64Value)
		}
		return uint64(int64Value), nil
	}

	// float64
	float64Value, ok := i.(float64)
	if ok {
		if float64Value < float64(0) {
			return uint64(0), fmt.Errorf("can not parse float %f to a uint64", float64Value)
		}
		intpart, div := math.Modf(float64Value)
		if div > float64(0) {
			return uint64(0), fmt.Errorf("can not parse float %f to a uint64", float64Value)
		}
		value := uint64(intpart)
		return value, nil
	}

	// string
	strValue := fmt.Sprintf("%v", i)
	value, err := Uint64FromString(strValue)
	if err != nil {
		return uint64(0), fmt.Errorf("can not parse string %s to a uint64", str)
	}

	return value, nil
}

// StringFromFloat64 cast a float64 to a string
func StringFromFloat64(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

// StringFromUInt64 cast a uint64 to a string
func StringFromUInt64(n uint64) string {
	return fmt.Sprintf("%d", n)
}

// Uint8FromString Cast a String to a Uint8
// If the string is not number casteable a zero value is returned
func Uint8FromString(s string) (uint8, error) {
	u, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return uint8(0), err
	}
	return uint8(u), nil
}

// GetLastDigit returns the last digit from the number applying mod 10
func GetLastDigit(n uint64) uint64 {
	return n % 10
}

// Float64FromString cast a float64 from a valid number string
func Float64FromString(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return float64(0), err
	}
	return f, nil
}
