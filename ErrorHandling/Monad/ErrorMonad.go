// Package Monad provides an error monad for making the chaining of functions returning errors more developer friendly
package Monad

type unaryThrower[T, U any] func(input T) (U, error)

type WithError[T any] interface {
	Value() T
	Error() error
	private()
}

type withError[T any] struct {
	value T
	err   error
}

func (w withError[T]) Value() T {
	return w.value
}

func (w withError[T]) Error() error {
	return w.err
}

func (w withError[T]) private() {}

// Wrap convert generic value to a wrapped version usable in the monad
func Wrap[T any](value T) WithError[T] {
	return withError[T]{value: value, err: nil}
}

// Unwrap unwraps the wrapped value to the value and error
func Unwrap[T any](we WithError[T]) (interface{}, error) {
	return we.Value(), we.Error()
}

// Curry transform a binary function to a curried unary function (T, U) => (V, error) into (T) => (U) => (V, error)
// allows for conveniently using binary functions in the error monad
func Curry[T, U, V any](i T, fn func(i T, j U) (V, error)) func(i U) (V, error) {
	return func(j U) (V, error) {
		return fn(i, j)
	}
}

// Bind operates fn on inner value of wrapped input if error equals nil, else return err and default value for U
func Bind[T, U any](we WithError[T], fn unaryThrower[T, U]) WithError[U] {
	if we.Error() != nil {
		return withError[U]{err: we.Error()}
	}

	value, newErr := fn(we.Value())
	return withError[U]{value: value, err: newErr}
}
