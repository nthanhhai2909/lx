package lxslices

// Reduce applies the given function to each element of the slice and returns the result.
// The order of the elements in the returned slice is the same as in the original slice.
func Reduce[T, U any](slice []T, fn func(accumulator U, element T) U, initial U) U {
	result := initial
	for _, e := range slice {
		result = fn(result, e)
	}
	return result
}
