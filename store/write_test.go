package store_test

import (
	"testing"

	"github.com/mrbenosborne/go-http-cache/store"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	for _, tc := range []struct {
		name    string
		key     string
		value   []byte
		want    []byte
		wantErr store.WriteError
	}{
		{
			name:    "empty key and nil value specified",
			key:     "",
			value:   nil,
			want:    nil,
			wantErr: store.NoKeySpecified,
		},
		{
			name:    "empty key specified",
			key:     "",
			value:   []byte("bar"),
			want:    nil,
			wantErr: store.NoKeySpecified,
		},
		{
			name:    "nil value specified",
			key:     "foo",
			value:   nil,
			want:    []byte(nil),
			wantErr: store.NoError,
		},
		{
			name:    "key and value",
			key:     "foo",
			value:   []byte("bar"),
			want:    []byte("bar"),
			wantErr: store.NoError,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			s := store.New()

			writeErr := s.Set(tc.key, tc.value)
			require.Equal(t, tc.wantErr, writeErr)

			if writeErr == store.NoError {
				require.Equal(t, tc.want, s.Get(tc.key))
			}
		})
	}
}
