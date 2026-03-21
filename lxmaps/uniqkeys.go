package lxmaps

// UniqKeys returns a slice of unique keys from a map.
func UniqKeys[K comparable, V any](in ...map[K]V) []K {
	if len(in) == 0 {
		return nil
	}

	seen := make(map[K]struct{})
	for _, m := range in {
		for k := range m {
			seen[k] = struct{}{}
		}
	}
	return Keys(seen)
}
