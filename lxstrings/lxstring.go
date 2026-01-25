package lxstrings

// IsEmpty checks if the given string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}