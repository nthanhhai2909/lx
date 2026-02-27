package lxslices

import "github.com/nthanhhai2909/lx/lxconstraints"

// Index returns the index of the first instance of elem in slice, or -1 if elem is not present in slice.
func Index[T comparable](slice []T, elem T) int {
	for i, e := range slice {
		if e == elem {
			return i
		}
	}
	return -1
}

// IndexFunc returns the index of the first element in slice for which predicate returns true, or -1 if none do.
func IndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i, e := range slice {
		if predicate(e) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of elem in slice, or -1 if elem is not present in slice.
func LastIndex[T comparable](slice []T, elem T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == elem {
			return i
		}
	}
	return -1
}

// LastIndexFunc returns the index of the last element in slice for which predicate returns true, or -1 if none do.
func LastIndexFunc[T any](slice []T, predicate func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(slice[i]) {
			return i
		}
	}
	return -1
}

// MinIndex returns the index of the minimum element in the slice, or -1 if the slice is empty.
func MinIndex[T lxconstraints.Ordered](slice []T) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	idx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[idx] {
			idx = i
		}
	}
	return idx, true
}

// MaxIndex returns the index of the maximum element in the slice, or -1 if the slice is empty.
func MaxIndex[T lxconstraints.Ordered](slice []T) (int, bool) {
	if len(slice) == 0 {
		return -1, false
	}
	idx := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] > slice[idx] {
			idx = i
		}
	}
	return idx, true
}
