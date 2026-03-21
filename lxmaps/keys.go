package lxmaps

// Keys returns a slice of keys from a map.
func Keys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}
	size := 0
	for i := range in {
		size += len(in[i])
	}

	keys := make([]K, 0, size)
	for _, m := range in {
		for k := range m {
			keys = append(keys, k)
		}
	}
	return keys
}
