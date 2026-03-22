package lxmaps

// ContainsAnyKey returns true if any of the keys are present in the map.
func ContainsAnyKey[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if _, ok := m[key]; ok {
			return true
		}
	}
	return false
}
