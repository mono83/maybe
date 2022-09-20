package maybe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {
	s := Nothing[string]()

	assert.True(t, s.IsEmpty())
	assert.False(t, s.IsPresent())
	assert.Equal(t, "", s.Value())

	i := Nothing[int]()
	assert.True(t, i.IsEmpty())
	assert.False(t, i.IsPresent())
	assert.Equal(t, 0, i.Value())

	e := Nothing[error]()
	assert.True(t, e.IsEmpty())
	assert.False(t, e.IsPresent())
	assert.Nil(t, e.Value())
}

func TestJust(t *testing.T) {
	s := Just("Hello")

	assert.False(t, s.IsEmpty())
	assert.True(t, s.IsPresent())
	assert.Equal(t, "Hello", s.Value())
}

func TestPtr(t *testing.T) {
	p := Ptr[string](nil)
	assert.True(t, p.IsEmpty())

	rs := "Hello"
	p = Ptr(&rs)
	assert.False(t, p.IsEmpty())
	assert.Equal(t, "Hello", p.Value())
}
