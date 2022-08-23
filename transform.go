package maybe

// Map constructs new Maybe container built using
// mapper function.
func Map[T any, R any](from Maybe[T], mapper func(T) R) Maybe[R] {
	if from.IsEmpty() || mapper == nil {
		return Nothing[R]()
	}

	return Nilable(mapper(from.Value()))
}
