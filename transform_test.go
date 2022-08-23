package maybe

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	x := Just(42)

	// Int to string
	assert.Equal(t, "42", Map(x, strconv.Itoa).Value())

	// Int to nil
	assert.True(t, Map(x, func(i int) error { return nil }).IsEmpty())

	// Nil mapper
	var mapper func(int) int
	assert.True(t, Map(x, mapper).IsEmpty())
}
