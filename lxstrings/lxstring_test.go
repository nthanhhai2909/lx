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