package store_test

import (
	"testing"

	"github.com/mrbenosborne/go-http-cache/store"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	s := store.New()
	require.NotNil(t, s)
}
