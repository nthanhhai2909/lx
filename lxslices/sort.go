package lxslices

import (
	"cmp"
	"sort"
)

// SortBy sorts the slice using the provided less function.
func SortBy[T any](slice []T, less func(T, T) bool) []T {
	sort.Slice(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	return slice
}

// StableSortBy sorts the slice using the provided less function.
// The order of equal elements is the same as in the original slice.
func StableSortBy[T any](slice []T, less func(T, T) bool) []T {
	sort.SliceStable(slice, func(i, j int) bool {
		return less(slice[i], slice[j])
	})
	return slice
}

// SortAsc sorts the slice in ascending order.
func SortAsc[T cmp.Ordered](slice []T) []T {
	return SortBy(slice, func(a, b T) bool { return a < b })
}

// SortDesc sorts the slice in descending order.
func SortDesc[T cmp.Ordered](slice []T) []T {
	return SortBy(slice, func(a, b T) bool { return a > b })
}
