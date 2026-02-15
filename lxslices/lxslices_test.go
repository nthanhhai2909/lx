package lxslices_test

import (
	"errors"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestEmpty(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		s := lxslices.Empty[string]()

		if s == nil {
			t.Fatalf("Empty[string]() = nil; want non-nil empty slice")
		}
		if len(s) != 0 {
			t.Fatalf("len(Empty[string]()) = %d; want 0", len(s))
		}

		s = append(s, "x")
		if len(s) != 1 || s[0] != "x" {
			t.Fatalf("append result = %v; want [x]", s)
		}
	})

	t.Run("int", func(t *testing.T) {
		s := lxslices.Empty[int]()

		if s == nil {
			t.Fatalf("Empty[int]() = nil; want non-nil empty slice")
		}
		if len(s) != 0 {
			t.Fatalf("len(Empty[int]()) = %d; want 0", len(s))
		}

		s = append(s, 42)
		if len(s) != 1 || s[0] != 42 {
			t.Fatalf("append result = %v; want [42]", s)
		}
	})
}

func TestEmptyInt(t *testing.T) {
	s := lxslices.EmptyInt()

	if s == nil {
		t.Fatalf("Empty[int]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[int]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyInt8(t *testing.T) {
	s := lxslices.EmptyInt8()

	if s == nil {
		t.Fatalf("Empty[int8]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[int8]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyInt16(t *testing.T) {
	s := lxslices.EmptyInt16()

	if s == nil {
		t.Fatalf("Empty[int16]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[int16]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyInt32(t *testing.T) {
	s := lxslices.EmptyInt32()

	if s == nil {
		t.Fatalf("Empty[int32]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[int32]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyInt64(t *testing.T) {
	s := lxslices.EmptyInt64()

	if s == nil {
		t.Fatalf("Empty[int64]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[int64]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyUint(t *testing.T) {
	s := lxslices.EmptyUint()

	if s == nil {
		t.Fatalf("Empty[uint]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[uint]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyUint8(t *testing.T) {
	s := lxslices.EmptyUint8()

	if s == nil {
		t.Fatalf("Empty[uint8]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[uint8]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyUint16(t *testing.T) {
	s := lxslices.EmptyUint16()

	if s == nil {
		t.Fatalf("Empty[uint16]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[uint16]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyUint32(t *testing.T) {
	s := lxslices.EmptyUint32()

	if s == nil {
		t.Fatalf("Empty[uint32]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[uint32]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyUint64(t *testing.T) {
	s := lxslices.EmptyUint64()

	if s == nil {
		t.Fatalf("Empty[uint64]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[uint64]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyFloat32(t *testing.T) {
	s := lxslices.EmptyFloat32()

	if s == nil {
		t.Fatalf("Empty[float32]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[float32]()) = %d; want 0", len(s))
	}

	s = append(s, 42.0)
	if len(s) != 1 || s[0] != 42.0 {
		t.Fatalf("append result = %v; want [42.0]", s)
	}
}

func TestEmptyFloat64(t *testing.T) {
	s := lxslices.EmptyFloat64()

	if s == nil {
		t.Fatalf("Empty[float64]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[float64]()) = %d; want 0", len(s))
	}

	s = append(s, 42.0)
	if len(s) != 1 || s[0] != 42.0 {
		t.Fatalf("append result = %v; want [42.0]", s)
	}
}

func TestEmptyBool(t *testing.T) {
	s := lxslices.EmptyBool()

	if s == nil {
		t.Fatalf("Empty[bool]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[bool]()) = %d; want 0", len(s))
	}
	s = append(s, false)

	if len(s) != 1 || s[0] != false {
		t.Fatalf("append result = %v; want [false]", s)
	}
}

func TestEmptyByte(t *testing.T) {
	s := lxslices.EmptyByte()

	if s == nil {
		t.Fatalf("Empty[byte]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[byte]()) = %d; want 0", len(s))
	}

	s = append(s, 42)
	if len(s) != 1 || s[0] != 42 {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyRune(t *testing.T) {
	s := lxslices.EmptyRune()

	if s == nil {
		t.Fatalf("Empty[rune]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[rune]()) = %d; want 0", len(s))
	}

	s = append(s, 'a')
	if len(s) != 1 || s[0] != 'a' {
		t.Fatalf("append result = %v; want [42]", s)
	}
}

func TestEmptyComplex64(t *testing.T) {
	s := lxslices.EmptyComplex64()

	if s == nil {
		t.Fatalf("Empty[complex64]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[complex64]()) = %d; want 0", len(s))
	}

	s = append(s, complex(1, 2))
	if len(s) != 1 || s[0] != complex(1, 2) {
		t.Fatalf("append result = %v; want [complex(1, 2)]", s)
	}
}

func TestEmptyComplex128(t *testing.T) {
	s := lxslices.EmptyComplex128()
	if s == nil {
		t.Fatalf("Empty[complex128]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[complex128]()) = %d; want 0", len(s))
	}

	s = append(s, complex(1, 2))
	if len(s) != 1 || s[0] != complex(1, 2) {
		t.Fatalf("append result = %v; want [complex(1, 2)]", s)
	}
}

func TestEmptyError(t *testing.T) {
	s := lxslices.EmptyError()
	if s == nil {
		t.Fatalf("Empty[error]() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(Empty[error]()) = %d; want 0", len(s))
	}

	errSentinel := errors.New("error")
	s = append(s, errSentinel)

	if len(s) != 1 || !errors.Is(s[0], errSentinel) {
		t.Fatalf("append result = %v; want [%v]", s, errSentinel)
	}
}

func TestEmptyString(t *testing.T) {
	s := lxslices.EmptyString()

	if s == nil {
		t.Fatalf("EmptyString() = nil; want non-nil empty slice")
	}
	if len(s) != 0 {
		t.Fatalf("len(EmptyString()) = %d; want 0", len(s))
	}

	s = append(s, "x")
	if len(s) != 1 || s[0] != "x" {
		t.Fatalf("append result = %v; want [x]", s)
	}
}
