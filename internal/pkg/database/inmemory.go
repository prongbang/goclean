package database

import (
	"errors"
	"sync"
)

type inMemoryStore struct {
	data sync.Map
}

func (s *inMemoryStore) Set(key string, value any) error {
	s.data.Store(key, value)
	return nil
}

func (s *inMemoryStore) Get(key string) (any, error) {
	val, ok := s.data.Load(key)
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (s *inMemoryStore) GetSlice(key string) ([]any, error) {
	val, ok := s.data.Load(key)
	if !ok {
		return nil, errors.New("key not found")
	}

	slice, ok := val.([]any)
	if !ok {
		return nil, errors.New("value is not a []any")
	}
	return slice, nil
}

func (s *inMemoryStore) Values() ([]any, error) {
	var all []any
	s.data.Range(func(_, value any) bool {
		all = append(all, value)
		return true
	})
	return all, nil
}

func NewInMemory() Store {
	return &inMemoryStore{}
}
