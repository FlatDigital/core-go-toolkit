package utils

import "fmt"

// Contains looks for string inside a []string
func Contains(slice []string, key string) bool {
	for _, value := range slice {
		if value == key {
			return true
		}
	}
	return false
}

// Remove removes an elem from a []string
func Remove(slice []string, elem string) []string {
	for index, value := range slice {
		if value == elem {
			return append(slice[:index], slice[index+1:]...)
		}
	}
	return slice
}

// ToStringSlice converts a slice of interface{} to a slice of string
func ToStringSlice(slice ...interface{}) []string {
	result := []string{}
	for _, value := range slice {
		result = append(result, fmt.Sprintf("%v", value))
	}
	return result
}

func Reverse(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(Reverse(input[1:]), input[0])
}

func FindPosition(slice []string, toFind string) int {
	for i, x := range slice {
		if x == toFind {
			return i
		}
	}
	return -1
}

func SwitchPositions(slice []string, firstPos int, secondPos int) []string {
	aux := slice[firstPos]
	slice[firstPos] = slice[secondPos]
	slice[secondPos] = aux
	return slice
}
