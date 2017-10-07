package options

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytes(t *testing.T) {
	input := []byte(`"DEADBEEF1234"`)
	expected := Bytes{0xDE, 0xAD, 0xBE, 0xEF, 0x12, 0x34}

	var b Bytes
	err := json.Unmarshal(input, &b)
	require.Nil(t, err)
	require.Equal(t, expected, b)

	out, err := json.Marshal(b)
	require.Nil(t, err)
	require.Equal(t, input, out)
}
