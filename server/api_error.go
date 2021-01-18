package server

import (
	"errors"
	"net/http"

	"github.com/mrbenosborne/go-http-cache/store"
)

// APIError a type of error that includes a
// friendly error message and a custom HTTP Status
// code.
type APIError struct {
	err      error
	Message  string `json:"error"`
	HTTPCode int    `json:"-"`
}

// Error return the error
func (a APIError) Error() string {
	return a.err.Error()
}

// NewAPIError return a new type of APIError.
func NewAPIError(internalErrorCode int, err error) error {
	friendlyErrors := getFriendlyErrors()
	if e, ok := friendlyErrors[internalErrorCode]; ok {
		if err != nil {
			e.err = err
			return e
		}

		e.err = errors.New(e.Message)
		return e
	}

	if err != nil {
		return err
	}
	return errors.New("unknown error")
}

func getFriendlyErrors() map[int]APIError {
	return map[int]APIError{
		int(store.NoKeySpecified): {
			Message:  "no key specified",
			HTTPCode: http.StatusBadRequest,
		},
	}
}
