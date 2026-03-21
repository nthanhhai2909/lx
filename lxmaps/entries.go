package lxmaps

import "github.com/nthanhhai2909/lx/lxtypes"

// Entries returns a slice of key-value pairs from a map.
func Entries[K comparable, V any](in ...map[K]V) []lxtypes.Pair[K, V] {
	if len(in) == 0 {
		return nil
	}

	size := 0
	for i := range in {
		size += len(in[i])
	}

	entries := make([]lxtypes.Pair[K, V], 0, size)
	for _, m := range in {
		for k, v := range m {
			entries = append(entries, lxtypes.Pair[K, V]{First: k, Second: v})
		}
	}

	return entries
}
