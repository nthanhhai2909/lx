package lxslices

import "fmt"

// Map applies the given function to each element of the slice and returns a new slice with the results.
// The order of the elements in the returned slice is the same as in the original slice.
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, e := range slice {
		result[i] = fn(e)
	}
	return result
}

// FlatMap applies the given function to each element of the slice and returns a new slice with the results.
// The order of the elements in the returned slice is the same as in the original slice.
func FlatMap[T, U any](slice []T, fn func(T) []U) []U {
	var result []U
	for _, e := range slice {
		result = append(result, fn(e)...)
	}
	return result
}

// Reverse returns a new slice with the elements in reverse order.
func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, e := range slice {
		result[len(slice)-1-i] = e
	}
	return result
}

// GroupBy groups the elements of the slice by the given function.
// The order of the elements in the returned slice is the same as in the original slice.
func GroupBy[T any, K comparable](slice []T, fn func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, e := range slice {
		key := fn(e)
		result[key] = append(result[key], e)
	}
	return result
}

// UniqueGroupBy groups the elements of the slice by the given function and returns a map with only unique keys.
// It returns an error if the function returns duplicate keys.
func UniqueGroupBy[T any, K comparable](slice []T, fn func(T) K) (map[K]T, error) {
	result := make(map[K]T)
	for _, e := range slice {
		key := fn(e)
		if _, exists := result[key]; exists {
			return nil, fmt.Errorf("%w: %v", ErrDuplicateKey, key)
		}
		result[key] = e
	}
	return result, nil
}
