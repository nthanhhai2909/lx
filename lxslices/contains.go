package lxslices

// Contains returns true if the specified element is present in the slice, false otherwise.
func Contains[T comparable](slice []T, elem T) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}

// ContainsAny returns true if the specified element is present in any of the slices, false otherwise.
func ContainsAny[T comparable](slice []T, elems ...T) bool {
	for _, e := range slice {
		for _, elem := range elems {
			if e == elem {
				return true
			}
		}
	}
	return false
}

func NotContainsAny[T comparable](slice []T, elems ...T) bool {
	return !ContainsAny(slice, elems...)
}

// ContainsAll returns true if all the specified elements are present in the slice, false otherwise.
func ContainsAll[T comparable](slice []T, elems ...T) bool {
	for _, elem := range elems {
		if !Contains(slice, elem) {
			return false
		}
	}
	return true
}

// ContainsFunc returns true if the specified predicate returns true for any element in the slice, false otherwise.
func ContainsFunc[T any](slice []T, predicate func(T) bool) bool {
	for _, e := range slice {
		if predicate(e) {
			return true
		}
	}
	return false
}
