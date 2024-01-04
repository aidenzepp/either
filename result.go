package either

import "fmt"

type Result[V any] struct {
	value V
	err error
}

func NewResult[V any](value V) Result[V] {
	return Result[V]{value, nil}
}

func ErrResult[V any](err error) Result[V] {
	var zero V
	return Result[V]{zero, err}
}

func (r Result[V]) IsValue() bool {
	return r.err == nil
}

func (r Result[V]) IsValueAnd(predicate  func(V) bool) bool {
	return r.IsValue() && predicate(r.value)
}

func (r Result[V]) IsError() bool {
	return r.err != nil
}

func (r Result[V]) IsErrorAnd(predicate func(error) bool) bool {
	return r.IsError() && predicate(r.err)
}

func (r Result[V]) ValueOption() Option[V] {
	if r.IsValue() {
		return NewOption(r.value)
	}

	return NilOption[V]()
}

func (r Result[V]) ErrorOption() Option[error] {
	if r.IsError() {
		return NewOption(r.err)
	}

	return NilOption[error]()
}

func (r Result[V]) Expect(message string) V {
	if r.IsValue() {
		return r.value
	}

	panic(message + ": " + r.err.Error())
}

func (r Result[V]) Unwrap() V {
	return r.Expect("called `Result.Unwrap()` on an error value")
}

func (r Result[V]) UnwrapOrDefault() V {
	if r.IsValue() {
		return r.value
	}

	var zero V
	return zero
}

func (r Result[V]) ExpectError(message string) error {
	if r.IsError() {
		return r.err
	}

	panic(fmt.Sprintf("%s: %v", message, r.value))
}

func (r Result[V]) UnwrapError() error {
	return r.ExpectError("called `Result.UnwrapError` on an okay value")
}

func (r Result[V]) And(other Result[V]) Result[V] {
	if r.IsValue() {
		return other
	}

	return r
}

func (r Result[V]) AndThen(f func(V) Result[V]) Result[V] {
	if r.IsValue() {
		return f(r.value)
	}

	return r
}

func (r Result[V]) Or(other Result[V]) Result[V] {
	if r.IsError() {
		return other
	}

	return r
}

func (r Result[V]) OrElse(f func(error) Result[V]) Result[V] {
	if r.IsError() {
		return f(r.err)
	}

	return r
}

func (r Result[V]) UnwrapOr(value V) V {
	if r.IsValue() {
		return r.value
	}

	return value
}

func (r Result[V]) UnwrapUnchecked() V {
	return r.value
}

func (r Result[V]) UnwrapErrorUnchecked() error {
	return r.err
}

