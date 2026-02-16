package lxslices

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
