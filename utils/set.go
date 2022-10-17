package utils

import (
	"fmt"
	"strings"
	"sync"
)

// Set defines a Set type
type Set struct {
	mutex       *sync.Mutex
	internalMap map[string]bool
}

// NewSet returns a new Set
func NewSet() *Set {
	return &Set{
		mutex:       &sync.Mutex{},
		internalMap: make(map[string]bool),
	}
}

// NewSetFromSlice returns a new Set from a slice
func NewSetFromSlice(slice []string) *Set {
	set := NewSet()
	for _, item := range slice {
		set.Add(item)
	}

	return set
}

// Add adds an element to the set
func (set Set) Add(elem string) {
	set.lock()
	defer set.unlock()

	set.internalMap[elem] = true
}

// Delete deletes an element from the set
func (set Set) Delete(elem string) {
	set.lock()
	defer set.unlock()

	delete(set.internalMap, elem)
}

// Size returns the size of the set
func (set Set) Size() int {
	set.lock()
	defer set.unlock()

	return len(set.internalMap)
}

// Contains returns if the set contains the element
func (set Set) Contains(elem string) bool {
	set.lock()
	defer set.unlock()

	return set.internalMap[elem]
}

// Union returns the union between the two sets
func (set Set) Union(other *Set) *Set {
	set.lock()
	other.lock()
	defer set.unlock()
	defer other.unlock()

	union := NewSet()

	for elem := range set.internalMap {
		union.Add(elem)
	}

	for elem := range other.internalMap {
		union.Add(elem)
	}

	return union
}

// Intersect returns the intersection between the two sets
func (set Set) Intersect(other *Set) *Set {
	set.lock()
	other.lock()
	defer set.unlock()
	defer other.unlock()
	intersection := NewSet()

	if len(set.internalMap) < len(other.internalMap) {
		for elem := range set.internalMap {
			if other.internalMap[elem] {
				intersection.internalMap[elem] = true
			}
		}
	} else {
		for elem := range other.internalMap {
			if set.internalMap[elem] {
				intersection.internalMap[elem] = true
			}
		}
	}

	return intersection
}

// Difference returns the difference between the two sets
func (set Set) Difference(other *Set) *Set {
	set.lock()
	other.lock()
	defer set.unlock()
	defer other.unlock()

	difference := NewSet()

	for elem := range set.internalMap {
		if !other.internalMap[elem] {
			difference.internalMap[elem] = true
		}
	}

	return difference
}

func (set Set) String() string {
	slice := set.ToSlice()

	return fmt.Sprintf("Set{%s}", strings.Join(slice, ", "))
}

// ToSlice returns an slice with the elements of the set
func (set Set) ToSlice() []string {
	set.lock()
	defer set.unlock()

	keys := make([]string, 0, len(set.internalMap))
	for elem := range set.internalMap {
		keys = append(keys, elem)
	}

	return keys
}

func (set Set) lock() {
	set.mutex.Lock()
}

func (set Set) unlock() {
	set.mutex.Unlock()
}
