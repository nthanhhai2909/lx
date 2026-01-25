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

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", true},
		{"HelloWorld", true},
		{"hello123", false},
		{"123", false},
		{"", false},
		{"こんにちは", true},
		{"😊emoji", false},
	}
	for _, test := range tests {
		result := lxstrings.IsAlpha(test.input)
		if result != test.expected {
			t.Errorf("IsAlpha(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"12345", true},
		{"00123", true},
		{"123abc", false},
		{"abc", false},
		{"", false},
		{"１２３４５", true}, // Full-width digits
		{"😊123", false},
	}
	for _, test := range tests {
		result := lxstrings.IsNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsNumeric(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{{"hello123", true},
		{"HelloWorld", true},
		{"12345", true},
		{"hello!", false},
		{"", false},
		{"こんにちは123", true},
		{"😊emoji123", false},
	}
	for _, test := range tests {
		result := lxstrings.IsAlphaNumeric(test.input)
		if result != test.expected {
			t.Errorf("IsAlphaNumeric(%q) = %v; want %v", test.input, result, test.expected)
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

func TestLastIndex(t *testing.T) {
	tests := []struct {
		s, substr string
		expected  int
	}{
		{"hello world world", "world", 12},
		{"hello world", "WORLD", -1},
		{"golang golang", "go", 7},
		{"test test test", "test", 10},
		{"", "", 0},
		{"non-empty", "", 9},
		{"", "non-empty", -1},
		{"こんにちは世界こんにちは", "こんにちは", 21},
		{"😊emoji😊emoji😊", "emoji", 13},
	}
	for _, test := range tests {
		result := lxstrings.LastIndex(test.s, test.substr)
		if result != test.expected {
			t.Errorf("LastIndex(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
		}
	}
}

func TestLastIndexIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			s, substr string
			expected  int
		}{
			{"hello world WORLD", "WORLD", 12},
			{"GoLang goLANG", "golang", 7},
			{"TestString teststring TESTSTRING", "teststring", 22},
			{"CaseSensitive CASESENSITIVE", "casesensitive", 14},
			{"NoMatch MATCH NOMATCH", "MATCH", 16},
			{"", "", 0},
			{"non-empty", "", 9},
			{"", "non-empty", -1},
			{"こんにちは世界こんにちは", "こんにちは", 21},
			{"😊Emoji😊emoji😊", "emoji", 13},
		}
		for _, test := range tests {
			result := lxstrings.LastIndexIgnoreCase(test.s, test.substr)
			if result != test.expected {
				t.Errorf("LastIndexIgnoreCase(%q, %q) = %d; want %d", test.s, test.substr, result, test.expected)
			}
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hello", 5},
		{"", 0},
		{"こんにちは", 15}, // Each Japanese character is 3 bytes in UTF-8
		{"😊emoji", 9}, // Emoji is 4 bytes, 'emoji' is 5 bytes
		{" ", 1},
		{"\n", 1},
		{"\t", 1},
	}
	for _, test := range tests {
		result := lxstrings.Length(test.input)
		if result != test.expected {
			t.Errorf("Length(%q) = %d; want %d", test.input, result, test.expected)
		}
	}
}

func TestLowerCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "hello"},
		{"Hello World", "hello world"},
		{"golang", "golang"},
		{"", ""},
		{"123ABC", "123abc"},
		{"こんにちは", "こんにちは"},
		{"😊EMOJI", "😊emoji"},
	}
	for _, test := range tests {
		result := lxstrings.LowerCase(test.input)
		if result != test.expected {
			t.Errorf("LowerCase(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestUpperCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "HELLO"},
		{"Hello World", "HELLO WORLD"},
		{"GOLANG", "GOLANG"},
		{"", ""},
		{"123abc", "123ABC"},
		{"こんにちは", "こんにちは"},
		{"😊emoji", "😊EMOJI"},
	}
	for _, test := range tests {
		result := lxstrings.UpperCase(test.input)
		if result != test.expected {
			t.Errorf("UpperCase(%q) = %q; want %q", test.input, result, test.expected)
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

func TestTrim(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "hello"},
		{"***golang***", "*", "golang"},
		{"   spaced   ", " ", "spaced"},
		{"!!!exciting!!!", "!", "exciting"},
		{"😊emoji😊", "😊", "emoji"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.Trim(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("Trim(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTrimLeft(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "hello---"},
		{"***golang***", "*", "golang***"},
		{"   spaced   ", " ", "spaced   "},
		{"!!!exciting!!!", "!", "exciting!!!"},
		{"😊emoji😊", "😊", "emoji😊"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimLeft(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("TrimLeft(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTrimRight(t *testing.T) {
	tests := []struct {
		input    string
		cutset   string
		expected string
	}{
		{"---hello---", "-", "---hello"},
		{"***golang***", "*", "***golang"},
		{"   spaced   ", " ", "   spaced"},
		{"!!!exciting!!!", "!", "!!!exciting"},
		{"😊emoji😊", "😊", "😊emoji"},
		{"no trim needed", "xyz", "no trim needed"},
		{"aaaaaa", "a", ""},
		{"", "abc", ""},
	}
	for _, test := range tests {
		result := lxstrings.TrimRight(test.input, test.cutset)
		if result != test.expected {
			t.Errorf("TrimRight(%q, %q) = %q; want %q", test.input, test.cutset, result, test.expected)
		}
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		input    string
		maxWidth int
		expected string
	}{
		{"Hello, World!", 5, "Hello"},
		{"Hello", 10, "Hello"},
		{"GoLang", 3, "GoL"},
		{"Short", 0, ""},
		{"Exact", 5, "Exact"},
		{"This is a longer string", 8, "This is "},
	}
	for _, test := range tests {
		result := lxstrings.Truncate(test.input, test.maxWidth)
		if result != test.expected {
			t.Errorf("Truncate(%q, %d) = %q; want %q", test.input, test.maxWidth, result, test.expected)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"a,b,c", ",", []string{"a", "b", "c"}},
		{"hello world", " ", []string{"hello", "world"}},
		{"one;two;three", ";", []string{"one", "two", "three"}},
		{"no separator", ",", []string{"no separator"}},
		{"", ",", []string{""}},
		{"a--b--c", "--", []string{"a", "b", "c"}},
		{"😊-emoji-😊", "-", []string{"😊", "emoji", "😊"}},
	}
	for _, test := range tests {
		result := lxstrings.Split(test.input, test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Split(%q, %q) = %v; want %v", test.input, test.sep, result, test.expected)
				break
			}
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		input    []string
		sep      string
		expected string
	}{
		{[]string{"a", "b", "c"}, ",", "a,b,c"},
		{[]string{"hello", "world"}, " ", "hello world"},
		{[]string{"one", "two", "three"}, ";", "one;two;three"},
		{[]string{"no", "separator"}, "", "noseparator"},
		{[]string{}, ",", ""},
		{[]string{"😊", "emoji", "😊"}, "-", "😊-emoji-😊"},
	}
	for _, test := range tests {
		result := lxstrings.Join(test.input, test.sep)
		if result != test.expected {
			t.Errorf("Join(%v, %q) = %q; want %q", test.input, test.sep, result, test.expected)
		}
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected string
	}{
		{"ha", 3, "hahaha"},
		{"go", 0, ""},
		{"!", 5, "!!!!!"},
		{"😊", 4, "😊😊😊😊"},
		{"abc", 1, "abc"},
		{"", 10, ""},
	}
	for _, test := range tests {
		result := lxstrings.Repeat(test.input, test.count)
		if result != test.expected {
			t.Errorf("Repeat(%q, %d) = %q; want %q", test.input, test.count, result, test.expected)
		}
	}
}

func TestStartBy(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"golang", "go", true},
		{"test", "TEST", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "こんにちは", true},
		{"😊emoji😊", "😊", true},
	}
	for _, test := range tests {
		result := lxstrings.StartBy(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("StartBy(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestStartByIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "HELLO", true},
		{"GoLang", "golang", true},
		{"TestString", "teststring", true},
		{"CaseSensitive", "casesensitive", true},
		{"NoMatch", "MATCH", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "こんにちは", true},
		{"😊Emoji😊", "emoji", false},
	}
	for _, test := range tests {
		result := lxstrings.StartByIgnoreCase(test.s, test.prefix)
		if result != test.expected {
			t.Errorf("StartByIgnoreCase(%q, %q) = %v; want %v", test.s, test.prefix, result, test.expected)
		}
	}
}

func TestStartByAny(t *testing.T) {
	tests := []struct {
		s        string
		prefixes []string
		expected bool
	}{
		{"hello world", []string{"hi", "hello"}, true},
		{"hello world", []string{"world", "planet"}, false},
		{"golang", []string{"go", "lang"}, true},
		{"test", []string{"TEST", "exam"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんにちは", "こんばんは"}, true},
		{"😊emoji😊", []string{"😊", "😢"}, true},
	}
	for _, test := range tests {
		result := lxstrings.StartByAny(test.s, test.prefixes...)
		if result != test.expected {
			t.Errorf("StartByAny(%q, %v) = %v; want %v", test.s, test.prefixes, result, test.expected)
		}
	}
}

func TestStartByAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		prefixes []string
		expected bool
	}{
		{"hello world", []string{"HI", "HELLO"}, true},
		{"GoLang", []string{"WORLD", "PLANET"}, false},
		{"TestString", []string{"teststring", "exam"}, true},
		{"CaseSensitive", []string{"CASESENSITIVE", "other"}, true},
		{"NoMatch", []string{"MATCH", "different"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんにちは", "こんばんは"}, true},
		{"😊Emoji😊", []string{"emoji", "😢"}, false},
	}
	for _, test := range tests {
		result := lxstrings.StartByAnyIgnoreCase(test.s, test.prefixes...)
		if result != test.expected {
			t.Errorf("StartByAnyIgnoreCase(%q, %v) = %v; want %v", test.s, test.prefixes, result, test.expected)
		}
	}
}

func TestEndBy(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
		{"golang", "lang", true},
		{"test", "TEST", false},
		{"", "", true},
		{"non-empty", "", true},
		{"", "non-empty", false},
		{"こんにちは世界", "世界", true},
		{"😊emoji😊", "😊", true},
	}
	for _, test := range tests {
		result := lxstrings.EndBy(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("EndBy(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestEndByIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		suffix   string
		expected bool
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
		{"😊Emoji😊", "emoji", false},
	}
	for _, test := range tests {
		result := lxstrings.EndByIgnoreCase(test.s, test.suffix)
		if result != test.expected {
			t.Errorf("EndByIgnoreCase(%q, %q) = %v; want %v", test.s, test.suffix, result, test.expected)
		}
	}
}

func TestEndByAny(t *testing.T) {
	tests := []struct {
		s        string
		suffixes []string
		expected bool
	}{
		{"hello world", []string{"planet", "world"}, true},
		{"hello world", []string{"hello", "hi"}, false},
		{"golang", []string{"lang", "go"}, true},
		{"test", []string{"TEST", "exam"}, false},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんばんは", "世界"}, true},
		{"😊emoji😊", []string{"😢", "😊"}, true},
	}
	for _, test := range tests {
		result := lxstrings.EndByAny(test.s, test.suffixes...)
		if result != test.expected {
			t.Errorf("EndByAny(%q, %v) = %v; want %v", test.s, test.suffixes, result, test.expected)
		}
	}
}

func TestEndByAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		suffixes []string
		expected bool
	}{
		{"hello world", []string{"PLANET", "WORLD"}, true},
		{"GoLang", []string{"HELLO", "HI"}, false},
		{"TestString", []string{"teststring", "exam"}, true},
		{"CaseSensitive", []string{"CASESENSITIVE", "other"}, true},
		{"NoMatch", []string{"MATCH", "different"}, true},
		{"", []string{""}, true},
		{"non-empty", []string{}, false},
		{"こんにちは世界", []string{"こんばんは", "世界"}, true},
		{"😊Emoji😊", []string{"emoji", "😢"}, false},
	}
	for _, test := range tests {
		result := lxstrings.EndByAnyIgnoreCase(test.s, test.suffixes...)
		if result != test.expected {
			t.Errorf("EndByAnyIgnoreCase(%q, %v) = %v; want %v", test.s, test.suffixes, result, test.expected)
		}
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		s        string
		old      string
		new      string
		n        int
		expected string
	}{
		{"hello world", "world", "gopher", -1, "hello gopher"},
		{"golang golang", "go", "GO", 1, "GOlang golang"},
		{"test test test", "test", "exam", 2, "exam exam test"},
		{"no match here", "xyz", "abc", -1, "no match here"},
		{"", "anything", "something", -1, ""},
		{"😊emoji😊emoji😊", "emoji", "EMOJI", 2, "😊EMOJI😊EMOJI😊"},
	}
	for _, test := range tests {
		result := lxstrings.Replace(test.s, test.old, test.new, test.n)
		if result != test.expected {
			t.Errorf("Replace(%q, %q, %q, %d) = %q; want %q", test.s, test.old, test.new, test.n, result, test.expected)
		}
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		s        string
		old      string
		new      string
		expected string
	}{
		{"hello world", "world", "gopher", "hello gopher"},
		{"golang golang", "go", "GO", "GOlang GOlang"},
		{"test test test", "test", "exam", "exam exam exam"},
		{"no match here", "xyz", "abc", "no match here"},
		{"", "anything", "something", ""},
		{"😊emoji😊emoji😊", "emoji", "EMOJI", "😊EMOJI😊EMOJI😊"},
	}
	for _, test := range tests {
		result := lxstrings.ReplaceAll(test.s, test.old, test.new)
		if result != test.expected {
			t.Errorf("ReplaceAll(%q, %q, %q) = %q; want %q", test.s, test.old, test.new, result, test.expected)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"golang golang", "go", "lang lang"},
		{"test test test", "test", "  "},
		{"no match here", "xyz", "no match here"},
		{"", "anything", ""},
		{"😊emoji😊emoji😊", "emoji", "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.Remove(test.s, test.substr)
		if result != test.expected {
			t.Errorf("Remove(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestRemoveIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello WORLD", "world", "hello "},
		{"GoLang golang", "GOLANG", " "},
		{"Test test TEST", "test", "  "},
		{"no match here", "XYZ", "no match here"},
		{"", "ANYTHING", ""},
		{"😊Emoji😊emoji😊", "EMOJI", "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.RemoveIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("RemoveIgnoreCase(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestRemoveAny(t *testing.T) {
	tests := []struct {
		s        string
		substrs  []string
		expected string
	}{
		{"hello world", []string{"hello", "world"}, " "},
		{"golang golang", []string{"go", "lang"}, " "},
		{"test test test", []string{"test", "exam"}, "  "},
		{"no match here", []string{"xyz", "abc"}, "no match here"},
		{"", []string{"anything", "something"}, ""},
		{"😊emoji😊emoji😊", []string{"emoji", "EMOJI"}, "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.RemoveAny(test.s, test.substrs...)
		if result != test.expected {
			t.Errorf("RemoveAny(%q, %v) = %q; want %q", test.s, test.substrs, result, test.expected)
		}
	}
}

func TestRemoveAnyIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substrs  []string
		expected string
	}{
		{"hello WORLD", []string{"HELLO", "world"}, " "},
		{"GoLang golang", []string{"GOLANG", "go"}, " "},
		{"Test test TEST", []string{"TEST", "test"}, "  "},
		{"no match here", []string{"XYZ", "ABC"}, "no match here"},
		{"", []string{"ANYTHING", "SOMETHING"}, ""},
		{"😊Emoji😊emoji😊", []string{"EMOJI", "emoji"}, "😊😊😊"},
	}
	for _, test := range tests {
		result := lxstrings.RemoveAnyIgnoreCase(test.s, test.substrs...)
		if result != test.expected {
			t.Errorf("RemoveAnyIgnoreCase(%q, %v) = %q; want %q", test.s, test.substrs, result, test.expected)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"golang", "gnalog"},
		{"😊emoji", "ijome😊"},
		{"こんにちは", "はちにんこ"},
	}
	for _, test := range tests {
		result := lxstrings.Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

func TestSubString(t *testing.T) {
	tests := []struct {
		s        string
		start    int
		end      int
		expected string
	}{
		{"hello world", 0, 5, "hello"},
		{"golang", 3, 6, "ang"},
		{"test string", 5, 11, "string"},
		{"😊emoji😊", 1, 6, "emoji"},
		{"short", 2, 10, "ort"},
	}
	for _, test := range tests {
		result := lxstrings.SubString(test.s, test.start, test.end)
		if result != test.expected {
			t.Errorf("SubString(%q, %d, %d) = %q; want %q", test.s, test.start, test.end, result, test.expected)
		}
	}
}

func TestSubStringBefore(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello world", "world", "hello "},
		{"golang programming", "programming", "golang "},
		{"test string example", "string", "test "},
		{"no match here", "xyz", ""},
	}
	for _, test := range tests {
		result := lxstrings.SubStringBefore(test.s, test.substr)
		if result != test.expected {
			t.Errorf("SubStringBefore(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}
func TestSubStringBeforeIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello WORLD", "world", "hello "},
		{"GoLang PROGRAMMING", "programming", "GoLang "},
		{"Test STRING example", "string", "Test "},
		{"no match here", "XYZ", ""},
	}
	for _, test := range tests {
		result := lxstrings.SubStringBeforeIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("SubStringBeforeIgnoreCase(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestSubStringAfter(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected string
	}{
		{"hello world", "hello", " world"},
		{"golang programming", "golang", " programming"},
		{"test string example", "string", " example"},
		{"no match here", "xyz", ""},
	}
	for _, test := range tests {
		result := lxstrings.SubStringAfter(test.s, test.substr)
		if result != test.expected {
			t.Errorf("SubStringAfter(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
		}
	}
}

func TestSubStringAfterIgnoreCase(t *testing.T) {
	{
		tests := []struct {
			s        string
			substr   string
			expected string
		}{
			{"hello WORLD", "HELLO", " WORLD"},
			{"GoLang PROGRAMMING", "GOLANG", " PROGRAMMING"},
			{"Test STRING example", "STRING", " example"},
			{"no match here", "XYZ", ""},
		}
		for _, test := range tests {
			result := lxstrings.SubStringAfterIgnoreCase(test.s, test.substr)
			if result != test.expected {
				t.Errorf("SubStringAfterIgnoreCase(%q, %q) = %q; want %q", test.s, test.substr, result, test.expected)
			}
		}
	}
}

func TestPadLeft(t *testing.T) {
	{
		tests := []struct {
			s        string
			length   int
			padChar  string
			expected string
		}{
			{"hello", 10, "*", "*****hello"},
			{"golang", 8, "0", "00golang"},
			{"test", 6, "-", "--test"},
			{"short", 3, "x", "short"},
		}
		for _, test := range tests {
			result := lxstrings.PadLeft(test.s, test.length, test.padChar)
			if result != test.expected {
				t.Errorf("PadLeft(%q, %d, %q) = %q; want %q", test.s, test.length, test.padChar, result, test.expected)
			}
		}
	}
}

func TestPadRight(t *testing.T) {
	tests := []struct {
		s        string
		length   int
		padChar  string
		expected string
	}{
		{"hello", 10, "*", "hello*****"},
		{"golang", 8, "0", "golang00"},
		{"test", 6, "-", "test--"},
		{"short", 3, "x", "short"},
	}
	for _, test := range tests {
		result := lxstrings.PadRight(test.s, test.length, test.padChar)
		if result != test.expected {
			t.Errorf("PadRight(%q, %d, %q) = %q; want %q", test.s, test.length, test.padChar, result, test.expected)
		}
	}
}

func TestPadCenter(t *testing.T) {
	tests := []struct {
		s        string
		length   int
		padChar  string
		expected string
	}{
		{"hello", 11, "*", "***hello***"},
		{"golang", 10, "0", "00golang00"},
		{"test", 8, "-", "--test--"},
		{"short", 3, "x", "short"},
	}
	for _, test := range tests {
		result := lxstrings.PadCenter(test.s, test.length, test.padChar)
		if result != test.expected {
			t.Errorf("PadCenter(%q, %d, %q) = %q; want %q", test.s, test.length, test.padChar, result, test.expected)
		}
	}
}
