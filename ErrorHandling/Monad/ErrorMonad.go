// Package Monad provides an error monad for making the chaining of functions returning errors more developer friendly
package Monad

type unaryThrower[T, U any] func(input T) (U, error)

type withError[T any] struct {
	Value T
	Err   error
}

// Wrap convert generic value to a wrapped version usable in the monad
func Wrap[T any](value T) withError[T] {
	return withError[T]{Value: value, Err: nil}
}

// Unwrap unwraps the wrapped value to the value and error
func Unwrap[T any](we withError[T]) (interface{}, error) {
	return we.Value, we.Err
}

// Curry transform a binary function to a curried unary function (T, U) => (V, error) into (T) => (U) => (V, error)
// allows for conveniently using binary functions in the error monad
func Curry[T, U, V any](fn func(i T, j U) (V, error), i T) func(i U) (V, error) {
	return func(j U) (V, error) {
		return fn(i, j)
	}
}

// Bind operates fn on inner value of wrapped input if error equals nil, else return err and default value for U
func Bind[T, U any](we withError[T], fn unaryThrower[T, U]) withError[U] {
	if we.Err != nil {
		return withError[U]{Err: we.Err}
	}

	value, newErr := fn(we.Value)
	return withError[U]{Value: value, Err: newErr}
}
