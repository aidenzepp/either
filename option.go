package either

type Option[V any] struct {
	value V
	empty bool
}

func NewOption[V any](value V) Option[V] {
	return Option[V]{value, false}
}

func NilOption[V any]() Option[V] {
	var zero V
	return Option[V]{zero, true}
}

func (o Option[V]) IsValue() bool {
	return !o.empty
}

func (o Option[V]) IsValueAnd(predicate func(V) bool) bool {
	return o.IsValue() && predicate(o.value)
}

func (o Option[V]) IsEmpty() bool {
	return o.empty
}

func (o *Option[V]) AsValue(value V) *Option[V] {
	*o = NewOption(value)
	return o
}

func (o *Option[V]) AsEmpty() *Option[V] {
	*o = NilOption[V]()
	return o
}

func (o Option[V]) AsSlice() []V {
	if o.IsValue() {
		return []V{o.value}
	}

	return []V{}
}

func (o Option[V]) Expect(message string) V {
	if o.IsValue() {
		return o.value
	}

	panic(message)
}

func (o Option[V]) Unwrap() V {
	return o.Expect("called `Option.Unwrap()` on an empty value")
}

func (o Option[V]) UnwrapOr(value V) V {
	if o.IsValue() {
		return o.value
	}

	return value
}

func (o Option[V]) UnwrapOrElse(f func() V) V {
	if o.IsValue() {
		return o.value
	}

	return f()
}

func (o Option[V]) UnwrapOrDefault() V {
	if o.IsValue() {
		return o.value
	}

	var zero V
	return zero
}

func (o Option[V]) UnwrapUnchecked() V {
	return o.value
}

func (o Option[V]) Inspect(fn func(*V)) Option[V] {
	if o.IsValue() {
		fn(&o.value)
	}

	return o
}

func (o Option[V]) OkOr(err error) Result[V] {
	if o.IsValue() {
		return NewResult(o.value)
	}

	return ErrResult[V](err)
}

func (o Option[V]) OkOrElse(err func() error) Result[V] {
	if o.IsValue() {
		return NewResult(o.value)
	}

	return ErrResult[V](err())
}

func (o Option[V]) And(other Option[V]) Option[V] {
	if o.IsValue() {
		return other
	}

	return o
}

func (o Option[V]) AndThen(f func(V) Option[V]) Option[V] {
	if o.IsValue() {
		return f(o.value)
	}

	return o
}

func (o Option[V]) Filter(predicate func(*V) bool) Option[V] {
	if o.IsValue() && predicate(&o.value) {
		return o
	}

	return NilOption[V]()
}

func (o Option[V]) Or(other Option[V]) Option[V] {
	if o.IsValue() {
		return o
	}

	return other
}

func (o Option[V]) OrElse(f func() Option[V]) Option[V] {
	if o.IsValue() {
		return o
	}

	return f()
}

func (o Option[V]) Xor(other Option[V]) Option[V] {
	if o.IsValue() != other.IsValue() {
		if o.IsValue() {
			return o
		}
		return other
	}

	return NilOption[V]()
}

func (o *Option[V]) Insert(value V) *V {
	o.value, o.empty = value, false
	return &o.value
}

func (o *Option[V]) GetOrInsert(value V) *V {
	if o.IsEmpty() {
		o.value = value
	}
	return &o.value
}

func (o *Option[V]) GetOrInsertDefault() *V {
	if o.IsEmpty() {
		var zero V
		o.value = zero
	}
	return &o.value
}

func (o *Option[V]) GetOrInsertWith(f func() V) *V {
	if o.IsEmpty() {
		o.value = f()
	}
	return &o.value
}

func (o *Option[V]) Take() Option[V] {
	if o.IsValue() {
		o.AsEmpty()
		return NewOption(o.value)
	}
	return NilOption[V]()
}

func (o *Option[V]) TakeIf(predicate func(*V) bool) Option[V] {
	if o.IsValue() && predicate(&o.value) {
		o.AsEmpty()
		return NewOption(o.value)
	}

	return NilOption[V]()
}

func (o *Option[V]) Replace(value V) Option[V] {
	opt := *o
	o.Insert(value)
	return opt
}


