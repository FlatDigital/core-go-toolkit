package utils_test

import (
	"testing"

	utils "github.com/FlatDigital/core-go-toolkit/utils"
	"github.com/stretchr/testify/assert"
)

func Test_Contains(t *testing.T) {
	ass := assert.New(t)

	s := utils.NewSet()

	s.Add("Axel")
	s.Add("Marco")

	ass.True(s.Contains("Axel"))
	ass.False(s.Contains("Pablo"))

	s.Delete("Marco")
	ass.False(s.Contains("Marco"))
}

func Test_NewSetFromSlice(t *testing.T) {
	ass := assert.New(t)

	arrayString := []string{
		"Axel",
		"Pablo",
	}

	s := utils.NewSetFromSlice(arrayString)

	ass.True(s.Contains("Axel"))
	ass.False(s.Contains("Marco"))
}

func Test_Size(t *testing.T) {
	ass := assert.New(t)

	s := utils.NewSet()

	s.Add("Axel")
	s.Add("Marco")

	ass.Equal(2, s.Size())
}

func Test_Union(t *testing.T) {
	ass := assert.New(t)

	s1 := utils.NewSet()

	s1.Add("Axel")
	s1.Add("Marco")

	s2 := utils.NewSetFromSlice([]string{
		"Pablo",
	})

	union := s1.Union(s2)

	ass.Equal(3, union.Size())
}

func Test_Intersection(t *testing.T) {
	ass := assert.New(t)

	s1 := utils.NewSet()

	s1.Add("Axel")
	s1.Add("Marco")

	s2 := utils.NewSetFromSlice([]string{
		"Pablo",
		"Axel",
	})

	intersection := s1.Intersect(s2)

	ass.Equal(1, intersection.Size())
}

func Test_Intersection2(t *testing.T) {
	ass := assert.New(t)

	s1 := utils.NewSet()

	s1.Add("Axel")
	s1.Add("Marco")

	s2 := utils.NewSetFromSlice([]string{
		"Pablo",
		"Axel",
		"Amilcar",
	})

	intersection := s1.Intersect(s2)

	ass.Equal(1, intersection.Size())
}

func Test_Difference(t *testing.T) {
	ass := assert.New(t)

	s1 := utils.NewSet()

	s1.Add("Axel")
	s1.Add("Marco")

	s2 := utils.NewSetFromSlice([]string{
		"Pablo",
		"Axel",
	})

	difference := s1.Difference(s2)

	ass.Equal(1, difference.Size())
}
