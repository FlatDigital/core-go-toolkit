package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Uint64FromString(t *testing.T) {
	ass := assert.New(t)

	var value uint64
	var err error

	// not valid
	value, err = Uint64FromString("")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("abcde")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("12345abcde")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("12345|@#¢∞")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("-123")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("123,456")
	ass.Equal(uint64(0), value)
	ass.Error(err)
	value, err = Uint64FromString("123.456")
	ass.Equal(uint64(0), value)
	ass.Error(err)

	// valid
	value, err = Uint64FromString("0")
	ass.Equal(uint64(0), value)
	ass.Nil(err)
	value, err = Uint64FromString("1")
	ass.Equal(uint64(1), value)
	ass.Nil(err)
	value, err = Uint64FromString("2")
	ass.Equal(uint64(2), value)
	ass.Nil(err)
	value, err = Uint64FromString("18446744073709551614")
	ass.Equal(uint64(18446744073709551614), value)
	ass.Nil(err)
	value, err = Uint64FromString("18446744073709551615")
	ass.Equal(uint64(18446744073709551615), value)
	ass.Nil(err)
}

func Test_Int64FromString(t *testing.T) {
	ass := assert.New(t)

	var value int64
	var err error

	// not valid
	value, err = Int64FromString("")
	ass.Equal(int64(0), value)
	ass.Error(err)
	value, err = Int64FromString("abcde")
	ass.Equal(int64(0), value)
	ass.Error(err)
	value, err = Int64FromString("12345abcde")
	ass.Equal(int64(0), value)
	ass.Error(err)
	value, err = Int64FromString("12345|@#¢∞")
	ass.Equal(int64(0), value)
	ass.Error(err)
	value, err = Int64FromString("123,456")
	ass.Equal(int64(0), value)
	ass.Error(err)
	value, err = Int64FromString("123.456")
	ass.Equal(int64(0), value)
	ass.Error(err)

	// valid
	value, err = Int64FromString("0")
	ass.Equal(int64(0), value)
	ass.Nil(err)
	value, err = Int64FromString("1")
	ass.Equal(int64(1), value)
	ass.Nil(err)
	value, err = Int64FromString("2")
	ass.Equal(int64(2), value)
	ass.Nil(err)
	value, err = Int64FromString("51614")
	ass.Equal(int64(51614), value)
	ass.Nil(err)
	value, err = Int64FromString("-51614")
	ass.Equal(int64(-51614), value)
	ass.Nil(err)
}

func Test_StringFromUInt64(t *testing.T) {
	ass := assert.New(t)

	// valid
	ass.Equal("0", StringFromUInt64(uint64(0)))
	ass.Equal("1", StringFromUInt64(uint64(1)))
	ass.Equal("2", StringFromUInt64(uint64(2)))
	ass.Equal("18446744073709551614", StringFromUInt64(uint64(18446744073709551614)))
	ass.Equal("18446744073709551615", StringFromUInt64(uint64(18446744073709551615)))
}

func Test_Uint8FromString(t *testing.T) {
	ass := assert.New(t)

	var value uint8
	var err error

	// not valid
	value, err = Uint8FromString("")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("abcde")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("123abcde")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("123|@#¢∞")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("-123")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("123,456")
	ass.Equal(uint8(0), value)
	ass.Error(err)
	value, err = Uint8FromString("123.456")
	ass.Equal(uint8(0), value)
	ass.Error(err)

	// valid
	value, err = Uint8FromString("0")
	ass.Equal(uint8(0), value)
	ass.Nil(err)
	value, err = Uint8FromString("1")
	ass.Equal(uint8(1), value)
	ass.Nil(err)
	value, err = Uint8FromString("2")
	ass.Equal(uint8(2), value)
	ass.Nil(err)
	value, err = Uint8FromString("254")
	ass.Equal(uint8(254), value)
	ass.Nil(err)
	value, err = Uint8FromString("255")
	ass.Equal(uint8(255), value)
	ass.Nil(err)
}

func Test_Float64FromString(t *testing.T) {
	ass := assert.New(t)

	// not valid
	_, err := Float64FromString("not_valid")
	ass.Error(err)

	// valid
	f, err := Float64FromString("0.123456789")
	ass.Equal(float64(.123456789), f)
	ass.Nil(err)
}

func TestGetLastDigit(t *testing.T) {
	ass := assert.New(t)

	cases := []uint64{0, 10, 1, 00, 52}
	expected := []uint64{0, 0, 1, 0, 2}

	ass.Equal(len(cases), len(expected))

	for i := range cases {
		ass.Equal(GetLastDigit(cases[i]), expected[i])
	}
}

func Test_StringFromFloat64(t *testing.T) {
	ass := assert.New(t)

	// valid
	ass.Equal("0", StringFromFloat64(float64(0)))
	ass.Equal("1", StringFromFloat64(float64(1)))
	ass.Equal("2", StringFromFloat64(float64(2)))
	ass.Equal("1446744070.551614", StringFromFloat64(float64(1446744070.551614)))
	ass.Equal("1844674.407371615", StringFromFloat64(float64(1844674.407371615)))
}

func Test_Uint64FromInterface(t *testing.T) {
	ass := assert.New(t)

	var i interface{}
	var u uint64
	var err error

	// from uint
	i = uint64(123456789)
	u, err = Uint64FromInterface(i)
	ass.Equal(uint64(123456789), u)
	ass.Nil(err)

	// from int
	i = int64(-123456789)
	u, err = Uint64FromInterface(i)
	ass.Zero(u)
	ass.Error(err)

	i = int64(123456789)
	u, err = Uint64FromInterface(i)
	ass.Equal(uint64(123456789), u)
	ass.Nil(err)

	// from float
	i = float64(-.123456789)
	u, err = Uint64FromInterface(i)
	ass.Zero(u)
	ass.Error(err)

	i = float64(.123456789)
	u, err = Uint64FromInterface(i)
	ass.Zero(u)
	ass.Error(err)

	i = float64(123456789)
	u, err = Uint64FromInterface(i)
	ass.Equal(uint64(123456789), u)
	ass.Nil(err)

	// from string
	i = string("-.123456789")
	u, err = Uint64FromInterface(i)
	ass.Zero(u)
	ass.Error(err)

	i = string("123456789")
	u, err = Uint64FromInterface(i)
	ass.Equal(uint64(123456789), u)
	ass.Nil(err)
}
