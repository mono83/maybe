package maybe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybe_Get(t *testing.T) {
	iv, ok := Nothing[int]().Get()

	assert.Equal(t, 0, iv)
	assert.False(t, ok)

	iv, ok = Just(10).Get()
	assert.Equal(t, 10, iv)
	assert.True(t, ok)
}

func TestMaybe_Filter(t *testing.T) {
	x := Just(10)

	assert.True(t, x.Filter(func(i int) bool { return i > 5 }).IsPresent())
	assert.True(t, x.Filter(func(i int) bool { return i < 5 }).IsEmpty())
}

func TestMaybe_IfPresent(t *testing.T) {
	j := Just("Foo")
	e := Nothing[string]()

	j.IfPresent(nil) // No panic
	e.IfPresent(nil) // No panic

	e.IfPresent(func(string) { t.Fail() })

	cnt := 0
	j.IfPresent(func(x string) {
		if assert.Equal(t, "Foo", x) {
			cnt++
		}
	})
	assert.Equal(t, 1, cnt)
}

func TestMaybe_IfPresentOrElse(t *testing.T) {
	j := Just("Foo")
	e := Nothing[string]()

	j.IfPresentOrElse(nil, nil) // No panic
	e.IfPresentOrElse(nil, nil) // No panic

	cnt := 0
	j.IfPresentOrElse(func(s string) {
		assert.Equal(t, "Foo", s)
		cnt++
	}, func() { t.Fail() })
	assert.Equal(t, 1, cnt)

	e.IfPresentOrElse(
		func(string) { t.Fail() },
		func() { cnt++ },
	)
	assert.Equal(t, 2, cnt)
}

func TestMaybe_Or(t *testing.T) {
	assert.Equal(t, 22, Just(22).Or(func() Maybe[int] { return Just(333) }).Value())
	assert.Equal(t, 33, Nothing[int]().Or(func() Maybe[int] { return Just(33) }).Value())
	assert.Equal(t, 22, Just(22).Or(nil).Value())
	assert.True(t, Nothing[int]().Or(nil).IsEmpty())
}

func TestMaybe_OrElse(t *testing.T) {
	assert.Equal(t, "foo", Just("foo").OrElse("bar"))
	assert.Equal(t, "bar", Nothing[string]().OrElse("bar"))
}

func TestMaybe_OrElseGet(t *testing.T) {
	assert.Equal(t, "foo", Just("foo").OrElseGet(func() string { return "bar" }))
	assert.Equal(t, "bar", Nothing[string]().OrElseGet(func() string { return "bar" }))
	assert.Equal(t, "foo", Just("foo").OrElseGet(nil))

	assert.Panics(t, func() { Nothing[string]().OrElseGet(nil) })
}
