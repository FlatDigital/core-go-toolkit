package utils

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode"
)

const longNameLength = 10

// TruncateString evaluates if value length is greater than max value length:
// True: truncate up to max value length
// False: returns original value
func TruncateString(value string, maxValueLength int) string {
	if len(value) > maxValueLength {
		return value[:maxValueLength]
	}
	return value
}

// FillWithTrailingChar adds filler strings as prefix until the expected length is reached
func FillWithTrailingChar(source, fill string, expectedLength int) string {
	ans := source
	for len(ans) < expectedLength {
		ans = fmt.Sprintf("%s%s", fill, ans)
	}
	return ans
}

// SplitByWords splits the given strings into two, keeping "n" words in the first string
func SplitByWords(s string, n int) (string, string) {
	words := strings.Split(s, " ")
	if len(words) <= n {
		return s, ""
	}
	firstStringWords := words[:n]
	secondStringWords := words[n:]
	return strings.Join(firstStringWords, " "), strings.Join(secondStringWords, " ")
}

// CompareStrings checks if the distance between string1 and string2 is less than the given tolerance
// Bigger tolerance means that both strings can be different in more percentage of characters
func CompareStrings(string1, string2 string, tolerance float64) bool {
	if len(strings.TrimSpace(string1)) == 0 || len(strings.TrimSpace(string2)) == 0 {
		return false
	}

	string1 = strings.ToUpper(string1)
	string2 = strings.ToUpper(string2)

	joinedString1 := strings.ReplaceAll(string1, " ", "")
	joinedString2 := strings.ReplaceAll(string2, " ", "")

	// string2 contains string1
	if len(joinedString1) > longNameLength && strings.Contains(joinedString2, joinedString1) {
		return levenshteinDistance(joinedString1, joinedString2) <= tolerance
	}

	// string1 contains string2
	if len(joinedString2) > longNameLength && strings.Contains(joinedString1, joinedString2) {
		return levenshteinDistance(joinedString2, joinedString1) <= tolerance
	}

	// Both words have the same amount of words
	if len(strings.Split(string1, " ")) == len(strings.Split(string2, " ")) {
		return levenshteinDistance(string1, string2) <= tolerance || levenshteinDistance(string2, string1) <= tolerance
	}

	// If it has different amount of words
	return false
}

// SanitizeString receives a string and eliminates special characters.
func SanitizeString(s string) (string, bool) {
	hasInvalidChar := false

	clean := strings.Map(func(r rune) rune {
		if !unicode.IsPrint(r) {
			hasInvalidChar = true
			return -1
		}
		return r
	}, s)

	return clean, hasInvalidChar
}

// levenshteinDistance returns a number between 0 and 1. 0 means both strings are equal,
// anything bigger represents the percentage of character that should be modified from s to get the string t
// IMPORTANT: both s and t should be non empty strings. This function assumes the length is at least 1
func levenshteinDistance(s, t string) float64 {
	// for all i and j, d[i,j] will hold the Levenshtein distance between
	// the first i characters of s and the first j characters of t
	// note that d has (m+1)*(n+1) values
	d := make([][]int, len(s))

	// source prefixes can be transformed into empty string by
	// dropping all characters
	for i := 0; i < len(s); i++ {
		row := make([]int, len(t))
		row[0] = i
		d[i] = row
	}

	// target prefixes can be reached from empty source prefix
	// by inserting every character
	for j := 0; j < len(t); j++ {
		d[0][j] = j
	}

	// Check len(s) > 0
	// If len(s) == 0, then the levenshtein distance will be 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			substitutionCost := 1
			if s[i] == t[j] {
				substitutionCost = 0
			}

			if i == 0 || j == 0 {
				d[i][j] = substitutionCost
			} else {
				d[i][j] = (int)(math.Min(
					math.Min(
						float64(d[i-1][j]+1),
						float64(d[i][j-1]+1)),
					float64(d[i-1][j-1]+substitutionCost)))
			}
		}
	}

	distance := float64(d[len(s)-1][len(t)-1]) / math.Min(float64(len(s)), float64(len(t)))
	return distance
}

func RemoveNotAlphaNumericValues(str string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(str, "")
}

func GetAllIndexPositions(in string, cr rune) []int {
	var indexArr []int
	for i, e := range []rune(in) {
		if e == cr {
			indexArr = append(indexArr, i)
			result := []rune(in)
			if cr != '?' {
				result[i] = '¿'
			}
			result[i] = '?'
			in = string(result)
		}
	}
	return indexArr
}

func ReplaceAtAllIndexs(in string, r rune, idxs []int) string {
	result := " "
	for i := 0; i < len(idxs); i++ {
		result = ReplaceAtIndex(in, r, idxs[i])
	}
	return result
}

func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

// Delete extra spaces in the string, when there are multiple spaces, only one space is reserved
func DeleteExtraSpace(s string) string {
	s1 := strings.Replace(s, " ", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}
