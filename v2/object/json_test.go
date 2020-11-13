package object_test

import (
	"testing"

	"github.com/nspcc-dev/neofs-api-go/v2/object"
	"github.com/stretchr/testify/require"
)

func TestShortHeaderJSON(t *testing.T) {
	h := generateShortHeader("id")

	data, err := h.MarshalJSON()
	require.NoError(t, err)

	h2 := new(object.ShortHeader)
	require.NoError(t, h2.UnmarshalJSON(data))

	require.Equal(t, h, h2)
}

func TestAttributeJSON(t *testing.T) {
	a := generateAttribute("key", "value")

	data, err := a.MarshalJSON()
	require.NoError(t, err)

	a2 := new(object.Attribute)
	require.NoError(t, a2.UnmarshalJSON(data))

	require.Equal(t, a, a2)
}