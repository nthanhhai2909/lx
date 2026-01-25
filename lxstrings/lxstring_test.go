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