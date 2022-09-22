package json

import (
	"bytes"
	"encoding/json"

	"github.com/mono83/maybe"
)

var jsonNull = []byte("null")

func Wrap[T any](m maybe.Maybe[T]) Maybe[T] {
	return Maybe[T]{Maybe: m}
}

// Maybe in json package is a wrapper over standard maybe.Maybe
// providing JSON serialization support
type Maybe[T any] struct {
	maybe.Maybe[T]
}

// MarshalJSON returns JSON bytes representation of Maybe contents
func (m Maybe[T]) MarshalJSON() ([]byte, error) {
	if m.IsEmpty() {
		return jsonNull, nil
	}
	return json.Marshal(m.Value())
}

// UnmarshalJSON fills Maybe container with decoded JSON data
func (m *Maybe[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		m.Maybe = maybe.Nothing[T]()
	} else {
		value := m.Value()
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Maybe = maybe.Nilable(value)
	}
	return nil
}
