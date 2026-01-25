package lxstrings

import (
	"strings"
	"unicode"
)

const (
	// Empty represents an empty string.
	Empty = ""

	// Space represents a string with a single space.
	Space = " "

	// LF represents a line feed character.
	LF = "\n"

	// CR represents a carriage return character.
	CR = "\r"

	// CRLF represents a carriage return followed by a line feed.
	CRLF = "\r\n"
)

// Abbreviate shortens the string to the specified maxWidth, adding "..." if truncated.
// If the string is shorter than or equal to maxWidth, it is returned unchanged.
// If maxWidth is less than or equal to 3, it returns the first maxWidth characters.
func Abbreviate(s string, maxWidth int) string {
	runes := []rune(s)
	if len(runes) <= maxWidth {
		return s
	}
	if maxWidth <= 3 {
		return string(runes[:maxWidth])
	}
	return string(runes[:maxWidth-3]) + "..."
}

// Capitalize capitalizes the first character of the string.
// If the string is empty or starts with a non-letter, it is returned unchanged.
func Capitalize(s string) string {
	if IsBlank(s) {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// Compare compares two strings lexicographically.
// It returns an integer comparing two strings lexicographically.
// The result will be 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func Compare(s1, s2 string) int {
	return strings.Compare(s1, s2)
}

// CompareIgnoreCase compares two strings lexicographically, ignoring case.
// It returns an integer comparing two strings lexicographically, ignoring case.
// The result will be 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.
func CompareIgnoreCase(s1, s2 string) int {
	s1Lower := strings.ToLower(s1)
	s2Lower := strings.ToLower(s2)
	return strings.Compare(s1Lower, s2Lower)
}

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
