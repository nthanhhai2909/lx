package lxconstraints

// This package defines reusable type constraints for the lx repo.

// Integer represents all signed and unsigned integer kinds.
// Use the ~ operator so named types with those underlying types are accepted.
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Float represents floating point kinds.
type Float interface {
	~float32 | ~float64
}

// Number includes integers and floats.
type Number interface {
	Integer | Float
}

// Ordered includes all types that support < and > comparisons.
// Include string here for convenience; adjust if you only want numeric ordering.
type Ordered interface {
	Number | ~string
}
