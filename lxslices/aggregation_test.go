package lxslices_test

import (
	"math"
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

func TestSum_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected int
	}{
		{
			name:     "sum of positives",
			slice:    []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "single element",
			slice:    []int{42},
			expected: 42,
		},
		{
			name:     "includes negatives",
			slice:    []int{-1, 2, -3, 4},
			expected: 2,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			expected: 0,
		},
		{
			name:     "nil slice",
			slice:    nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Sum(tt.slice)
			if result != tt.expected {
				t.Errorf("Sum() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestSum_Float64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected float64
	}{
		{
			name:     "sum of floats",
			slice:    []float64{1.5, 2.5, 3.0},
			expected: 7.0,
		},
		{
			name:     "single element",
			slice:    []float64{3.14},
			expected: 3.14,
		},
		{
			name:     "empty slice",
			slice:    []float64{},
			expected: 0.0,
		},
		{
			name:     "nil slice",
			slice:    nil,
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.Sum(tt.slice)
			if result != tt.expected {
				t.Errorf("Sum() = %v; want %v", result, tt.expected)
			}
		})
	}
}

func TestMin_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected struct {
			value int
			found bool
		}
	}{
		{
			name:  "min at beginning",
			slice: []int{1, 2, 3},
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:  "min in middle",
			slice: []int{3, 1, 2},
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:  "min at end",
			slice: []int{3, 2, 1},
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:  "duplicate minima returns first value",
			slice: []int{2, 1, 1, 3},
			expected: struct {
				value int
				found bool
			}{1, true},
		},
		{
			name:  "empty slice",
			slice: []int{},
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:  "nil slice",
			slice: nil,
			expected: struct {
				value int
				found bool
			}{0, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Min(tt.slice)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("Min() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestMin_Float64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected struct {
			value float64
			found bool
		}
	}{
		{
			name:  "min at beginning",
			slice: []float64{1.0, 2.0, 3.0},
			expected: struct {
				value float64
				found bool
			}{1.0, true},
		},
		{
			name:  "min in middle",
			slice: []float64{3.0, 1.0, 2.0},
			expected: struct {
				value float64
				found bool
			}{1.0, true},
		},
		{
			name:  "min at end",
			slice: []float64{3.0, 2.0, 1.0},
			expected: struct {
				value float64
				found bool
			}{1.0, true},
		},
		{
			name:  "duplicate minima returns first value",
			slice: []float64{2.0, 1.0, 1.0, 3.0},
			expected: struct {
				value float64
				found bool
			}{1.0, true},
		},
		{
			name:  "empty slice",
			slice: []float64{},
			expected: struct {
				value float64
				found bool
			}{0.0, false},
		},
		{
			name:  "nil slice",
			slice: nil,
			expected: struct {
				value float64
				found bool
			}{0.0, false},
		},
		{
			name:  "NaN first returns NaN",
			slice: []float64{math.NaN(), 1.0, 2.0},
			expected: struct {
				value float64
				found bool
			}{math.NaN(), true},
		},
		{
			name:  "NaN later ignored",
			slice: []float64{1.0, math.NaN(), 2.0},
			expected: struct {
				value float64
				found bool
			}{1.0, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Min(tt.slice)
			if tt.expected.found {
				if math.IsNaN(tt.expected.value) {
					if !found {
						t.Fatalf("Min() found=false; want true")
					}
					if !math.IsNaN(value) {
						t.Errorf("Min() value = %v; want NaN", value)
					}
				} else {
					if value != tt.expected.value || found != tt.expected.found {
						t.Errorf("Min() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
					}
				}
			} else {
				if found {
					t.Errorf("Min() found=true; want false for empty/nil slice")
				}
			}
		})
	}
}

func TestMax_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected struct {
			value int
			found bool
		}
	}{
		{
			name:  "max at beginning",
			slice: []int{3, 2, 1},
			expected: struct {
				value int
				found bool
			}{3, true},
		},
		{
			name:  "max in middle",
			slice: []int{1, 3, 2},
			expected: struct {
				value int
				found bool
			}{3, true},
		},
		{
			name:  "max at end",
			slice: []int{1, 2, 5},
			expected: struct {
				value int
				found bool
			}{5, true},
		},
		{
			name:  "duplicate maxima returns first value",
			slice: []int{5, 3, 5, 2},
			expected: struct {
				value int
				found bool
			}{5, true},
		},
		{
			name:  "empty slice",
			slice: []int{},
			expected: struct {
				value int
				found bool
			}{0, false},
		},
		{
			name:  "nil slice",
			slice: nil,
			expected: struct {
				value int
				found bool
			}{0, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Max(tt.slice)
			if value != tt.expected.value || found != tt.expected.found {
				t.Errorf("Max() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestMax_Float64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected struct {
			value float64
			found bool
		}
	}{
		{
			name:  "max at beginning",
			slice: []float64{3.0, 2.0, 1.0},
			expected: struct {
				value float64
				found bool
			}{3.0, true},
		},
		{
			name:  "max in middle",
			slice: []float64{1.0, 3.0, 2.0},
			expected: struct {
				value float64
				found bool
			}{3.0, true},
		},
		{
			name:  "max at end",
			slice: []float64{1.0, 2.0, 5.0},
			expected: struct {
				value float64
				found bool
			}{5.0, true},
		},
		{
			name:  "duplicate maxima returns first value",
			slice: []float64{5.0, 3.0, 5.0, 2.0},
			expected: struct {
				value float64
				found bool
			}{5.0, true},
		},
		{
			name:  "empty slice",
			slice: []float64{},
			expected: struct {
				value float64
				found bool
			}{0.0, false},
		},
		{
			name:  "nil slice",
			slice: nil,
			expected: struct {
				value float64
				found bool
			}{0.0, false},
		},
		{
			name:  "NaN first returns NaN",
			slice: []float64{math.NaN(), 1.0, 2.0},
			expected: struct {
				value float64
				found bool
			}{math.NaN(), true},
		},
		{
			name:  "NaN later ignored",
			slice: []float64{1.0, math.NaN(), 2.0},
			expected: struct {
				value float64
				found bool
			}{2.0, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := lxslices.Max(tt.slice)
			if tt.expected.found {
				if math.IsNaN(tt.expected.value) {
					if !found {
						t.Fatalf("Max() found=false; want true")
					}
					if !math.IsNaN(value) {
						t.Errorf("Max() value = %v; want NaN", value)
					}
				} else {
					if value != tt.expected.value || found != tt.expected.found {
						t.Errorf("Max() = (%v, %v); want (%v, %v)", value, found, tt.expected.value, tt.expected.found)
					}
				}
			} else {
				if found {
					t.Errorf("Max() found=true; want false for empty/nil slice")
				}
			}
		})
	}
}

func TestAverage_Int(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected struct {
			value float64
			found bool
		}
	}{
		{name: "average of ints", slice: []int{1, 2, 3, 4}, expected: struct {
			value float64
			found bool
		}{2.5, true}},
		{name: "single element", slice: []int{5}, expected: struct {
			value float64
			found bool
		}{5.0, true}},
		{name: "negatives", slice: []int{-1, -2, -3}, expected: struct {
			value float64
			found bool
		}{-2.0, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := lxslices.Average(tt.slice)
			if v != tt.expected.value || ok != tt.expected.found {
				t.Errorf("Average() = (%v, %v); want (%v, %v)", v, ok, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestAverage_Float64(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected struct {
			value float64
			found bool
		}
	}{
		{name: "average floats", slice: []float64{1.5, 2.5, 3.0}, expected: struct {
			value float64
			found bool
		}{7.0 / 3.0, true}},
		{name: "single float", slice: []float64{2.2}, expected: struct {
			value float64
			found bool
		}{2.2, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, ok := lxslices.Average(tt.slice)
			if v != tt.expected.value || ok != tt.expected.found {
				t.Errorf("Average() = (%v, %v); want (%v, %v)", v, ok, tt.expected.value, tt.expected.found)
			}
		})
	}
}

func TestAverage_Empty(t *testing.T) {
	// empty and nil slices should return (0, false)
	if v, ok := lxslices.Average([]int{}); ok || v != 0 {
		t.Errorf("Average(empty int) = (%v, %v); want (0, false)", v, ok)
	}
	if v, ok := lxslices.Average([]float64{}); ok || v != 0 {
		t.Errorf("Average(empty float64) = (%v, %v); want (0, false)", v, ok)
	}
	if v, ok := lxslices.Average([]int(nil)); ok || v != 0 {
		t.Errorf("Average(nil int) = (%v, %v); want (0, false)", v, ok)
	}
}
