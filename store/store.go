package store

import "sync"

// Store a cache store that holds
// key value pairs.
type Store struct {
	mu   sync.RWMutex
	data map[string][]byte
}

// New initialize a new store.
func New() *Store {
	return &Store{
		data: make(map[string][]byte),
	}
}
