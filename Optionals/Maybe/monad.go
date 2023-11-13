package Maybe

// todo goal is to make the maybe monad not definable outside this package, not even with a default type value

type mapper[T, U any] func(input T) U

type Maybe[T any] interface {
	Value() T
	HasValue() bool
	private()
}

type monad[T any] struct {
	hasValue bool
	value    T
}

func (m monad[T]) Value() T {
	return m.value
}

func (m monad[T]) HasValue() bool {
	return m.hasValue
}

func (m monad[T]) private() {}

func Some[T any](value T) Maybe[T] {
	return monad[T]{
		hasValue: true,
		value:    value,
	}
}

func None[T any]() Maybe[T] {
	return monad[T]{
		hasValue: false,
	}
}

func Map[T, U any](m Maybe[T], fn mapper[T, U]) Maybe[U] {
	if m.HasValue() {
		return monad[U]{
			hasValue: true,
			value:    fn(m.Value()),
		}
	}

	return monad[U]{
		hasValue: false,
	}
}

// map is used when you have a function T => U
// flatmap is used when you have a function T => Maybe<U>
