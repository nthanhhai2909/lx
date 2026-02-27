package lxslices

import "github.com/nthanhhai2909/lx/lxconstraints"

// Reduce applies the given function to each element of the slice and returns the result.
// The order of the elements in the returned slice is the same as in the original slice.
func Reduce[T, U any](slice []T, fn func(accumulator U, element T) U, initial U) U {
	result := initial
	for _, e := range slice {
		result = fn(result, e)
	}
	return result
}

// Sum returns the sum of all elements in the slice.
func Sum[T lxconstraints.Number](slice []T) T {
	var total T
	for _, v := range slice {
		total += v
	}
	return total
}

// Min returns (value, found) for safety on empty slices
func Min[T lxconstraints.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	res := slice[0]
	for _, v := range slice[1:] {
		if v < res {
			res = v
		}
	}
	return res, true
}

// Max returns (value, found) for safety on empty slices
func Max[T lxconstraints.Ordered](slice []T) (T, bool) {
	if len(slice) == 0 {
		var zero T
		return zero, false
	}
	res := slice[0]
	for _, v := range slice[1:] {
		if v > res {
			res = v
		}
	}
	return res, true
}
