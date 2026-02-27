package lxtuples

// Pair is a generic tuple type.
// This package is intentionally minimal; it can be extended with Triple, Quad, etc.
type Pair[T any, U any] struct {
	First  T
	Second U
}
