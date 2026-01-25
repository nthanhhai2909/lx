package lxstrings_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxstrings"
)

func TestAbbreviate(t *testing.T) {
	tests := []struct {
		input    string
		maxWidth int
		expected string
	}{
		{"Hello, World!", 5, "He..."},
		{"Hello", 10, "Hello"},
		{"GoLang", 3, "GoL"},
		{"Short", 0, ""},
		{"Exact", 5, "Exact"},
		{"This is a longer string", 8, "This ..."},
		{"😊😊😊😊😊", 4, "😊..."},
		{"こんにちは世界", 6, "こんに..."},
		{"", 5, ""},
		{"Test", 2, "Te"},
		{"Test", 3, "Tes"},
		{"Test", 4, "Test"},
	}
	for _, test := range tests {
		result := lxstrings.Abbreviate(test.input, test.maxWidth)
		if result != test.expected {
			t.Errorf("Abbreviate(%q, %d) = %q; want %q", test.input, test.maxWidth, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"hELLO", "HELLO"},
		{"", ""},
		{" ", " "},
		{"\n", "\n"},
		{"\r", "\r"},
		{"\t", "\t"},
		{"a \n", "A \n"},
		{"\rtest", "\rtest"},
		{"/", "/"},
		{"\r\n", "\r\n"},
		{"aBC", "ABC"},
		{"111", "111"},
		{"a", "A"},
		{"1test", "1test"},
		{"😊emoji", "😊emoji"},
		{"こんにちは", "こんにちは"},
	}
	for _, test := range tests {
		result := lxstrings.Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"apple", "banana", -1},
		{"banana", "apple", 1},
		{"cherry", "cherry", 0},
		{"", "", 0},
		{"a", "A", 1},
		{"A", "a", -1},
		{"abc", "abcd", -1},
		{"abcd", "abc", 1},
		{"こんにちは", "こんばんは", -1},
		{"😊", "😊", 0},
		{"😊", "😢", -1},
		{"😢", "😊", 1},
	}
	for _, test := range tests {
		result := lxstrings.Compare(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("Compare(%q, %q) = %d; want %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestCompareIgnoreCase(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected int
	}{
		{"apple", "BANANA", -1},
		{"BANANA", "apple", 1},
		{"cherry", "CHERRY", 0},
		{"", "", 0},
		{"a", "A", 0},
		{"A", "a", 0},
		{"abc", "ABCD", -1},
		{"ABCD", "abc", 1},
		{"GoLang", "golang", 0},
		{"HELLO", "hello", 0},
	}
	for _, test := range tests {
		result := lxstrings.CompareIgnoreCase(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("CompareIgnoreCase(%q, %q) = %d; want %d", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "world", true},
		{"hello world", "WORLD", false},
		{"golang", "go", true},
		{"golang", "lang", true},
		{"test", "TEST", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "世界", true},
		{"😊emoji😊", "emoji", true},
	}
	for _, test := range tests {
		result := lxstrings.Contains(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Contains(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  bool
	}{
		{"hello world", "WORLD", true},
		{"GoLang", "golang", true},
		{"TestString", "teststring", true},
		{"CaseSensitive", "casesensitive", true},
		{"NoMatch", "MATCH", true},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "世界", true},
		{"😊Emoji😊", "emoji", true},
	}
	for _, test := range tests {
		result := lxstrings.ContainsIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %v; want %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		s        string
		chars    []rune
		expected bool
	}{
		{"hello", []rune{'a', 'e', 'i'}, true},
		{"world", []rune{'x', 'y', 'z'}, false},
		{"golang", []rune{'g', 'o'}, true},
		{"test", []rune{'1', '2', '3'}, false},
		{"", []rune{'a', 'b'}, false},
		{"non-empty", []rune{}, false},
		{"こんにちは", []rune{'に', 'は'}, true},
		{"😊emoji😊", []rune{'😊'}, true},
	}
	for _, test := range tests {
		result := lxstrings.ContainsAny(test.s, test.chars...)
		if result != test.expected {
			t.Errorf("ContainsAny(%q, %v) = %v; want %v", test.s, test.chars, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"not empty", false},
		{" ", false},
		{"\n", false},
		{"\t", false},
		{"hello", false},
		{"こんにちは", false},
		{"😊", false},
		{"\r", false},
		{"\u200B", false},
	}
	for _, test := range tests {
		result := lxstrings.IsEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsEmpty(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"not empty", true},
		{" ", true},
		{"\n", true},
		{"\t", true},
		{"hello", true},
		{"こんにちは", true},
		{"😊", true},
		{"\r", true},
		{"\u200B", true},
	}
	for _, test := range tests {
		result := lxstrings.IsNotEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsNotEmpty(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"\n\t\r", true},
		{" not blank ", false},
		{"hello", false},
		{"こんにちは", false},
		{"😊", false},
		{" \n hello \t ", false},
	}
	for _, test := range tests {
		result := lxstrings.IsBlank(test.input)
		if result != test.expected {
			t.Errorf("IsBlank(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNotBlank(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"   ", false},
		{"\n\t\r", false},
		{" not blank ", true},
		{"hello", true},
		{"こんにちは", true},
		{"😊", true},
		{" \n hello \t ", true},
	}
	for _, test := range tests {
		result := lxstrings.IsNotBlank(test.input)
		if result != test.expected {
			t.Errorf("IsNotBlank(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "hello", true},
		{"hello", "Hello", false},
		{"", "", true},
		{"not empty", "", false},
		{"こんにちは", "こんにちは", true},
		{"😊", "😊", true},
		{"test", "test ", false},
	}
	for _, test := range tests {
		result := lxstrings.Equals(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("Equals(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestNotEquals(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "hello", false},
		{"hello", "Hello", true},
		{"", "", false},
		{"not empty", "", true},
		{"こんにちは", "こんにちは", false},
		{"😊", "😊", false},
		{"test", "test ", true},
	}
	for _, test := range tests {
		result := lxstrings.NotEquals(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("NotEquals(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "HELLO", true},
		{"GoLang", "golang", true},
		{"Test", "test ", false},
		{"こんにちは", "こんにちは", true},
		{"😊", "😊", true},
		{"NotEqual", "Different", false},
	}
	for _, test := range tests {
		result := lxstrings.EqualsIgnoreCase(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("EqualsIgnoreCase(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestNotEqualsIgnoreCase(t *testing.T) {
	tests := []struct {
		s1, s2   string
		expected bool
	}{
		{"hello", "HELLO", false},
		{"GoLang", "golang", false},
		{"Test", "test ", true},
		{"こんにちは", "こんにちは", false},
		{"😊", "😊", false},
		{"NotEqual", "Different", true},
	}
	for _, test := range tests {
		result := lxstrings.NotEqualsIgnoreCase(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("NotEqualsIgnoreCase(%q, %q) = %v; want %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  int
	}{
		{"hello world", "world", 6},
		{"hello world", "WORLD", -1},
		{"golang", "go", 0},
		{"golang", "lang", 2},
		{"test", "TEST", -1},
		{"", "", 0},
		{"non-empty", "", 0},
		{"", "non-empty", -1},
		{"こんにちは世界", "世界", 15},
		{"😊emoji😊", "emoji", 4},
	}
	for _, test := range tests {
		result := lxstrings.Index(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Index(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
		}
	}
}

func TestIndexIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			s, substr string
			expected  int
		}{
			{"hello world", "WORLD", 6},
			{"GoLang", "golang", 0},
			{"TestString", "teststring", 0},
			{"CaseSensitive", "casesensitive", 0},
			{"NoMatch", "MATCH", 2},
			{"", "", 0},
			{"non-empty", "", 0},
			{"", "non-empty", -1},
			{"こんにちは世界", "世界", 15},
			{"😊Emoji😊", "emoji", 4},
		}
		for _, test := range tests {
			result := lxstrings.IndexIgnoreCase(test.s, test.substr)
			if result != test.expected {
				t.Errorf("IndexIgnoreCase(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
			}
		}
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"   hello world   ", "hello world"},
		{"\n\tgolang\r\n", "golang"},
		{" no leading or trailing spaces ", "no leading or trailing spaces"},
		{"\u200Bzero width space\u200B", "\u200Bzero width space\u200B"},
		{"😊 emoji 😊", "😊 emoji 😊"},
		{"      ", ""},
		{"", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimSpace(test.input)
		if result != test.expected {
			t.Errorf("TrimSpace(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
