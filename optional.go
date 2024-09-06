package gateapi

type Optional[T any] struct {
	isSet bool
	value T
}

func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{
		true,
		value,
	}
}

func EmptyOptional[T any]() Optional[T] {
	return Optional[T]{
		isSet: false,
	}
}

func (i Optional[T]) IsSet() bool {
	return i.isSet
}

func (i Optional[T]) Value() T {
	return i.value
}

func (i Optional[T]) Default(defaultValue T) T {
	if i.isSet {
		return i.value
	}
	return defaultValue
}
