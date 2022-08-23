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

func TestMaybe_MapToNilMapper(t *testing.T) {
	assert.True(t, Just(15).MapToInt16(nil).IsEmpty())
}

func TestMaybe_MapTo(t *testing.T) {
	assert.Equal(t, "8", Just(8).MapToAny(func(x int) any { return strconv.Itoa(x) }).Value())

	assert.Equal(t, 8, Just(uint(8)).MapToInt(func(x uint) int { return int(x) }).Value())
	assert.Equal(t, int64(8), Just(8).MapToInt64(func(x int) int64 { return int64(x) }).Value())
	assert.Equal(t, int32(8), Just(8).MapToInt32(func(x int) int32 { return int32(x) }).Value())
	assert.Equal(t, int16(8), Just(8).MapToInt16(func(x int) int16 { return int16(x) }).Value())
	assert.Equal(t, uint(8), Just(8).MapToUint(func(x int) uint { return uint(x) }).Value())
	assert.Equal(t, uint16(8), Just(8).MapToUint16(func(x int) uint16 { return uint16(x) }).Value())
	assert.Equal(t, uint32(8), Just(8).MapToUint32(func(x int) uint32 { return uint32(x) }).Value())
	assert.Equal(t, uint64(8), Just(8).MapToUint64(func(x int) uint64 { return uint64(x) }).Value())
	assert.Equal(t, float32(8), Just(8).MapToFloat32(func(x int) float32 { return float32(x) }).Value())
	assert.Equal(t, float64(8), Just(8).MapToFloat64(func(x int) float64 { return float64(x) }).Value())
	assert.Equal(t, "8", Just(8).MapToString(strconv.Itoa).Value())
}
