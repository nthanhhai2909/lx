package lxstrings

import "strings"

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

// Equals checks if two strings are equal.
func Equals(s1, s2 string) bool {
	return s1 == s2
}

// NotEquals checks if two strings are not equal.
func NotEquals(s1, s2 string) bool {
	return !Equals(s1, s2)
}

// EqualsIgnoreCase checks if two strings are equal, ignoring case.
func EqualsIgnoreCase(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

// NotEqualsIgnoreCase checks if two strings are not equal, ignoring case.
func NotEqualsIgnoreCase(s1, s2 string) bool {
	return !EqualsIgnoreCase(s1, s2)
}