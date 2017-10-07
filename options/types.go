package options

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// Bytes is a []byte array that outputs hex
type Bytes []byte

func (b Bytes) MarshalJSON() ([]byte, error) {
	s := strings.ToUpper(hex.EncodeToString(b))
	return json.Marshal(s)
}

func (b *Bytes) UnmarshalJSON(src []byte) (err error) {
	var s string
	err = json.Unmarshal(src, &s)
	if err != nil {
		return errors.Wrap(err, "parse string")
	}
	// and interpret that string as hex
	*b, err = hex.DecodeString(s)
	return err
}

func (b *Bytes) Size() int {
	if b == nil {
		return 0
	}
	return len(*b)
}

// Marshal for protobuf
func (b Bytes) Marshal() ([]byte, error) {
	return b, nil
}

// MarshalTo for protobuf
func (b Bytes) MarshalTo(data []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	copy(data, b)
	return len(b), nil
}

// Unmarshal for protobuf
func (b *Bytes) Unmarshal(data []byte) error {
	if len(data) == 0 {
		b = nil
		return nil
	}
	*b = data
	return nil
}

func (b Bytes) Equal(other Bytes) bool {
	return bytes.Equal(b, other)
}
