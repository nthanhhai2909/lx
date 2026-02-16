package lxslices_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestUnique_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "remove duplicates",
			slice:    []int{1, 2, 3, 2, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "no duplicates",
			slice:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:  "all the same",
			slice: []int{1, 1, 1, 1, 1},
			expected: []int{
				1,
			},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Unique(tt.slice)
			if len(result) != len(tt.expected) {
				t.Errorf("Unique() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Unique() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestUnique_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "remove duplicates",
			slice:    []string{"a", "b", "c", "b", "a"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "no duplicates",
			slice:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:  "all the same",
			slice: []string{"a", "a", "a", "a", "a"},
			expected: []string{
				"a",
			},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Unique(tt.slice)
			if len(result) != len(tt.expected) {
				t.Errorf("Unique() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Unique() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestUnique_Struct(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Active bool
	}

	tests := []struct {
		name     string
		slice    []User
		expected []User
	}{
		{
			name: "remove duplicates",
			slice: []User{
				{1, "a", true},
				{2, "b", false},
				{3, "c", true},
				{2, "b", false},
				{1, "a", true},
			},
			expected: []User{
				{1, "a", true},
				{2, "b", false},
				{3, "c", true},
			},
		},
		{
			name: "no duplicates",
			slice: []User{
				{1, "a", true},
				{2, "b", false},
				{3, "c", true},
			},
			expected: []User{
				{1, "a", true},
				{2, "b", false},
				{3, "c", true},
			},
		},
		{
			name: "all the same",
			slice: []User{
				{1, "a", true},
				{1, "a", true},
				{1, "a", true},
				{1, "a", true},
				{1, "a", true},
			},
			expected: []User{
				{1, "a", true},
			},
		},
		{
			name:     "empty slice",
			slice:    []User{},
			expected: []User{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Unique(tt.slice)
			if len(result) != len(tt.expected) {
				t.Errorf("Unique() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Unique() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}
