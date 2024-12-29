package types

type Optional[T comparable] struct {
	Value T
	Valid bool
}

func OptionalValue[T comparable](value T) Optional[T] {
	return Optional[T]{value, true}
}

func OptionalEmpty[T comparable]() Optional[T] {
	return Optional[T]{Valid: false}
}

func OptionalZeroed[T comparable](value T) Optional[T] {
	var zeroed T

	return Optional[T]{Value: value, Valid: value != zeroed}
}

func OptionalPointer[T comparable](value *T) Optional[T] {
	if value == nil {
		return OptionalEmpty[T]()
	}

	return OptionalValue(*value)
}
