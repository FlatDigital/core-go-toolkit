package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bool(t *testing.T) {
	ass := assert.New(t)

	ass.True(BoolFromString("true"))
	ass.True(BoolFromString("1"))
	ass.True(BoolFromString("T"))

	ass.False(BoolFromString("false"))
	ass.False(BoolFromString("0"))
	ass.False(BoolFromString("F"))

	ass.False(BoolFromString("string that not represent a bool value"))

	ass.Equal("false", StringFromBool(false))
	ass.Equal("true", StringFromBool(true))
}

func Test_FormatBoolBinary_WithTrue_ShouldReturn_1StringValue(t *testing.T) {
	assertions := assert.New(t)

	actual := FormatBoolBinary(true)
	assertions.EqualValues("1", actual)
}

func Test_FormatBoolBinary_WithFalse_ShouldReturn_0StringValue(t *testing.T) {
	assertions := assert.New(t)

	actual := FormatBoolBinary(false)
	assertions.EqualValues("0", actual)
}
