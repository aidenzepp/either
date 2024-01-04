package either

import "fmt"

const (
	lhs = false
	rhs = true
)

type Either[L, R any] struct {
	lhs L
	rhs R
	side bool
}

func NewEitherLhs[L, R any](value L) Either[L, R] {
	var zero R
	return Either[L, R]{value, zero, lhs}
}

func NewEitherRhs[L, R any](value R) Either[L, R] {
	var zero L
	return Either[L, R]{zero, value, rhs}
}

func (e Either[L, R]) IsLhs() bool {
	return e.side == lhs
}

func (e Either[L, R]) IsRhs() bool {
	return e.side == rhs
}

func (e Either[L, R]) Lhs() Option[L] {
	if e.IsLhs() {
		return NewOption(e.lhs)
	}

	return NilOption[L]()
}

func (e Either[L, R]) Rhs() Option[R] {
	if e.IsRhs() {
		return NewOption(e.rhs)
	}
	
	return NilOption[R]()
}

func (e Either[L, R]) Flip() Either[R, L] {
	if e.IsLhs() {
		return NewEitherRhs[R, L](e.lhs)
	}

	return NewEitherLhs[R, L](e.rhs)
}

func (e Either[L, R]) LhsOr(value L) L {
	if e.IsLhs() {
		return e.lhs
	}

	return value
}

func (e Either[L, R]) LhsOrDefault() L {
	if e.IsLhs() {
		return e.lhs
	}

	var zero L
	return zero
}

func (e Either[L, R]) LhsOrElse(f func(R) L) L {
	if e.IsLhs() {
		return e.lhs
	}

	return f(e.rhs)
}

func (e Either[L, R]) LhsOrError(err error) Result[L] {
	if e.IsLhs() {
		return NewResult(e.lhs)
	}

	return ErrResult[L](err)
}

func (e Either[L, R]) RhsOr(value R) R {
	if e.IsRhs() {
		return e.rhs
	}

	return value
}

func (e Either[L, R]) RhsOrDefault() R {
	if e.IsRhs() {
		return e.rhs
	}

	var zero R
	return zero
}

func (e Either[L, R]) RhsOrElse(f func(L) R) R {
	if e.IsRhs() {
		return e.rhs
	}

	return f(e.lhs)
}

func (e Either[L, R]) RhsOrError(err error) Result[R] {
	if e.IsRhs() {
		return NewResult(e.rhs)
	}

	return ErrResult[R](err)
}

func (e Either[L, R]) UnwrapLhs() L {
	return e.ExpectLhs("called `Either.UnwrapLhs()` on a right-hand side value")
}

func (e Either[L, R]) UnwrapRhs() R {
	return e.ExpectRhs("called `Either.UnwrapRhs()` on a left-hand side value")
}

func (e Either[L, R]) ExpectLhs(message string) L {
	if e.IsLhs() {
		return e.lhs
	}

	panic(fmt.Sprintf("%s: %v", message, e.rhs))
}

func (e Either[L, R]) ExpectRhs(message string) R {
	if e.IsRhs() {
		return e.rhs
	}

	panic(fmt.Sprintf("%s: %v", message, e.lhs))
}


