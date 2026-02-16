package lxslices_test

import (
	"reflect"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestSortBy_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		less     func(int, int) bool
		expected []int
	}{
		{
			name:  "ascending",
			input: []int{5, 2, 4, 1, 3},
			less: func(a, b int) bool {
				return a < b
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "same values",
			input: []int{5, 5, 5, 5, 5},
			less: func(a, b int) bool {
				return a < b
			},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:  "sorted by ascending, no change",
			input: []int{1, 2, 3, 4, 5},
			less: func(a, b int) bool {
				return a < b
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "sorted by descending, no change",
			input: []int{5, 4, 3, 2, 1},
			less: func(a, b int) bool {
				return a > b
			},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:  "descending",
			input: []int{5, 2, 4, 1, 3},
			less: func(a, b int) bool {
				return a > b
			},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "empty slice",
			input:    []int{},
			less:     func(a, b int) bool { return a < b },
			expected: []int{},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     func(a, b int) bool { return a < b },
			expected: nil,
		},
		{
			name:     "single element",
			input:    []int{42},
			less:     func(a, b int) bool { return a < b },
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortBy(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortBy_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		less     func(string, string) bool
		expected []string
	}{
		{
			name:  "ascending",
			input: []string{"a", "b", "c", "d", "e"},
			less: func(a, b string) bool {
				return a < b
			},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:  "descending",
			input: []string{"e", "d", "c", "b", "a"},
			less: func(a, b string) bool {
				return a > b
			},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:  "same values",
			input: []string{"a", "a", "a", "a", "a"},
			less: func(a, b string) bool {
				return a < b
			},
			expected: []string{"a", "a", "a", "a", "a"},
		},
		{
			name:  "sorted by ascending, no change",
			input: []string{"a", "b", "c", "d", "e"},
			less: func(a, b string) bool {
				return a < b
			},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:  "sorted by descending, no change",
			input: []string{"e", "d", "c", "b", "a"},
			less: func(a, b string) bool {
				return a > b
			},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			less:     func(a, b string) bool { return a < b },
			expected: []string{},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     func(a, b string) bool { return a < b },
			expected: nil,
		},
		{
			name:     "single element",
			input:    []string{"a"},
			less:     func(a, b string) bool { return a < b },
			expected: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortBy(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortBy_Struct(t *testing.T) {
	type User struct {
		Name string
	}
	tests := []struct {
		name     string
		input    []User
		less     func(User, User) bool
		expected []User
	}{
		{
			name:  "ascending",
			input: []User{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}},
			less: func(a, b User) bool {
				return a.Name < b.Name
			},
			expected: []User{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}},
		},
		{
			name:  "descending",
			input: []User{{"e"}, {"d"}, {"c"}, {"b"}, {"a"}},
			less: func(a, b User) bool {
				return a.Name > b.Name
			},
			expected: []User{{"e"}, {"d"}, {"c"}, {"b"}, {"a"}},
		},
		{
			name:  "same values",
			input: []User{{"a"}, {"a"}, {"a"}, {"a"}, {"a"}},
			less: func(a, b User) bool {
				return a.Name < b.Name
			},
			expected: []User{{"a"}, {"a"}, {"a"}, {"a"}, {"a"}},
		},
		{
			name:  "sorted by ascending, no change",
			input: []User{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}},
			less: func(a, b User) bool {
				return a.Name < b.Name
			},
			expected: []User{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}},
		},
		{
			name:  "sorted by descending, no change",
			input: []User{{"e"}, {"d"}, {"c"}, {"b"}, {"a"}},
			less: func(a, b User) bool {
				return a.Name > b.Name
			},
			expected: []User{{"e"}, {"d"}, {"c"}, {"b"}, {"a"}},
		},
		{
			name:     "empty slice",
			input:    []User{},
			less:     func(a, b User) bool { return a.Name < b.Name },
			expected: []User{},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     func(a, b User) bool { return a.Name < b.Name },
			expected: nil,
		},
		{
			name:     "single element",
			input:    []User{{"a"}},
			less:     func(a, b User) bool { return a.Name < b.Name },
			expected: []User{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortBy(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestStableSortBy(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	tests := []struct {
		name     string
		input    []User
		less     func(User, User) bool
		expected []User
	}{
		{
			name: "stable sort by age keeps original asc order for equal ages",
			input: []User{
				{Name: "Alice", Age: 20},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 20},
				{Name: "David", Age: 30},
			},
			less: func(a, b User) bool {
				return a.Age < b.Age
			},
			expected: []User{
				// Age 20 group — original order preserved
				{Name: "Alice", Age: 20},
				{Name: "Charlie", Age: 20},

				// Age 30 group — original order preserved
				{Name: "Bob", Age: 30},
				{Name: "David", Age: 30},
			},
		},
		{
			name: "stable sort by age keeps original desc order for equal ages",
			input: []User{
				{Name: "Alice", Age: 20},
				{Name: "Bob", Age: 30},
				{Name: "Charlie", Age: 20},
				{Name: "David", Age: 30},
			},
			less: func(a, b User) bool {
				return a.Age > b.Age
			},
			expected: []User{
				// Age 30 group — original order preserved
				{Name: "Bob", Age: 30},
				{Name: "David", Age: 30},

				// Age 20 group — original order preserved
				{Name: "Alice", Age: 20},
				{Name: "Charlie", Age: 20},
			},
		},
		{
			name: "all equal values",
			input: []User{
				{Name: "Alice", Age: 20},
				{Name: "Bob", Age: 20},
				{Name: "Charlie", Age: 20},
			},
			less: func(a, b User) bool {
				return a.Age < b.Age
			},
			expected: []User{
				{Name: "Alice", Age: 20},
				{Name: "Bob", Age: 20},
				{Name: "Charlie", Age: 20},
			},
		},
		{
			name:     "empty slice",
			input:    []User{},
			less:     func(a, b User) bool { return a.Age < b.Age },
			expected: []User{},
		},
		{
			name:     "nil slice",
			input:    nil,
			less:     func(a, b User) bool { return a.Age < b.Age },
			expected: nil,
		},
		{
			name:     "single element",
			input:    []User{{"Alice", 20}},
			less:     func(a, b User) bool { return a.Age < b.Age },
			expected: []User{{"Alice", 20}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.StableSortBy(tt.input, tt.less)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortAsc_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "unsorted slice",
			input:    []int{5, 2, 4, 1, 3},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sorted asc slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sorted desc slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all same values",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortAsc(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortAsc_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "unsorted slice",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "sorted asc slice",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "sorted desc slice",
			input:    []string{"e", "d", "c", "b", "a"},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "all same values",
			input:    []string{"a", "a", "a", "a", "a"},
			expected: []string{"a", "a", "a", "a", "a"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortAsc(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortDesc_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "unsorted slice",
			input:    []int{5, 2, 4, 1, 3},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "sorted asc slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "sorted desc slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "all same values",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortDesc(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSortDesc_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "unsorted slice",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "sorted asc slice",
			input:    []string{"a", "b", "c", "d", "e"},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "sorted desc slice",
			input:    []string{"e", "d", "c", "b", "a"},
			expected: []string{"e", "d", "c", "b", "a"},
		},
		{
			name:     "all same values",
			input:    []string{"a", "a", "a", "a", "a"},
			expected: []string{"a", "a", "a", "a", "a"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "single element",
			input:    []string{"a"},
			expected: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SortDesc(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
