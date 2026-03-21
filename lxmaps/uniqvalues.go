package lxmaps

// UniqValues returns a slice of unique values from a map.
func UniqValues[K comparable, V comparable](in ...map[K]V) []V {
	if len(in) == 0 {
		return nil
	}

	uniqValueMap := make(map[V]struct{})
	for _, m := range in {
		for _, v := range m {
			uniqValueMap[v] = struct{}{}
		}
	}

	uniqValues := make([]V, 0, len(uniqValueMap))
	for v := range uniqValueMap {
		uniqValues = append(uniqValues, v)
	}

	return uniqValues
}
