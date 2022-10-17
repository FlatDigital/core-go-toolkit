package utils

import "strconv"

// BoolFromString Cast a String to a Bool
// If the string is not bool casteable a false value is returned
func BoolFromString(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

// StringFromBool cast a bool to a string
func StringFromBool(b bool) string {
	return strconv.FormatBool(b)
}

// FormatBoolBinary returns "1" or "0" according to the value of b
func FormatBoolBinary(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
