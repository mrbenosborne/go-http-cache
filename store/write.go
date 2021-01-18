package store

import "strings"

// WriteError an integer indicating a type
// of error.
type WriteError int

const (
	// NoError no error has occurred, 0.
	NoError WriteError = iota

	// NoKeySpecified a key was not specified.
	NoKeySpecified
)

// Set set a key value pair in the
// store.
func (s *Store) Set(key string, value []byte) WriteError {
	key = strings.ReplaceAll(key, " ", "")
	if key == "" {
		return NoKeySpecified
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value

	return NoError
}
