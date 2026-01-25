package lxstrings_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxstrings"
)

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