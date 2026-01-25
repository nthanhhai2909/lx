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

// Length returns the length of the string.
func Length(s string) int {
	return len(s)
}

// LowerCase converts the string to lowercase.
func LowerCase(s string) string {
	return strings.ToLower(s)
}

// UpperCase converts the string to uppercase.
func UpperCase(s string) string {
	return strings.ToUpper(s)
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

// Trim removes all leading and trailing characters specified in cutset from the string.
func Trim(s string, cutset string) string {
	return strings.Trim(s, cutset)
}

// TrimLeft removes all leading characters specified in cutset from the string.
func TrimLeft(s string, cutset string) string {
	return strings.TrimLeft(s, cutset)
}

// TrimRight removes all trailing characters specified in cutset from the string.
func TrimRight(s string, cutset string) string {
	return strings.TrimRight(s, cutset)
}

// Split splits the string by the specified separator and returns a slice of substrings.
func Split(s, sep string) []string {
	return strings.Split(s, sep)
}

// Join joins a slice of strings into a single string with the specified separator.
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// Repeat returns a new string consisting of count copies of the string s.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StartBy checks if the string starts with the specified prefix.
func StartBy(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// StartByIgnoreCase checks if the string starts with the specified prefix, ignoring case.
func StartByIgnoreCase(s, prefix string) bool {
	sLower := strings.ToLower(s)
	prefixLower := strings.ToLower(prefix)
	return strings.HasPrefix(sLower, prefixLower)
}

// StartByAny checks if the string starts with any of the specified prefixes.
func StartByAny(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func StartByAnyIgnoreCase(s string, prefixes ...string) bool {
	sLower := strings.ToLower(s)
	for _, prefix := range prefixes {
		prefixLower := strings.ToLower(prefix)
		if strings.HasPrefix(sLower, prefixLower) {
			return true
		}
	}
	return false
}

// EndBy checks if the string ends with the specified suffix.
func EndBy(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// EndByIgnoreCase checks if the string ends with the specified suffix, ignoring case.
func EndByIgnoreCase(s, suffix string) bool {
	sLower := strings.ToLower(s)
	suffixLower := strings.ToLower(suffix)
	return strings.HasSuffix(sLower, suffixLower)
}

// EndByAny checks if the string ends with any of the specified suffixes.
func EndByAny(s string, suffixes ...string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// EndByAnyIgnoreCase checks if the string ends with any of the specified suffixes, ignoring case.
func EndByAnyIgnoreCase(s string, suffixes ...string) bool {
	sLower := strings.ToLower(s)
	for _, suffix := range suffixes {
		suffixLower := strings.ToLower(suffix)
		if strings.HasSuffix(sLower, suffixLower) {
			return true
		}
	}
	return false
}

// Replace replaces occurrences of old with new in the string s, up to n times.
// If n is -1, all occurrences are replaced.
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// ReplaceAll replaces all occurrences of old with new in the string s.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Remove removes all occurrences of substr from the string s.
func Remove(s, substr string) string {
	return strings.ReplaceAll(s, substr, "")
}

func RemoveIgnoreCase(s, substr string) string {
	if substr == "" {
		return s
	}

	lowerS := strings.ToLower(s)
	lowerSub := strings.ToLower(substr)

	var b strings.Builder
	i := 0

	for {
		j := Index(lowerS[i:], lowerSub)
		if j < 0 {
			b.WriteString(s[i:])
			break
		}

		j += i
		b.WriteString(s[i:j])
		i = j + len(substr)
	}

	return b.String()
}

// RemoveAny removes all occurrences of the specified substrings from the string s.
func RemoveAny(s string, substrs ...string) string {
	result := s
	for _, substr := range substrs {
		result = Remove(result, substr)
	}
	return result
}

// RemoveAnyIgnoreCase removes all occurrences of the specified substrings from the string s, ignoring case.
func RemoveAnyIgnoreCase(s string, substrs ...string) string {
	result := s
	for _, substr := range substrs {
		result = RemoveIgnoreCase(result, substr)
	}
	return result
}
