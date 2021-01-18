package store

// Get get an item from the store, if the
// key does not exist "nil" will be returned.
func (s *Store) Get(key string) []byte {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data[key]
}
