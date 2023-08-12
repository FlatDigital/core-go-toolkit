package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	list map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
	return s
}

func NewSetFromSlice[T comparable](list []T) *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
	s.AddMulti(list...)
	return s
}

func (s *Set[T]) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Set[T]) Has(t T) bool {
	_, ok := s.list[t]
	return ok
}

func (s *Set[T]) Add(t T) {
	s.list[t] = struct{}{}
}

func (s *Set[T]) Size() int {
	return len(s.list)
}

func (s *Set[T]) AddMulti(list ...T) {
	for _, v := range list {
		s.Add(v)
	}
}

func (s *Set[T]) Remove(t T) {
	delete(s.list, t)
}

func (s *Set[T]) ToSlice() []T {

	list := make([]T, 0, len(s.list))

	for k := range s.list {
		list = append(list, k)
	}
	return list
}

func (s *Set[T]) MarshalJSON() ([]byte, error) {

	array := s.ToSlice()

	return json.Marshal(array)
}

func (s *Set[T]) UnmarshalJSON(bytes []byte) error {
	var array []T

	err := json.Unmarshal(bytes, &array)
	if err != nil {
		return err
	}

	s.list = make(map[T]struct{})
	s.AddMulti(array...)

	return nil
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	unionSet := NewSet[T]()
	unionSet.AddMulti(s.ToSlice()...)
	unionSet.AddMulti(other.ToSlice()...)
	return unionSet
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	intersectionSet := NewSet[T]()
	for k := range s.list {
		if other.Has(k) {
			intersectionSet.Add(k)
		}
	}
	return intersectionSet
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	differenceSet := NewSet[T]()
	for k := range s.list {
		if !other.Has(k) {
			differenceSet.Add(k)
		}
	}
	return differenceSet
}

func (s *Set[T]) String() string {
	list := s.ToSlice()
	stringSlice := make([]string, len(list))
	for i, v := range list {
		stringSlice[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(stringSlice, ", ")
}
