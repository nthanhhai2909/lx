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

// Contains checks if the substring is present in the string.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ContainsIgnoreCase checks if the substring is present in the string, ignoring case.
func ContainsIgnoreCase(s, substr string) bool {
	sLower := strings.ToLower(s)
	substrLower := strings.ToLower(substr)
	return strings.Contains(sLower, substrLower)
}

// ContainsAny checks if any of the specified characters are present in the string.
func ContainsAny(s string, chars ...rune) bool {
	for _, c := range s {
		for _, char := range chars {
			if c == char {
				return true
			}
		}
	}
	return false
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

// IsAlpha checks if the given string contains only alphabetic characters.
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsNumeric checks if the given string contains only numeric characters.
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsAlphaNumeric checks if the given string contains only alphanumeric characters.
func IsAlphaNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// Index returns the index of the first occurrence of substr in s, or -1 if not found.
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// LastIndex returns the index of the last occurrence of substr in s, or -1 if not found.
func LastIndex(s, substr string) int {
	return strings.LastIndex(s, substr)
}

// LastIndexIgnoreCase returns the index of the last occurrence of substr in s, ignoring case, or -1 if not found.
func LastIndexIgnoreCase(s, substr string) int {
	sLower := strings.ToLower(s)
	substrLower := strings.ToLower(substr)
	return strings.LastIndex(sLower, substrLower)
}

// IndexIgnoreCase returns the index of the first occurrence of substr in s, ignoring case, or -1 if not found.
func IndexIgnoreCase(s, substr string) int {
	sLower := strings.ToLower(s)
	substrLower := strings.ToLower(substr)
	return strings.Index(sLower, substrLower)
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

// TrimSpace removes leading and trailing whitespace from the string.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}