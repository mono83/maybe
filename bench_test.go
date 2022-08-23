package maybe

import "testing"

func BenchmarkNothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nothing[int]()
	}
}

func BenchmarkJust(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Just(i)
	}
}

func BenchmarkPtr(b *testing.B) {
	x := 4
	y := &x
	for i := 0; i < b.N; i++ {
		Ptr(y)
	}
}

func BenchmarkPtrNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ptr[string](nil)
	}
}

func BenchmarkNilableInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nilable(i)
	}
}

func BenchmarkNilableNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Nilable[error](nil)
	}
}
