package lxmaps

// Filter returns a new map with all elements that satisfy the predicate.
func Filter[K comparable, V any](m map[K]V, predicate func(k K, v V) bool) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}
