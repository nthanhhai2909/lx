package lxtypes_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxtypes"
)

func TestOptionalOf(t *testing.T) {
	opt := lxtypes.Of(42)

	if !opt.IsPresent() {
		t.Error("Expected Of to return true for IsPresent()")
	}
	if opt.IsEmpty() {
		t.Error("Expected Of to return false for IsEmpty()")
	}
	if got := opt.Get(); got != 42 {
		t.Errorf("Get() = %v, want 42", got)
	}
}

func TestOptionalEmpty(t *testing.T) {
	opt := lxtypes.Empty[int]()

	if opt.IsPresent() {
		t.Error("Expected Empty to return false for IsPresent()")
	}
	if !opt.IsEmpty() {
		t.Error("Expected Empty to return true for IsEmpty()")
	}
}

func TestOptionalGetPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Get on Empty to panic")
		}
	}()
	opt := lxtypes.Empty[int]()
	opt.Get()
}

func TestOptionalOrElse(t *testing.T) {
	present := lxtypes.Of(42)
	empty := lxtypes.Empty[int]()

	if got := present.OrElse(0); got != 42 {
		t.Errorf("Of.OrElse(0) = %v, want 42", got)
	}
	if got := empty.OrElse(99); got != 99 {
		t.Errorf("Empty.OrElse(99) = %v, want 99", got)
	}
}

func TestOptionalOrElseGet(t *testing.T) {
	present := lxtypes.Of(42)
	empty := lxtypes.Empty[int]()

	if got := present.OrElseGet(func() int { return 0 }); got != 42 {
		t.Errorf("Of.OrElseGet(...) = %v, want 42", got)
	}
	if got := empty.OrElseGet(func() int { return 99 }); got != 99 {
		t.Errorf("Empty.OrElseGet(...) = %v, want 99", got)
	}
}

func TestOptionalalMap(t *testing.T) {
	double := func(n int) int { return n * 2 }

	present := lxtypes.Of(21)
	mapped := lxtypes.OptionalMap(present, double)

	if !mapped.IsPresent() {
		t.Error("Expected mapped Of to be present")
	}
	if got := mapped.Get(); got != 42 {
		t.Errorf("Mapped value = %v, want 42", got)
	}

	empty := lxtypes.Empty[int]()
	mappedEmpty := lxtypes.OptionalMap(empty, double)

	if !mappedEmpty.IsEmpty() {
		t.Error("Expected mapped Empty to be empty")
	}
}

func TestOptionalAndThen(t *testing.T) {
	safeDivide := func(n int) lxtypes.Optional[int] {
		if n == 0 {
			return lxtypes.Empty[int]()
		}
		return lxtypes.Of(100 / n)
	}

	present := lxtypes.Of(10)
	result := lxtypes.OptionalAndThen(present, safeDivide)

	if !result.IsPresent() {
		t.Error("Expected present result after AndThen")
	}
	if got := result.Get(); got != 10 {
		t.Errorf("Result = %v, want 10", got)
	}

	presentZero := lxtypes.Of(0)
	resultEmpty := lxtypes.OptionalAndThen(presentZero, safeDivide)

	if !resultEmpty.IsEmpty() {
		t.Error("Expected empty result after AndThen with zero")
	}

	empty := lxtypes.Empty[int]()
	resultOriginal := lxtypes.OptionalAndThen(empty, safeDivide)

	if !resultOriginal.IsEmpty() {
		t.Error("Expected Empty to propagate through AndThen")
	}
}

func TestOptionalOr(t *testing.T) {
	opt1 := lxtypes.Of(42)
	opt2 := lxtypes.Of(99)
	empty := lxtypes.Empty[int]()

	if got := opt1.Or(opt2).Get(); got != 42 {
		t.Errorf("Of.Or(Of) = %v, want 42", got)
	}
	if got := opt1.Or(empty).Get(); got != 42 {
		t.Errorf("Of.Or(Empty) = %v, want 42", got)
	}
	if got := empty.Or(opt2).Get(); got != 99 {
		t.Errorf("Empty.Or(Of) = %v, want 99", got)
	}
	if !empty.Or(lxtypes.Empty[int]()).IsEmpty() {
		t.Error("Empty.Or(Empty) should be Empty")
	}
}

func TestOptionalOrElseSupply(t *testing.T) {
	present := lxtypes.Of(42)
	empty := lxtypes.Empty[int]()

	fallback := func() lxtypes.Optional[int] {
		return lxtypes.Of(99)
	}

	if got := present.OrElseSupply(fallback).Get(); got != 42 {
		t.Errorf("Of.OrElseSupply(...) = %v, want 42", got)
	}
	if got := empty.OrElseSupply(fallback).Get(); got != 99 {
		t.Errorf("Empty.OrElseSupply(...) = %v, want 99", got)
	}
}

func TestOptionalChaining(t *testing.T) {
	// Complex chaining scenario with Map and AndThen
	opt1 := lxtypes.Of(10)
	opt2 := lxtypes.OptionalMap(opt1, func(n int) int { return n * 2 })

	// Use AndThen to conditionally continue
	opt3 := lxtypes.OptionalAndThen(opt2, func(n int) lxtypes.Optional[int] {
		if n > 15 {
			return lxtypes.Of(n + 5)
		}
		return lxtypes.Empty[int]()
	})
	result := opt3.OrElse(0)

	if result != 25 {
		t.Errorf("Chained result = %v, want 25", result)
	}

	// Chaining that results in Empty
	opt4 := lxtypes.Of(5)
	opt5 := lxtypes.OptionalMap(opt4, func(n int) int { return n * 2 })
	opt6 := lxtypes.OptionalAndThen(opt5, func(n int) lxtypes.Optional[int] {
		if n > 15 {
			return lxtypes.Of(n)
		}
		return lxtypes.Empty[int]()
	})
	filtered := opt6.OrElse(99)

	if filtered != 99 {
		t.Errorf("Filtered result = %v, want 99", filtered)
	}
}

func TestOptionalOfNullable(t *testing.T) {
	// Non-nil pointer
	value := 42
	opt := lxtypes.OfNullable(&value)

	if !opt.IsPresent() {
		t.Error("Expected OfNullable with non-nil to be present")
	}
	if got := opt.Get(); got != 42 {
		t.Errorf("OfNullable value = %v, want 42", got)
	}

	// Nil pointer
	var nilPtr *int
	optNil := lxtypes.OfNullable(nilPtr)

	if !optNil.IsEmpty() {
		t.Error("Expected OfNullable with nil to be empty")
	}
}
