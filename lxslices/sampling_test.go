package lxslices_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxslices"
)

func TestSample_Int(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "single element",
			slice: []int{42},
		},
		{
			name:  "multiple elements",
			slice: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "large slice",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := lxslices.Sample(tt.slice)

			if !found {
				t.Errorf("Sample() returned found=false for non-empty slice %v", tt.slice)
				return
			}
			// Check that result is one of the elements in the slice
			inSlice := false
			for _, v := range tt.slice {
				if v == result {
					inSlice = true
					break
				}
			}
			if !inSlice {
				t.Errorf("Sample() = %v; not found in slice %v", result, tt.slice)
			}
		})
	}
}

func TestSample_EmptyAndNil(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "empty slice",
			slice: []int{},
		},
		{
			name:  "nil slice",
			slice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := lxslices.Sample(tt.slice)
			if found {
				t.Errorf("Sample() returned found=true for empty/nil slice; got value %v", result)
			}
			if result != 0 {
				t.Errorf("Sample() = %v; want zero value 0 for empty/nil slice", result)
			}
		})
	}
}

func TestSample_String(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
	}{
		{
			name:  "single element",
			slice: []string{"hello"},
		},
		{
			name:  "multiple elements",
			slice: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := lxslices.Sample(tt.slice)

			if !found {
				t.Errorf("Sample() returned found=false for non-empty slice %v", tt.slice)
				return
			}
			inSlice := false
			for _, v := range tt.slice {
				if v == result {
					inSlice = true
					break
				}
			}
			if !inSlice {
				t.Errorf("Sample() = %v; not found in slice %v", result, tt.slice)
			}
		})
	}
}

func TestSampleN_Int(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		n           int
		expectedLen int
	}{
		{
			name:        "sample 3 from 5",
			slice:       []int{1, 2, 3, 4, 5},
			n:           3,
			expectedLen: 3,
		},
		{
			name:        "sample 1 from many",
			slice:       []int{1, 2, 3, 4, 5},
			n:           1,
			expectedLen: 1,
		},
		{
			name:        "sample all elements",
			slice:       []int{1, 2, 3},
			n:           3,
			expectedLen: 3,
		},
		{
			name:        "n greater than length",
			slice:       []int{1, 2, 3},
			n:           10,
			expectedLen: 3,
		},
		{
			name:        "n equals zero",
			slice:       []int{1, 2, 3},
			n:           0,
			expectedLen: 0,
		},
		{
			name:        "n negative",
			slice:       []int{1, 2, 3},
			n:           -5,
			expectedLen: 0,
		},
		{
			name:        "empty slice",
			slice:       []int{},
			n:           5,
			expectedLen: 0,
		},
		{
			name:        "nil slice",
			slice:       nil,
			n:           5,
			expectedLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SampleN(tt.slice, tt.n)

			// Check length
			if len(result) != tt.expectedLen {
				t.Errorf("SampleN() length = %v; want %v", len(result), tt.expectedLen)
				return
			}

			// Check all elements are from original slice
			for _, v := range result {
				found := false
				for _, orig := range tt.slice {
					if v == orig {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("SampleN() contains %v which is not in original slice %v", v, tt.slice)
				}
			}

			// Check no duplicates (sampling without replacement)
			if len(result) > 0 {
				seen := make(map[int]bool)
				for _, v := range result {
					if seen[v] {
						t.Errorf("SampleN() contains duplicate: %v", v)
					}
					seen[v] = true
				}
			}
		})
	}
}

func TestSampleN_String(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		n           int
		expectedLen int
	}{
		{
			name:        "sample 2 from 4",
			slice:       []string{"a", "b", "c", "d"},
			n:           2,
			expectedLen: 2,
		},
		{
			name:        "sample all",
			slice:       []string{"x", "y", "z"},
			n:           3,
			expectedLen: 3,
		},
		{
			name:        "n greater than length",
			slice:       []string{"a", "b"},
			n:           5,
			expectedLen: 2,
		},
		{
			name:        "empty slice",
			slice:       []string{},
			n:           3,
			expectedLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SampleN(tt.slice, tt.n)

			if len(result) != tt.expectedLen {
				t.Errorf("SampleN() length = %v; want %v", len(result), tt.expectedLen)
				return
			}

			for _, v := range result {
				found := false
				for _, orig := range tt.slice {
					if v == orig {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("SampleN() contains %v which is not in original slice %v", v, tt.slice)
				}
			}

			// Check no duplicates
			if len(result) > 0 {
				seen := make(map[string]bool)
				for _, v := range result {
					if seen[v] {
						t.Errorf("SampleN() contains duplicate: %v", v)
					}
					seen[v] = true
				}
			}
		})
	}
}

func TestSampleN_Struct(t *testing.T) {
	type Item struct {
		ID   int
		Name string
	}

	tests := []struct {
		name        string
		slice       []Item
		n           int
		expectedLen int
	}{
		{
			name: "sample 2 from 4",
			slice: []Item{
				{1, "a"},
				{2, "b"},
				{3, "c"},
				{4, "d"},
			},
			n:           2,
			expectedLen: 2,
		},
		{
			name: "sample all",
			slice: []Item{
				{1, "x"},
				{2, "y"},
			},
			n:           2,
			expectedLen: 2,
		},
		{
			name: "n greater than length",
			slice: []Item{
				{1, "a"},
			},
			n:           5,
			expectedLen: 1,
		},
		{
			name:        "empty slice",
			slice:       []Item{},
			n:           3,
			expectedLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lxslices.SampleN(tt.slice, tt.n)

			if len(result) != tt.expectedLen {
				t.Errorf("SampleN() length = %v; want %v", len(result), tt.expectedLen)
				return
			}

			for _, v := range result {
				found := false
				for _, orig := range tt.slice {
					if v == orig {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("SampleN() contains %v which is not in original slice %v", v, tt.slice)
				}
			}

			// Check no duplicates
			if len(result) > 0 {
				seen := make(map[Item]bool)
				for _, v := range result {
					if seen[v] {
						t.Errorf("SampleN() contains duplicate: %v", v)
					}
					seen[v] = true
				}
			}
		})
	}
}

// TestSampleEmptyAndSingle verifies Sample returns (zero, false) for empty/nil
// and (element, true) for a single-element slice.
func TestSampleEmptyAndSingle(t *testing.T) {
// nil input
v, ok := lxslices.Sample[int](nil)
if ok || v != 0 {
t.Errorf("Sample(nil) = (%v, %v); want (0, false)", v, ok)
}

// empty non-nil input
v, ok = lxslices.Sample([]int{})
if ok || v != 0 {
t.Errorf("Sample([]int{}) = (%v, %v); want (0, false)", v, ok)
}

// single element
v, ok = lxslices.Sample([]int{99})
if !ok || v != 99 {
t.Errorf("Sample([]int{99}) = (%v, %v); want (99, true)", v, ok)
}
}

// TestSampleNBounds verifies SampleN nil/empty/n-bounds semantics.
func TestSampleNBounds(t *testing.T) {
// nil input → nil output
result := lxslices.SampleN[int](nil, 3)
if result != nil {
t.Errorf("SampleN(nil, 3) = %v; want nil", result)
}

// non-nil empty input → non-nil empty output
result = lxslices.SampleN([]int{}, 3)
if result == nil {
t.Error("SampleN([]int{}, 3) = nil; want non-nil empty slice")
}
if len(result) != 0 {
t.Errorf("SampleN([]int{}, 3) length = %d; want 0", len(result))
}

// n == 0 → empty non-nil
result = lxslices.SampleN([]int{1, 2, 3}, 0)
if result == nil {
t.Error("SampleN(slice, 0) = nil; want non-nil empty slice")
}
if len(result) != 0 {
t.Errorf("SampleN(slice, 0) length = %d; want 0", len(result))
}

// n > len(slice) → all elements returned (shuffled)
src := []int{1, 2, 3}
result = lxslices.SampleN(src, 10)
if len(result) != len(src) {
t.Errorf("SampleN(src, 10) length = %d; want %d", len(result), len(src))
}
}

// TestSampleNShuffles verifies SampleN returns distinct elements from the input.
func TestSampleNShuffles(t *testing.T) {
src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
n := 5
result := lxslices.SampleN(src, n)
if len(result) != n {
t.Fatalf("SampleN() length = %d; want %d", len(result), n)
}

seen := make(map[int]bool)
for _, v := range result {
if seen[v] {
t.Errorf("SampleN() contains duplicate value %d", v)
}
seen[v] = true

found := false
for _, s := range src {
if s == v {
found = true
break
}
}
if !found {
t.Errorf("SampleN() result contains %d which is not in input", v)
}
}
}
