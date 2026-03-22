package lxmaps

// ContainsAllKeys returns true if all of the keys are present in the map.
func ContainsAllKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	if len(keys) == 0 {
		return true
	}
	for _, key := range keys {
		if _, ok := m[key]; !ok {
			return false
		}
	}
	return true
}
