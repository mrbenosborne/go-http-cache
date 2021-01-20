package server_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mrbenosborne/go-http-cache/server"
	"github.com/mrbenosborne/go-http-cache/store"

	"github.com/stretchr/testify/require"
)

func TestJSONHandler(t *testing.T) {
	for _, tc := range []struct {
		name           string
		handler        server.JSONHandler
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "standard error",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, errors.New("standard error")
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"error":"standard error"}`,
		},
		{
			name: "api error with zero http code and nil error",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, server.NewAPIError(int(store.NoError), nil)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"error":"unknown error"}`,
		},
		{
			name: "api error with zero http code and error",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, server.NewAPIError(int(store.NoError), errors.New("foo error"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       `{"error":"foo error"}`,
		},
		{
			name: "api error with http code",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, server.NewAPIError(int(store.NoKeySpecified), nil)
			},
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"error":"no key specified"}`,
		},
		{
			name: "no content and no errors",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return nil, nil
			},
			wantStatusCode: http.StatusNoContent,
			wantBody:       "",
		},
		{
			name: "content and no errors",
			handler: func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
				return struct {
					Data string `json:"data"`
				}{Data: "Hello, World!"}, nil
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `{"data":"Hello, World!"}`,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "", nil)
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tc.handler.Func())
			handler.ServeHTTP(rr, req)

			require.Equal(t, tc.wantStatusCode, rr.Code)
			require.Equal(t, tc.wantBody, strings.Trim(rr.Body.String(), "\n"))
		})
	}
}
