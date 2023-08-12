package utils

import (
	"database/sql/driver"
	"encoding/json"
)

type Set[T comparable] struct {
	list map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	s := &Set[T]{}
	s.list = make(map[T]struct{})
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
