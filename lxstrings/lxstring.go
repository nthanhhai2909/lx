package lxstrings

// IsEmpty checks if the given string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty checks if the given string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsBlank checks if the given string is blank (empty or only whitespace).
func IsBlank(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\n' && r != '\t' && r != '\r' {
			return false
		}
	}
	return true
}

// IsNotBlank checks if the given string is not blank.
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}