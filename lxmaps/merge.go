package lxmaps

// Merge merges multiple maps into a single map.
// If there are duplicate keys, the value from the last map will be used.
func Merge[K comparable, V any](in ...map[K]V) map[K]V {
	out := make(map[K]V)
	for _, m := range in {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}
