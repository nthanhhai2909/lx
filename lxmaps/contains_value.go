package lxmaps

// ContainsValue returns true if the map contains the value.
func ContainsValue[K comparable, V comparable](m map[K]V, in V) bool {
	for _, v := range m {
		if v == in {
			return true
		}
	}
	return false
}
