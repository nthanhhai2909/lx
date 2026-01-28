package lxptrs_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxptrs"
)

func TestRef(t *testing.T) {

	t.Run("String", func(t *testing.T) {
		v := "ref check"
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		v := true
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int", func(t *testing.T) {
		v := 30
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int64", func(t *testing.T) {
		var v int64
		v = 98
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int32", func(t *testing.T) {
		var v int32
		v = 12
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int16", func(t *testing.T) {
		var v int16
		v = -55
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Int8", func(t *testing.T) {
		var v int8
		v = 12
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit", func(t *testing.T) {
		var v uint
		v = 12
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit64", func(t *testing.T) {
		var v uint64
		v = 1999
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit32", func(t *testing.T) {
		var v uint32
		v = 19921
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit16", func(t *testing.T) {
		var v uint16
		v = 1000
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})

	t.Run("Unit8", func(t *testing.T) {
		var v uint8
		v = 1
		p := lxptrs.Ref(v)
		if *p != v {
			t.Errorf("Ref(%v) = %v; want %v", v, *p, v)
		}
	})
}

func TestDeref(t *testing.T) {

	t.Run("String", func(t *testing.T) {
		v := "my string"
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		v := true
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int", func(t *testing.T) {
		v := 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int64", func(t *testing.T) {
		var v int64
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int32", func(t *testing.T) {
		var v int32
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int16", func(t *testing.T) {
		var v int16
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("Int8", func(t *testing.T) {
		var v int8
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt", func(t *testing.T) {
		var v uint
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt64", func(t *testing.T) {
		var v uint64
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt32", func(t *testing.T) {
		var v uint32
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt16", func(t *testing.T) {
		var v uint16
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})

	t.Run("UInt8", func(t *testing.T) {
		var v uint8
		v = 10
		p := &v
		result := lxptrs.Deref(p)
		if result != v {
			t.Errorf("Deref(%v) = %v; want %v", *p, result, v)
		}
	})
}
