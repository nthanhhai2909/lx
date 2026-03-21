package lxmaps_test

import (
	"testing"

	"github.com/nthanhhai2909/lx/lxmaps"
	"github.com/nthanhhai2909/lx/lxslices"
)

func TestKeys_String(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{"nil map", nil, nil},
		{"empty map", map[string]int{}, []string{}},
		{"single key", map[string]int{"a": 1}, []string{"a"}},
		{"multiple keys small", map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "b", "c"}},
		{"empty string key", map[string]int{"": 0, "x": 1}, []string{"", "x"}},
		{"unicode keys", map[string]int{"こんにちは": 1, "世界": 2}, []string{"こんにちは", "世界"}},
		{"emoji keys", map[string]int{"😊": 1, "emoji": 2}, []string{"😊", "emoji"}},
		{"long strings", map[string]int{"long_string_example_0123456789": 1, "other_long_string_abcdefgh": 2}, []string{"long_string_example_0123456789", "other_long_string_abcdefgh"}},
		{"separator-like keys", map[string]int{"a,b": 1, "c;d": 2}, []string{"a,b", "c;d"}},
		{"numeric-like strings", map[string]int{"1": 1, "2": 2}, []string{"1", "2"}},
		{"special chars", map[string]int{"!@#": 1, "$%": 2}, []string{"!@#", "$%"}},
		{"space keys", map[string]int{" ": 1, "  ": 2}, []string{" ", "  "}},
		{"newline key", map[string]int{"\n": 1, "ok": 2}, []string{"\n", "ok"}},
		{"tab key", map[string]int{"\t": 1, "ok": 2}, []string{"\t", "ok"}},
		{"case sensitive", map[string]int{"Test": 1, "test": 2}, []string{"Test", "test"}},
		{"mixed ascii unicode", map[string]int{"a": 1, "ä": 2, "ö": 3, "ü": 4}, []string{"a", "ä", "ö", "ü"}},
		{"prefix suffix", map[string]int{"pre_": 1, "_suf": 2}, []string{"pre_", "_suf"}},
		{"zwj emoji", map[string]int{"👩‍💻": 1, "👨‍🚀": 2}, []string{"👩‍💻", "👨‍🚀"}},
		{"many keys small", map[string]int{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5}, []string{"k1", "k2", "k3", "k4", "k5"}},
		{"mixed small", map[string]int{"X": 1, "y": 2, "Z": 3}, []string{"X", "y", "Z"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input)

			// Distinguish between expected nil vs empty slice
			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Keys(%v) missing key %q in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestKeys_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    map[int]string
		expected []int
	}{
		{"nil map", nil, nil},
		{"empty map", map[int]string{}, []int{}},
		{"single key", map[int]string{1: "one"}, []int{1}},
		{"multiple small", map[int]string{1: "one", 2: "two", 3: "three"}, []int{1, 2, 3}},
		{"zero key", map[int]string{0: "zero", 1: "one"}, []int{0, 1}},
		{"negative keys", map[int]string{-1: "neg", -2: "neg2"}, []int{-1, -2}},
		{"mixed sign", map[int]string{-5: "a", 5: "b"}, []int{-5, 5}},
		{"large numbers", map[int]string{1000000: "m", 999999: "n"}, []int{1000000, 999999}},
		{"sequential 5", map[int]string{10: "a", 11: "b", 12: "c", 13: "d", 14: "e"}, []int{10, 11, 12, 13, 14}},
		{"sparse keys", map[int]string{2: "a", 100: "b"}, []int{2, 100}},
		{"single large", map[int]string{999: "x"}, []int{999}},
		{"many keys 6", map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}, []int{1, 2, 3, 4, 5, 6}},
		{"mixed small", map[int]string{7: "g", 8: "h", 9: "i"}, []int{7, 8, 9}},
		{"even odd mix", map[int]string{2: "even", 3: "odd", 4: "even"}, []int{2, 3, 4}},
		{"small negatives", map[int]string{-10: "a", -11: "b"}, []int{-10, -11}},
		{"zero only", map[int]string{0: "zero"}, []int{0}},
		{"range 3", map[int]string{20: "a", 21: "b", 22: "c"}, []int{20, 21, 22}},
		{"random small", map[int]string{42: "x", 7: "y"}, []int{42, 7}},
		{"mixed many", map[int]string{100: "x", -1: "y", 0: "z"}, []int{100, -1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input)

			// Distinguish between expected nil vs empty slice
			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Keys(%v) missing key %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}

func TestKeys_Struct(t *testing.T) {
	// define a comparable struct type for keys
	type KS struct {
		I int
		S string
	}

	tests := []struct {
		name     string
		input    map[KS]int
		expected []KS
	}{
		{"nil map", nil, nil},
		{"empty map", map[KS]int{}, []KS{}},
		{"single key", map[KS]int{{I: 1, S: "one"}: 1}, []KS{{I: 1, S: "one"}}},
		{"multiple small", map[KS]int{{I: 1, S: "a"}: 1, {I: 2, S: "b"}: 2}, []KS{{I: 1, S: "a"}, {I: 2, S: "b"}}},
		{"zero and neg", map[KS]int{{I: 0, S: "z"}: 0, {I: -1, S: "n"}: -1}, []KS{{I: 0, S: "z"}, {I: -1, S: "n"}}},
		{"unicode fields", map[KS]int{{I: 10, S: "こんにちは"}: 1, {I: 11, S: "世界"}: 2}, []KS{{I: 10, S: "こんにちは"}, {I: 11, S: "世界"}}},
		{"emoji fields", map[KS]int{{I: 2, S: "😊"}: 1, {I: 3, S: "👋"}: 1}, []KS{{I: 2, S: "😊"}, {I: 3, S: "👋"}}},
		{"many small", map[KS]int{{I: 1, S: "a"}: 1, {I: 2, S: "b"}: 2, {I: 3, S: "c"}: 3}, []KS{{I: 1, S: "a"}, {I: 2, S: "b"}, {I: 3, S: "c"}}},
		{"mixed fields", map[KS]int{{I: 5, S: "X"}: 1, {I: -5, S: "Y"}: 2}, []KS{{I: 5, S: "X"}, {I: -5, S: "Y"}}},
		{"prefix suffix fields", map[KS]int{{I: 7, S: "pre_"}: 1, {I: 8, S: "_suf"}: 2}, []KS{{I: 7, S: "pre_"}, {I: 8, S: "_suf"}}},
		{"long string field", map[KS]int{{I: 100, S: "long_string_example_012345"}: 1}, []KS{{I: 100, S: "long_string_example_012345"}}},
		{"mixed unicode ascii", map[KS]int{{I: 9, S: "a"}: 1, {I: 10, S: "ä"}: 2}, []KS{{I: 9, S: "a"}, {I: 10, S: "ä"}}},
		{"random small", map[KS]int{{I: 42, S: "x"}: 1, {I: 7, S: "y"}: 2}, []KS{{I: 42, S: "x"}, {I: 7, S: "y"}}},
		{"single different", map[KS]int{{I: 999, S: "unique"}: 1}, []KS{{I: 999, S: "unique"}}},
		{"two similar", map[KS]int{{I: 1, S: "same"}: 1, {I: 1, S: "same2"}: 2}, []KS{{I: 1, S: "same"}, {I: 1, S: "same2"}}},
		{"negative many", map[KS]int{{I: -2, S: "a"}: 1, {I: -3, S: "b"}: 2, {I: -4, S: "c"}: 3}, []KS{{I: -2, S: "a"}, {I: -3, S: "b"}, {I: -4, S: "c"}}},
		{"mixed many", map[KS]int{{I: 0, S: "0"}: 1, {I: 1, S: "1"}: 2, {I: 2, S: "2"}: 3}, []KS{{I: 0, S: "0"}, {I: 1, S: "1"}, {I: 2, S: "2"}}},
		{"final small", map[KS]int{{I: 3, S: "end"}: 1, {I: 4, S: "stop"}: 2}, []KS{{I: 3, S: "end"}, {I: 4, S: "stop"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lxmaps.Keys(tt.input)

			// Distinguish between expected nil vs empty slice
			if tt.expected == nil {
				if got != nil {
					t.Fatalf("Keys(%v) = %v; want nil", tt.input, got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Keys(%v) = nil; want non-nil with length %d", tt.input, len(tt.expected))
			}

			if len(got) != len(tt.expected) {
				t.Fatalf("Keys(%v) length = %d; want %d", tt.input, len(got), len(tt.expected))
			}

			for _, exp := range tt.expected {
				if !lxslices.Contains(got, exp) {
					t.Fatalf("Keys(%v) missing key %v in result %v", tt.input, exp, got)
				}
			}
		})
	}
}
