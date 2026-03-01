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
