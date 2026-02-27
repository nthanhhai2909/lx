package lxslices

// Unique returns a new slice with the unique elements of the original slice.
// The order of the elements in the returned slice is the same as in the original slice.
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, e := range slice {
		if !seen[e] {
			seen[e] = true
			result = append(result, e)
		}
	}
	return result
}

// Difference returns a new slice containing elements that are in slice1 but not in slice2.
// The order of elements from slice1 is preserved. Duplicates in slice1 are preserved when
// they are not present in slice2 (i.e., this is not a set-deduplication operation).
func Difference[T comparable](slice1, slice2 []T) []T {
	if len(slice1) == 0 {
		return nil
	}

	m := make(map[T]struct{}, len(slice2))
	for _, e := range slice2 {
		m[e] = struct{}{}
	}
	var result []T
	for _, e := range slice1 {
		if _, found := m[e]; !found {
			result = append(result, e)
		}
	}
	return result
}

// Intersection TODO: implement
func Intersection[T comparable](slice1, slice2 []T) []T {
	return nil
}
