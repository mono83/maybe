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

type Maybe[T any] struct {
	maybe.Maybe[T]
}

func (m Maybe[T]) MarshalJSON() ([]byte, error) {
	if m.IsEmpty() {
		return jsonNull, nil
	}
	return json.Marshal(m.Value())
}

func (m *Maybe[T]) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, jsonNull) == 0 {
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
