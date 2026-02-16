package lxslices_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestIndex_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elem     int
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []int{10, 20, 30},
			elem:     10,
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []int{10, 20, 30},
			elem:     20,
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []int{10, 20, 30},
			elem:     30,
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []int{10, 20, 30},
			elem:     40,
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elem:     1,
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     1,
			expected: -1,
		},
		{
			name:     "duplicate elements returns first index",
			slice:    []int{5, 7, 5, 9},
			elem:     5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndex_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elem     string
		expected int
	}{
		{
			name:     "element in middle",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: 1,
		},
		{
			name:     "element at beginning",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "apple",
			expected: 0,
		},
		{
			name:     "element at end",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "orange",
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "duplicate elements returns first index",
			slice:    []string{"go", "java", "go", "rust"},
			elem:     "go",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %q) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndex_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected int
	}{
		{
			name: "element at beginning",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
		{
			name: "element in middle",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element at end",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element not found",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			elem:     User{3, "Charlie"},
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name: "duplicate returns first index",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{1, "Alice"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Index(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("Index(%v, %+v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []int{1, 3, 4, 7},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2, // 4
		},
		{
			name:      "match at beginning",
			slice:     []int{2, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []int{1, 3, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns first",
			slice:     []int{1, 4, 6, 8},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "banana" },
			expected:  1,
		},
		{
			name:      "match at beginning",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "apple" },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "cherry" },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return len(v) == 4 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns first",
			slice:     []string{"go", "java", "go", "rust"},
			predicate: func(v string) bool { return len(v) == 2 },
			expected:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestIndexFunc_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  int
	}{
		{
			name: "match by ID",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			predicate: func(u User) bool {
				return u.ID == 2
			},
			expected: 1,
		},
		{
			name: "no match",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			predicate: func(u User) bool {
				return u.Name == "David"
			},
			expected: -1,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name: "multiple matches returns first",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Alice"
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.IndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("IndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		elem     int
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []int{10, 20, 30},
			elem:     10,
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []int{10, 20, 30},
			elem:     20,
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []int{10, 20, 30},
			elem:     30,
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []int{10, 20, 30},
			elem:     40,
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			elem:     1,
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     1,
			expected: -1,
		},
		{
			name:     "duplicate elements returns last index",
			slice:    []int{5, 7, 5, 9},
			elem:     5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_String(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elem     string
		expected int
	}{
		{
			name:     "element at beginning",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "apple",
			expected: 0,
		},
		{
			name:     "element in middle",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "banana",
			expected: 1,
		},
		{
			name:     "element at end",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "cherry",
			expected: 2,
		},
		{
			name:     "element not found",
			slice:    []string{"apple", "banana", "cherry"},
			elem:     "orange",
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []string{},
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     "apple",
			expected: -1,
		},
		{
			name:     "duplicate elements returns last index",
			slice:    []string{"go", "java", "go", "rust"},
			elem:     "go",
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %q) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndex_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name     string
		slice    []User
		elem     User
		expected int
	}{
		{
			name: "element at beginning",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{1, "Alice"},
			expected: 0,
		},
		{
			name: "element in middle",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{2, "Bob"},
			expected: 1,
		},
		{
			name: "element at end",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{3, "Charlie"},
			expected: 2,
		},
		{
			name: "element not found",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Charlie"},
			},
			elem:     User{4, "David"},
			expected: -1,
		},
		{
			name:     "empty slice",
			slice:    []User{},
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name:     "nil slice",
			slice:    nil,
			elem:     User{1, "Alice"},
			expected: -1,
		},
		{
			name: "duplicate elements returns last index",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{1, "Alice"},
			},
			elem:     User{1, "Alice"},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndex(tt.slice, tt.elem)
			if result != tt.expected {
				t.Errorf("LastIndex(%v, %+v) = %d; want %d",
					tt.slice, tt.elem, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_Int(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []int{1, 3, 4, 7},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "match at beginning",
			slice:     []int{2, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []int{1, 3, 6},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns last",
			slice:     []int{1, 4, 6, 8},
			predicate: func(v int) bool { return v%2 == 0 },
			expected:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_String(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		predicate func(string) bool
		expected  int
	}{
		{
			name:      "match in middle",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "banana" },
			expected:  1,
		},
		{
			name:      "match at beginning",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "apple" },
			expected:  0,
		},
		{
			name:      "match at end",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return v == "cherry" },
			expected:  2,
		},
		{
			name:      "no match",
			slice:     []string{"apple", "banana", "cherry"},
			predicate: func(v string) bool { return len(v) == 4 },
			expected:  -1,
		},
		{
			name:      "empty slice",
			slice:     []string{},
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v string) bool { return true },
			expected:  -1,
		},
		{
			name:      "multiple matches returns last",
			slice:     []string{"go", "java", "go", "rust"},
			predicate: func(v string) bool { return v == "go" },
			expected:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}

func TestLastIndexFunc_Struct(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}
	tests := []struct {
		name      string
		slice     []User
		predicate func(User) bool
		expected  int
	}{
		{
			name: "match by ID",
			slice: []User{
				{1, "Alice"},
			},
			predicate: func(u User) bool {
				return u.ID == 1
			},
			expected: 0,
		},
		{
			name: "no match",
			slice: []User{
				{1, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Bob"
			},
			expected: -1,
		},
		{
			name:      "empty slice",
			slice:     []User{},
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(u User) bool { return true },
			expected:  -1,
		},
		{
			name: "multiple matches returns last",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
				{3, "Alice"},
			},
			predicate: func(u User) bool {
				return u.Name == "Alice"
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.LastIndexFunc(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("LastIndexFunc(%v) = %d; want %d",
					tt.slice, result, tt.expected)
			}
		})
	}
}
