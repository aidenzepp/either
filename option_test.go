package either

import "testing"

func TestOptionIsValue(t *testing.T) {
	o := NewOption[int](3)
	if p, e := o.IsValue(), true; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}

	o = NilOption[int]()
	if p, e := o.IsValue(), false; p != e{
		t.Errorf("provided %t, expected %t", p, e)
	}
}

func TestOptionIsValueAnd(t *testing.T) {
	predicate := func(value int) bool {
		return value == 3
	}
	o := NewOption[int](3)
	if p, e := o.IsValueAnd(predicate), true; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}

	o = NewOption[int](6)
	if p, e := o.IsValueAnd(predicate), false; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}
	
	o = NilOption[int]()
	if p, e := o.IsValueAnd(predicate), false; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}
}

func TestOptionIsEmpty(t *testing.T) {
	o := NewOption[int](3)
	if p, e := o.IsEmpty(), false; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}

	o = NilOption[int]()
	if p, e := o.IsEmpty(), true; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}
}

func TestOptionAsValue(t *testing.T) {
	predicate := func(value int) bool {
		return value == 6
	}
	o := NewOption[int](3)
	if p, e := o.AsValue(6).IsValueAnd(predicate), true; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}	
}

func TestOptionAsEmpty(t *testing.T) {
	o := NewOption[int](3)
	if p, e := o.AsEmpty().IsEmpty(), true; p != e {
		t.Errorf("provided %t, expected %t", p, e)
	}	
}

func TestOptionAsSlice(t *testing.T) {
	o := NewOption[int](3)
	p, e := o.AsSlice(), []int{3}
	if p == nil {
		t.Errorf("provided `nil`, expected a non-nil slice")
	} else if pl, el := len(p), len(e); pl != el {
		t.Errorf("provided length %d, expected length %d", pl, el)
	} else {
		for i, v := range p {
			if v != e[i] {
				t.Errorf("provided %v, expected %v", p, e)
			}
		}
	}

	o = NilOption[int]()
	p, e = o.AsSlice(), []int{}
	if p == nil {
		t.Errorf("provided `nil`, expected a non-nil slice")
	} else if pl, el := len(p), len(e); pl != el {
		t.Errorf("provided length %d, expected length %d", pl, el)
	}
}

