package lxslices_test

import (
	"strings"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestReduce_IntSum(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(int, int) int
		initial  int
		expected int
	}{
		{
			name:     "sum of integers",
			slice:    []int{1, 2, 3, 4, 5},
			fn:       func(acc, v int) int { return acc + v },
			initial:  0,
			expected: 15,
		},
		{
			name:     "sum with non-zero initial",
			slice:    []int{1, 2, 3},
			fn:       func(acc, v int) int { return acc + v },
			initial:  10,
			expected: 16,
		},
		{
			name:     "product of integers",
			slice:    []int{1, 2, 3, 4},
			fn:       func(acc, v int) int { return acc * v },
			initial:  1,
			expected: 24,
		},
		{
			name:  "find maximum",
			slice: []int{3, 7, 2, 9, 1},
			fn: func(acc, v int) int {
				if v > acc {
					return v
				}
				return acc
			},
			initial:  0,
			expected: 9,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(acc, v int) int { return acc + v },
			initial:  42,
			expected: 42,
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(acc, v int) int { return acc + v },
			initial:  42,
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Reduce(tt.slice, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce_StringConcat(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		fn       func(string, string) string
		initial  string
		expected string
	}{
		{
			name:     "concatenate strings",
			slice:    []string{"Hello", " ", "World", "!"},
			fn:       func(acc, v string) string { return acc + v },
			initial:  "",
			expected: "Hello World!",
		},
		{
			name:  "join with separator",
			slice: []string{"apple", "banana", "cherry"},
			fn: func(acc, v string) string {
				if acc == "" {
					return v
				}
				return acc + ", " + v
			},
			initial:  "",
			expected: "apple, banana, cherry",
		},
		{
			name:     "concatenate with prefix",
			slice:    []string{"a", "b", "c"},
			fn:       func(acc, v string) string { return acc + v },
			initial:  "prefix:",
			expected: "prefix:abc",
		},
		{
			name:     "empty slice",
			slice:    []string{},
			fn:       func(acc, v string) string { return acc + v },
			initial:  "default",
			expected: "default",
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(acc, v string) string { return acc + v },
			initial:  "default",
			expected: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Reduce(tt.slice, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce_IntToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(string, int) string
		initial  string
		expected string
	}{
		{
			name:  "convert to comma-separated string",
			slice: []int{1, 2, 3, 4},
			fn: func(acc string, v int) string {
				if acc == "" {
					return strings.Repeat("*", v)
				}
				return acc + "," + strings.Repeat("*", v)
			},
			initial:  "",
			expected: "*,**,***,****",
		},
		{
			name:     "build string with numbers",
			slice:    []int{1, 2, 3},
			fn:       func(acc string, v int) string { return acc + strings.Repeat("#", v) },
			initial:  "start:",
			expected: "start:######",
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(acc string, v int) string { return acc + strings.Repeat("*", v) },
			initial:  "empty",
			expected: "empty",
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(acc string, v int) string { return acc + strings.Repeat("*", v) },
			initial:  "empty",
			expected: "empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Reduce(tt.slice, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestReduce_StructAggregation(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Age    int
		Active bool
	}

	type Summary struct {
		TotalUsers  int
		TotalAge    int
		ActiveCount int
	}

	tests := []struct {
		name     string
		slice    []User
		fn       func(Summary, User) Summary
		initial  Summary
		expected Summary
	}{
		{
			name: "aggregate user statistics",
			slice: []User{
				{1, "Alice", 25, true},
				{2, "Bob", 30, false},
				{3, "Charlie", 35, true},
			},
			fn: func(acc Summary, u User) Summary {
				acc.TotalUsers++
				acc.TotalAge += u.Age
				if u.Active {
					acc.ActiveCount++
				}
				return acc
			},
			initial:  Summary{},
			expected: Summary{TotalUsers: 3, TotalAge: 90, ActiveCount: 2},
		},
		{
			name: "count active users only",
			slice: []User{
				{1, "Alice", 25, true},
				{2, "Bob", 30, false},
				{3, "Charlie", 35, true},
			},
			fn: func(acc Summary, u User) Summary {
				if u.Active {
					acc.ActiveCount++
				}
				return acc
			},
			initial:  Summary{},
			expected: Summary{TotalUsers: 0, TotalAge: 0, ActiveCount: 2},
		},
		{
			name:  "empty slice",
			slice: []User{},
			fn: func(acc Summary, u User) Summary {
				acc.TotalUsers++
				return acc
			},
			initial:  Summary{TotalUsers: 5},
			expected: Summary{TotalUsers: 5},
		},
		{
			name:  "nil slice",
			slice: nil,
			fn: func(acc Summary, u User) Summary {
				acc.TotalUsers++
				return acc
			},
			initial:  Summary{TotalUsers: 5},
			expected: Summary{TotalUsers: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Reduce(tt.slice, tt.fn, tt.initial)
			if result != tt.expected {
				t.Errorf("Reduce() = %v; want %v", result, tt.expected)
			}
		})
	}
}
