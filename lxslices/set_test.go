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

func TestDifference_Int(t *testing.T) {
	tests := []struct {
		name     string
		s1       []int
		s2       []int
		expected []int
	}{
		{
			name:     "basic difference",
			s1:       []int{1, 2, 3, 4},
			s2:       []int{2, 4},
			expected: []int{1, 3},
		},
		{
			name:     "no overlap",
			s1:       []int{1, 2},
			s2:       []int{3, 4},
			expected: []int{1, 2},
		},
		{
			name:     "all removed",
			s1:       []int{1, 2},
			s2:       []int{1, 2},
			expected: nil,
		},
		{
			name:     "duplicates preserved",
			s1:       []int{1, 2, 2, 3},
			s2:       []int{4},
			expected: []int{1, 2, 2, 3},
		},
		{
			name:     "empty s1",
			s1:       []int{},
			s2:       []int{1},
			expected: nil,
		},
		{
			name:     "nil s1",
			s1:       nil,
			s2:       []int{1},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Difference(tt.s1, tt.s2)
			if len(result) != len(tt.expected) {
				t.Errorf("Difference() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Difference() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestDifference_String(t *testing.T) {
	tests := []struct {
		name     string
		s1       []string
		s2       []string
		expected []string
	}{
		{
			name:     "basic difference",
			s1:       []string{"a", "b", "c"},
			s2:       []string{"b"},
			expected: []string{"a", "c"},
		},
		{
			name:     "no overlap",
			s1:       []string{"a"},
			s2:       []string{"b"},
			expected: []string{"a"},
		},
		{
			name:     "all removed",
			s1:       []string{"a"},
			s2:       []string{"a"},
			expected: nil,
		},
		{
			name:     "empty s1",
			s1:       []string{},
			s2:       []string{"a"},
			expected: nil,
		},
		{
			name:     "nil s1",
			s1:       nil,
			s2:       []string{"a"},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Difference(tt.s1, tt.s2)
			if len(result) != len(tt.expected) {
				t.Errorf("Difference() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Difference() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestDifference_Struct(t *testing.T) {
	type Item struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		s1       []Item
		s2       []Item
		expected []Item
	}{
		{
			name:     "basic difference",
			s1:       []Item{{1, "a"}, {2, "b"}, {3, "c"}},
			s2:       []Item{{2, "b"}},
			expected: []Item{{1, "a"}, {3, "c"}},
		},
		{
			name:     "no overlap",
			s1:       []Item{{1, "a"}},
			s2:       []Item{{2, "b"}},
			expected: []Item{{1, "a"}},
		},
		{
			name:     "all removed",
			s1:       []Item{{1, "a"}},
			s2:       []Item{{1, "a"}},
			expected: nil,
		},
		{
			name:     "empty s1",
			s1:       []Item{},
			s2:       []Item{{1, "a"}},
			expected: nil,
		},
		{
			name:     "nil s1",
			s1:       nil,
			s2:       []Item{{1, "a"}},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Difference(tt.s1, tt.s2)
			if len(result) != len(tt.expected) {
				t.Errorf("Difference() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Difference() = %v; want %v", result, tt.expected)
				}
			}
		})
	}
}
