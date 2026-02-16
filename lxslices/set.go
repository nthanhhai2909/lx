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

// Difference TODO: implement
func Difference[T comparable](slice1, slice2 []T) []T {
	return nil
}

// Intersection TODO: implement
func Intersection[T comparable](slice1, slice2 []T) []T {
	return nil
}
