package utils_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	utils "github.com/FlatDigital/core-go-toolkit/utils"
)

func TestAnStringWith10LengthAnd3AsMaxValueLengthShouldTruncateStringTo3Length(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(3, len(utils.TruncateString("0123456789", 3)))
	assertion.Equal("012", utils.TruncateString("0123456789", 3))
}

func TestAnStringWith10LengthAnd15AsMaxValueLengthShouldReturnOriginalValue(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(10, len(utils.TruncateString("0123456789", 15)))
	assertion.Equal("0123456789", utils.TruncateString("0123456789", 15))
}

func TestSplitByWords(t *testing.T) {
	ass := assert.New(t)

	firstStr, secondStr := utils.SplitByWords("AA AA BB BB", 2)
	ass.Equal("AA AA", firstStr)
	ass.Equal("BB BB", secondStr)

	firstStr, secondStr = utils.SplitByWords("AA AA BB BB", 4)
	ass.Equal("AA AA BB BB", firstStr)
	ass.Equal("", secondStr)
}

func Test_CompareStrings_OK(t *testing.T) {
	ass := assert.New(t)

	// The distance between both strings should be exactly 0.2 = 1 / 5
	isValid := utils.CompareStrings("aaaaa", "aabaa", 0.2)
	ass.True(isValid)
}

func Test_CompareStrings_EmptyString(t *testing.T) {
	ass := assert.New(t)

	// The distance between both strings should be exactly 0.2 = 1 / 5
	isValid := utils.CompareStrings(" ", "aabaa", 0.2)
	ass.False(isValid)
}

func Test_CompareStrings_LongNames(t *testing.T) {
	ass := assert.New(t)

	name1 := "WOARLEN JONAS MATIAS DOS SANTOS JONAS"
	name2 := "WOARLEN JONAS MATIAS DOS SANTOS"
	ass.True(utils.CompareStrings(name1, name2, 0.2))
	ass.True(utils.CompareStrings(name2, name1, 0.2))

	name1 = "SAMUEL HENRIQUE DA SILVA PEREIRA FO"
	name2 = "SAMUEL HENRIQUE DA SILVA PEREIRA"
	ass.True(utils.CompareStrings(name1, name2, 0.2))
	ass.True(utils.CompareStrings(name2, name1, 0.2))

	name1 = "Adrielly Iris Costa Dos Santos"
	name2 = "ADRIELLY IRIS COSTA DOS SANTOS NOU"
	ass.True(utils.CompareStrings(name1, name2, 0.2))
	ass.True(utils.CompareStrings(name2, name1, 0.2))

	name1 = "SONAIDE GOUVEIA DA SILVA"
	name2 = "SONA I DE GOUVEIA DA SILVA"
	ass.True(utils.CompareStrings(name1, name2, 0.2))
	ass.True(utils.CompareStrings(name2, name1, 0.2))

	ass.False(utils.CompareStrings("este es un nombre largo", "este es otro nombre largo pero no deberian matchear", 0.5))
	ass.False(utils.CompareStrings("SAMUEL HENRIQUE DA SILVA PEREIRA", "CARLOS HENRIQUE DA SILVA", 0.5))
}

func Test_CompareStrings_DifferentLens(t *testing.T) {
	ass := assert.New(t)

	isValid := utils.CompareStrings("aabaa a", "aabaa", 0.2)
	ass.False(isValid)
}

func Test_RemoveNotAlphaNumericValues(t *testing.T) {
	ass := assert.New(t)
	str := "(*&^%$#@1234*&^%$#@"
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	result := utils.RemoveNotAlphaNumericValues(str)

	ass.Equal(result, "1234")
	ass.False(reg.MatchString(result))
}

func Test_GetAllIndexPositions(t *testing.T) {
	ass := assert.New(t)
	str1 := "ÀAZÀ"
	str2 := "ÑAÑZ"
	str3 := "AAZA"
	str4 := "AA?Z?A"

	ass.Equal([]int{0, 3}, utils.GetAllIndexPositions(str1, 'À'))
	ass.Equal([]int{0, 2}, utils.GetAllIndexPositions(str2, 'Ñ'))
	ass.Equal([]int{0, 1, 3}, utils.GetAllIndexPositions(str3, 'A'))
	ass.Equal([]int{2, 4}, utils.GetAllIndexPositions(str4, '?'))
}

func Test_SanitizeString_Success(t *testing.T) {
	// Given
	ass := assert.New(t)
	// string with special character "\u00"
	testString := "Alan Mat­as Moreno"

	// When
	sanitizedString, hasInvalidChar := utils.SanitizeString(testString)

	// Then
	ass.Equal("Alan Mat\u00adas Moreno", testString)
	ass.Equal(sanitizedString, "Alan Matas Moreno")
	ass.True(hasInvalidChar)
}

func Test_SanitizeString_Success_No_Invalid_Chars(t *testing.T) {
	// Given
	ass := assert.New(t)
	// string with no special character "\u00"
	testString := "Alan Matas Moreno"

	// When
	_, hasInvalidChar := utils.SanitizeString(testString)

	// Then
	ass.False(hasInvalidChar)
}

func Test_SanitizeString_Success_JSON(t *testing.T) {
	// Given
	ass := assert.New(t)
	testString := `[
        {
            "information_bucket": {
                "users_api": {
                    "first_name": "Alan Mat­as Moreno",
                }
            },
        }
    ]`

	expectedString := `[        {            "information_bucket": {                "users_api": ` +
		`{                    "first_name": "Alan Matas Moreno",                }            },        }    ]`

	// When
	sanitizedString, hasInvalidChar := utils.SanitizeString(testString)

	// Then
	ass.Equal(expectedString, sanitizedString)
	ass.NotEqual(sanitizedString, testString)
	ass.True(hasInvalidChar)
}

func Test_DeleteExtraSpace(t *testing.T) {
	ass := assert.New(t)

	ass.Equal(utils.DeleteExtraSpace("eireli  Me"), "eireli Me")
	ass.Equal(utils.DeleteExtraSpace("2135   Empresario individual mei"), "2135 Empresario individual mei")
}
