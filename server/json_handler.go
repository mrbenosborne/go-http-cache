package server

import (
	"encoding/json"
	"net/http"
)

// ErrResp a standard error response.
type ErrResp struct {
	Error string `json:"error"`
}

// JSONHandler a handler for JSON
// responses.
type JSONHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// Func the HandlerFunc
func (j JSONHandler) Func() http.HandlerFunc {
	return j.ServeHTTP
}

// ServeHTTP executes the handler to receive the response and
// error.
func (j JSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := j(w, r)
	if err != nil {
		j.SendErrorResponse(w, r, err)
		return
	}

	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		j.SendErrorResponse(w, r, err)
	}
}

// SendErrorResponse send a JSON error response.
func (j JSONHandler) SendErrorResponse(w http.ResponseWriter, r *http.Request, apiError error) {
	switch e := apiError.(type) {
	case APIError:
		if e.HTTPCode > 0 {
			w.WriteHeader(e.HTTPCode)
		}

		err := json.NewEncoder(w).Encode(e)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusInternalServerError)

		if apiError != nil {
			resp := ErrResp{Error: apiError.Error()}
			_ = json.NewEncoder(w).Encode(resp)
		}
	}
}
