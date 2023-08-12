package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/utils"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()
	assert.False(set.Has("test"))

	set.Add("test")
	assert.True(set.Has("test"))

}

func TestSize(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()
	assert.Equal(0, set.Size())

	set.Add("test")
	assert.Equal(1, set.Size())

	set.Add("test")
	assert.Equal(1, set.Size())

}

func TestAddMulti(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()
	assert.Equal(0, set.Size())

	set.AddMulti("test1", "test2", "test3")
	assert.Equal(3, set.Size())

}

func TestRemove(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()

	set.AddMulti("test1", "test2", "test3")
	assert.Equal(3, set.Size())

	set.Remove("test")
	assert.Equal(3, set.Size())

	set.Remove("test1")
	assert.Equal(2, set.Size())

}

func TestToSlice(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()

	expected := []string{"test1", "test2", "test3"}
	set.AddMulti(expected...)

	list := set.ToSlice()
	assert.Equal(3, len(list))

	for _, v := range list {
		assert.Contains(expected, v)
	}

}

func TestMarshalJSON(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()

	expected := []string{"test1", "test2", "test3"}
	set.AddMulti(expected...)

	bytes, err := set.MarshalJSON()
	assert.NoError(err)

	var list []string
	err = json.Unmarshal(bytes, &list)
	assert.NoError(err)

	assert.Equal(3, len(list))

	for _, v := range list {
		assert.Contains(expected, v)
	}

}

func TestUnmarshalJSON(t *testing.T) {

	assert := assert.New(t)

	set := utils.NewSet[string]()

	expected := []string{"test1", "test2", "test3"}

	bytes, err := json.Marshal(expected)
	assert.NoError(err)

	err = set.UnmarshalJSON(bytes)
	assert.NoError(err)

	assert.Equal(3, set.Size())

	for _, v := range expected {
		assert.True(set.Has(v))
	}

}

func TestSetUnion(t *testing.T) {
	assert := assert.New(t)

	set1 := utils.NewSetFromSlice([]int{1, 2, 3})
	set2 := utils.NewSetFromSlice([]int{3, 4, 5})

	result := set1.Union(set2).ToSlice()
	expected := []int{1, 2, 3, 4, 5}

	assert.ElementsMatch(expected, result)
}

func TestSetIntersect(t *testing.T) {
	assert := assert.New(t)

	set1 := utils.NewSetFromSlice([]int{1, 2, 3})
	set2 := utils.NewSetFromSlice([]int{3, 4, 5})

	result := set1.Intersect(set2).ToSlice()
	expected := []int{3}

	assert.ElementsMatch(expected, result)
}

func TestSetDifference(t *testing.T) {
	assert := assert.New(t)

	set1 := utils.NewSetFromSlice([]int{1, 2, 3})
	set2 := utils.NewSetFromSlice([]int{3, 4, 5})

	result := set1.Difference(set2).ToSlice()
	expected := []int{1, 2}

	assert.ElementsMatch(expected, result)
}

func TestSetString(t *testing.T) {
	assert := assert.New(t)

	set := utils.NewSetFromSlice([]string{"apple", "banana", "orange"})

	result := set.String()
	expected := "apple, banana, orange"

	assert.Equal(expected, result)
}

func TestSetStringWithEmptySet(t *testing.T) {
	assert := assert.New(t)

	set := utils.NewSet[string]() // Crear un conjunto vacío de strings

	result := set.String()
	expected := ""

	assert.Equal(expected, result)
}

func TestSetFunctions(t *testing.T) {
	assert := assert.New(t)

	set := utils.NewSet[string]()

	assert.False(set.Has("test"))

	set.Add("test")
	assert.True(set.Has("test"))

	assert.Equal(1, set.Size())

	set.Add("test")
	assert.Equal(1, set.Size())

	set.AddMulti("test1", "test2", "test3")
	assert.Equal(4, set.Size())

	set.Remove("test")
	assert.Equal(3, set.Size())

	set.Remove("test1")
	assert.Equal(2, set.Size())

	expected := []string{"test2", "test3"}
	list := set.ToSlice()
	assert.Equal(2, len(list))
	assert.ElementsMatch(expected, list)

	bytes, err := set.MarshalJSON()
	assert.NoError(err)

	var jsonList []string
	err = json.Unmarshal(bytes, &jsonList)
	assert.NoError(err)

	assert.ElementsMatch(expected, jsonList)

	err = set.UnmarshalJSON(bytes)
	assert.NoError(err)

	assert.Equal(2, set.Size())
	for _, v := range expected {
		assert.True(set.Has(v))
	}
}
