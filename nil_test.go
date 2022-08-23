package maybe

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilable(t *testing.T) {
	assert.True(t, Nilable[error](nil).IsEmpty())
	assert.True(t, Nilable[*int](nil).IsEmpty())
	assert.True(t, Nilable[[]int](nil).IsEmpty())
	assert.True(t, Nilable[map[int]int](nil).IsEmpty())
	assert.True(t, Nilable[[]int]([]int{}).IsEmpty())
	assert.True(t, Nilable[map[int]int](map[int]int{}).IsEmpty())

	i := 10
	assert.True(t, Nilable[error](errors.New("foo")).IsPresent())
	assert.True(t, Nilable[*int](&i).IsPresent())
	assert.True(t, Nilable[[]int]([]int{22}).IsPresent())
	assert.True(t, Nilable[map[int]int](map[int]int{5: 6}).IsPresent())
}

func TestFilterNonNil(t *testing.T) {
	assert.True(t, Nilable[error](nil).FilterNotNil().IsEmpty())
	assert.False(t, Just[error](nil).IsEmpty())
	assert.True(t, Just[error](nil).FilterNotNil().IsEmpty())
}
