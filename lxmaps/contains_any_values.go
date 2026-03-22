package lxmaps

// ContainsAnyValues returns true if any of the values are present in the map.
func ContainsAnyValues[K comparable, V comparable](m map[K]V, values ...V) bool {
	for _, value := range values {
		for _, v := range m {
			if v == value {
				return true
			}
		}
	}
	return false
}
