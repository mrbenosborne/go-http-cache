package server_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/mrbenosborne/go-http-cache/server"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	s := server.New()
	require.NotNil(t, s)
}

func TestReady(t *testing.T) {
	s := server.New()

	go s.Serve(context.TODO())
	<-s.Ready
	defer s.Close()

	resp, err := http.Get("http://localhost:8901/")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
