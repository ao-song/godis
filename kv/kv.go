package kv

import (
	"sync"
)

type Store struct {
	kv  map[string]interface{}
	mux sync.RWMutex
}

func (s *Store) Set(k string, v interface{}) {
	s.mux.Lock()
	s.kv[k] = v
	s.mux.Unlock()
}

func (s *Store) Get(k string) (interface{}, bool) {
	v, found := s.kv[k]
	return v, found
}

func (s *Store) Del(k string) {
	s.mux.Lock()
	delete(s.kv, k)
	s.mux.Unlock()
}
