package server_test

import (
	"errors"
	"testing"

	"github.com/mrbenosborne/go-http-cache/server"
	"github.com/mrbenosborne/go-http-cache/store"

	"github.com/stretchr/testify/require"
)

func TestNewAPIError(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    int
		inputErr error
		want     error
	}{
		{
			name:     "no error",
			input:    int(store.NoError),
			inputErr: nil,
			want:     errors.New("unknown error"),
		},
		{
			name:     "no key specified",
			input:    int(store.NoKeySpecified),
			inputErr: nil,
			want:     server.NewAPIError(int(store.NoKeySpecified), nil),
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := server.NewAPIError(tc.input, tc.inputErr)
			require.Equal(t, tc.want, result)
		})
	}
}
