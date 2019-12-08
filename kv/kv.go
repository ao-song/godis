package kv

import (
	"sync"
)

type Store struct {
	KV  map[string]interface{}
	Mux sync.RWMutex
}

func (s *Store) Set(k string, v interface{}) {
	s.Mux.Lock()
	s.KV[k] = v
	s.Mux.Unlock()
}

func (s *Store) Get(k string) (interface{}, bool) {
	v, found := s.KV[k]
	return v, found
}

func (s *Store) Del(k string) {
	s.Mux.Lock()
	delete(s.KV, k)
	s.Mux.Unlock()
}
