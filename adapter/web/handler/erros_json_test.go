package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"Message":"Hello json"}`), result)
}
