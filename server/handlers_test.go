package server_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mrbenosborne/go-http-cache/server"

	"github.com/stretchr/testify/require"
)

func TestStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.StatusHandler())
	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	require.Equal(t, "", rr.Body.String())
}

func TestSetHandler(t *testing.T) {
	r := server.New()
	ts := httptest.NewServer(r.GetRouter())
	defer ts.Close()

	for _, tc := range []struct {
		name           string
		key            string
		value          string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "empty key / empty value",
			key:            " ",
			value:          "",
			wantStatusCode: http.StatusBadRequest,
			wantBody:       `{"error":"no key specified"}`,
		},
		{
			name:           "empty value",
			key:            "foo",
			value:          "",
			wantStatusCode: http.StatusOK,
			wantBody:       `{"acknowledged":true}`,
		},
		{
			name:           "with value",
			key:            "foo",
			value:          "bar",
			wantStatusCode: http.StatusOK,
			wantBody:       `{"acknowledged":true}`,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", ts.URL, tc.key), strings.NewReader(tc.value))
			require.NoError(t, err)
			require.NotNil(t, req)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.NotNil(t, resp)

			respBody, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, tc.wantStatusCode, resp.StatusCode)
			require.Equal(t, tc.wantBody, strings.Trim(string(respBody), "\n"))
		})
	}
}

func TestGetHandler(t *testing.T) {
	r := server.New()
	ts := httptest.NewServer(r.GetRouter())
	defer ts.Close()

	for _, tc := range []struct {
		name           string
		key            string
		value          string
		wantStatusCode int
		want           string
	}{
		{
			name:           "empty value",
			key:            "foo",
			value:          "",
			wantStatusCode: http.StatusOK,
			want:           "",
		},
		{
			name:           "with value",
			key:            "foo",
			value:          "bar",
			wantStatusCode: http.StatusOK,
			want:           "bar",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// set the value
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", ts.URL, tc.key), strings.NewReader(tc.value))
			require.NoError(t, err)
			require.NotNil(t, req)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.NotNil(t, resp)

			// get the value
			req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", ts.URL, tc.key), nil)
			require.NoError(t, err)
			require.NotNil(t, req)

			resp, err = http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.NotNil(t, resp)

			respBody, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, tc.wantStatusCode, resp.StatusCode)
			require.Equal(t, tc.want, strings.Trim(string(respBody), "\n"))
		})
	}
}
