package utils_test

import (
	"testing"

	utils "github.com/FlatDigital/core-go-toolkit/v2/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Contains_True(t *testing.T) {
	assertion := assert.New(t)

	slice := []string{"a", "b", "c"}
	contains := utils.Contains(slice, "a")

	assertion.True(contains)
}

func Test_Contains_False(t *testing.T) {
	assertion := assert.New(t)

	slice := []string{"a", "b", "c"}
	contains := utils.Contains(slice, "X")

	assertion.False(contains)
}

func Test_Remove(t *testing.T) {
	assertion := assert.New(t)

	slice := []string{"a", "b", "c", "d"}

	assertion.Len(utils.Remove(slice, "a"), 3)
	assertion.Len(utils.Remove(slice, "d"), 3)
	assertion.Len(utils.Remove(slice, "A"), 4)
	assertion.Len(utils.Remove(slice, "e"), 4)
}

func Test_ToStringSlice(t *testing.T) {
	assertion := assert.New(t)

	result := utils.ToStringSlice("a", 2, "c", 3)

	assertion.Len(result, 4)
	assertion.Equal("a", result[0])
	assertion.Equal("2", result[1])
	assertion.Equal("c", result[2])
	assertion.Equal("3", result[3])
}

func Test_Reverse(t *testing.T) {
	ass := assert.New(t)
	arr1 := []string{"Penn", "Teller"}
	expectedResult := []string{"Teller", "Penn"}
	arr2 := []string{"David"}

	ass.Equal(utils.Reverse(arr2), arr2)
	ass.Equal(expectedResult, utils.Reverse(arr1))
}

func Test_FindPosition(t *testing.T) {
	assertion := assert.New(t)

	slice := []string{"a", "b", "c"}
	contains := utils.FindPosition(slice, "X")

	assertion.Equal(contains, -1)

	contains = utils.FindPosition(slice, "a")

	assertion.Equal(contains, 0)
}

func Test_SwitchPositions(t *testing.T) {
	assertion := assert.New(t)

	slicePreOrder := []string{"a", "b", "c"}
	slicePostOrder := utils.SwitchPositions(slicePreOrder, 0, 1)

	assertion.Equal(slicePostOrder, []string{"b", "a", "c"})
}
