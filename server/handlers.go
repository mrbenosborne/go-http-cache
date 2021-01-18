package server

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/mrbenosborne/go-http-cache/store"
)

// StatusHandler a handler for a status endpoint,
// simply returns a "200 OK" response.
func StatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// SetHandler set a key/value pair in
// the store.
func SetHandler(s *store.Store) JSONHandler {
	return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		defer r.Body.Close()

		key := chi.URLParam(r, "key")
		value, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, NewAPIError(int(store.NoError), errors.New("failed to retrieve request body"))
		}

		writeErr := s.Set(key, value)
		if writeErr != store.NoError {
			return nil, NewAPIError(int(writeErr), nil)
		}
		return Acknowledged{Acknowledged: true}, nil
	}
}

// GetHandler get a value by the key.
func GetHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		key := chi.URLParam(r, "key")

		w.Write(s.Get(key))
	}
}
