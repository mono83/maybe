package maybe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilable(t *testing.T) {
	assert.True(t, Nilable[error](nil).IsEmpty())
	assert.True(t, Nilable[*int](nil).IsEmpty())
	assert.True(t, Nilable[[]int](nil).IsEmpty())
	assert.True(t, Nilable[map[int]int](nil).IsEmpty())
	assert.True(t, Nilable([]int{}).IsEmpty())
	assert.True(t, Nilable(map[int]int{}).IsEmpty())

	i := 10
	assert.True(t, Nilable(errors.New("foo")).IsPresent())
	assert.True(t, Nilable(&i).IsPresent())
	assert.True(t, Nilable([]int{22}).IsPresent())
	assert.True(t, Nilable(map[int]int{5: 6}).IsPresent())
}

func TestFilterNonNil(t *testing.T) {
	assert.True(t, Nilable[error](nil).FilterNotNil().IsEmpty())
	assert.False(t, Just[error](nil).IsEmpty())
	assert.True(t, Just[error](nil).FilterNotNil().IsEmpty())
}
