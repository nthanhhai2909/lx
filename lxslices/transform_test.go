package lxslices_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestMap_IntToInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(int) int
		expected []int
	}{
		{
			name:     "double each number",
			slice:    []int{1, 2, 3, 4},
			fn:       func(v int) int { return v * 2 },
			expected: []int{2, 4, 6, 8},
		},
		{
			name:     "add 10 to each number",
			slice:    []int{1, 2, 3},
			fn:       func(v int) int { return v + 10 },
			expected: []int{11, 12, 13},
		},
		{
			name:     "square each number",
			slice:    []int{1, 2, 3, 4},
			fn:       func(v int) int { return v * v },
			expected: []int{1, 4, 9, 16},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(v int) int { return v * 2 },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(v int) int { return v * 2 },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_IntToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		fn       func(int) string
		expected []string
	}{
		{
			name:     "convert numbers to stars",
			slice:    []int{1, 2, 3},
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{"*", "**", "***"},
		},
		{
			name:  "convert to even/odd",
			slice: []int{1, 2, 3, 4},
			fn: func(v int) string {
				if v%2 == 0 {
					return "even"
				}
				return "odd"
			},
			expected: []string{"odd", "even", "odd", "even"},
		},
		{
			name:     "empty slice",
			slice:    []int{},
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(v int) string { return strings.Repeat("*", v) },
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StringToInt(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		fn       func(string) int
		expected []int
	}{
		{
			name:     "string length",
			slice:    []string{"a", "bb", "ccc", "dddd"},
			fn:       func(s string) int { return len(s) },
			expected: []int{1, 2, 3, 4},
		},
		{
			name:  "count uppercase letters",
			slice: []string{"Hello", "WORLD", "Go"},
			fn: func(s string) int {
				count := 0
				for _, r := range s {
					if r >= 'A' && r <= 'Z' {
						count++
					}
				}
				return count
			},
			expected: []int{1, 5, 1},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			fn:       func(s string) int { return len(s) },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(s string) int { return len(s) },
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StringToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		fn       func(string) string
		expected []string
	}{
		{
			name:     "convert to uppercase",
			slice:    []string{"hello", "world", "go"},
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{"HELLO", "WORLD", "GO"},
		},
		{
			name:     "add prefix",
			slice:    []string{"apple", "banana"},
			fn:       func(s string) string { return "fruit: " + s },
			expected: []string{"fruit: apple", "fruit: banana"},
		},
		{
			name:     "empty slice",
			slice:    []string{},
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			fn:       func(s string) string { return strings.ToUpper(s) },
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map() = %v; want %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestMap_StructTransform(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	type UserDTO struct {
		ID       int
		FullName string
	}

	tests := []struct {
		name     string
		slice    []User
		fn       func(User) UserDTO
		expected []UserDTO
	}{
		{
			name: "convert User to UserDTO",
			slice: []User{
				{1, "Alice"},
				{2, "Bob"},
			},
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: "Mr/Ms " + u.Name}
			},
			expected: []UserDTO{
				{1, "Mr/Ms Alice"},
				{2, "Mr/Ms Bob"},
			},
		},
		{
			name:  "empty slice",
			slice: []User{},
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: u.Name}
			},
			expected: []UserDTO{},
		},
		{
			name:  "nil slice",
			slice: nil,
			fn: func(u User) UserDTO {
				return UserDTO{ID: u.ID, FullName: u.Name}
			},
			expected: []UserDTO{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Map(tt.slice, tt.fn)
			if len(result) != len(tt.expected) {
				t.Errorf("Map() length = %v; want %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Map()[%d] = %v; want %v", i, result[i], tt.expected[i])
					return
				}
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) []int { return []int{n, n * 2} }
		expected := []int{1, 2, 2, 4, 3, 6}
		result := lxslices.FlatMap(slice, fn)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FlatMap() = %v; want %v", result, expected)
		}
	})

	t.Run("strings", func(t *testing.T) {
		slice := []string{"a", "b"}
		fn := func(s string) []string { return []string{s, s + s} }
		expected := []string{"a", "aa", "b", "bb"}
		result := lxslices.FlatMap(slice, fn)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FlatMap() = %v; want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		fn := func(n int) []int { return []int{n} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		fn := func(n int) []int { return []int{n} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})

	t.Run("mapping to empty", func(t *testing.T) {
		slice := []int{1, 2, 3}
		fn := func(n int) []int { return []int{} }
		result := lxslices.FlatMap(slice, fn)
		if len(result) != 0 {
			t.Errorf("FlatMap() = %v; want empty slice", result)
		}
	})
}

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

func TestReverse(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		result := lxslices.Reverse(slice)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})

	t.Run("strings", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		expected := []string{"c", "b", "a"}
		result := lxslices.Reverse(slice)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		expected := []int{}
		result := lxslices.Reverse(slice)
		if len(result) != 0 {
			t.Errorf("Reverse() = %v; want empty slice", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		expected := []int{}
		result := lxslices.Reverse(slice)
		if len(result) != 0 {
			t.Errorf("Reverse() = %v; want empty slice", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})

	t.Run("single element", func(t *testing.T) {
		slice := []int{42}
		expected := []int{42}
		result := lxslices.Reverse(slice)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})

	t.Run("structs", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		slice := []Person{
			{"Alice", 30},
			{"Bob", 25},
			{"Charlie", 35},
		}
		expected := []Person{
			{"Charlie", 35},
			{"Bob", 25},
			{"Alice", 30},
		}
		result := lxslices.Reverse(slice)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Reverse() = %v; want %v", result, expected)
		}
	})
}

func TestGroupBy(t *testing.T) {
	t.Run("integers by parity", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := map[string][]int{
			"even": {2, 4},
			"odd":  {1, 3, 5},
		}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("strings by length", func(t *testing.T) {
		slice := []string{"a", "bb", "ccc", "dd", "e"}
		expected := map[int][]string{
			1: {"a", "e"},
			2: {"bb", "dd"},
			3: {"ccc"},
		}
		result := lxslices.GroupBy(slice, func(s string) int {
			return len(s)
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		expected := map[string][]int{}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if len(result) != 0 {
			t.Errorf("GroupBy() = %v; want empty map", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("nil slice", func(t *testing.T) {
		var slice []int
		expected := map[string][]int{}
		result := lxslices.GroupBy(slice, func(n int) string {
			if n%2 == 0 {
				return "even"
			}
			return "odd"
		})
		if len(result) != 0 {
			t.Errorf("GroupBy() = %v; want empty map", result)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})

	t.Run("structs by property", func(t *testing.T) {
		type Person struct {
			Name   string
			Active bool
		}
		slice := []Person{
			{"Alice", true},
			{"Bob", false},
			{"Charlie", true},
		}
		expected := map[bool][]Person{
			true:  {{"Alice", true}, {"Charlie", true}},
			false: {{"Bob", false}},
		}
		result := lxslices.GroupBy(slice, func(p Person) bool {
			return p.Active
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("GroupBy() = %v; want %v", result, expected)
		}
	})
}
