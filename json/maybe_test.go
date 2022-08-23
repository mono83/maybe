package json

import (
	"encoding/json"
	"github.com/mono83/maybe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonMarshal(t *testing.T) {
	if bts, err := json.Marshal(Wrap(maybe.Just("Hello, world"))); assert.NoError(t, err) {
		assert.Equal(t, "\"Hello, world\"", string(bts))
	}
	if bts, err := json.Marshal(Wrap(maybe.Just(12345))); assert.NoError(t, err) {
		assert.Equal(t, "12345", string(bts))
	}
	if bts, err := json.Marshal(Wrap(maybe.Nothing[int]())); assert.NoError(t, err) {
		assert.Equal(t, "null", string(bts))
	}
}

func TestJsonUnmarshal(t *testing.T) {
	var x Maybe[string]
	if err := json.Unmarshal([]byte("null"), &x); assert.NoError(t, err) {
		assert.True(t, x.IsEmpty())
	}
	if err := json.Unmarshal([]byte("\"foo\""), &x); assert.NoError(t, err) && assert.True(t, x.IsPresent()) {
		assert.Equal(t, "foo", x.Value())
	}
}

type someStruct struct {
	ID       Maybe[int]
	ParentID Maybe[int]
	Name     Maybe[string]
}

func TestSomeStruct(t *testing.T) {
	x := someStruct{
		ID:       Wrap(maybe.Just(5243)),
		ParentID: Wrap(maybe.Nothing[int]()),
		Name:     Wrap(maybe.Nothing[string]()),
	}
	if bts, err := json.Marshal(x); assert.NoError(t, err) {
		assert.Equal(t, "{\"ID\":5243,\"ParentID\":null,\"Name\":null}", string(bts))
	}

	var y someStruct
	if err := json.Unmarshal([]byte("{\"ID\":5243,\"ParentID\":null,\"Name\":null}"), &y); assert.NoError(t, err) {
		assert.Equal(t, x, y)
	}
}
