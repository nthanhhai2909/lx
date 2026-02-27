package lxtuples_test

import (
	"reflect"
	"testing"

	"github.com/nthanhhai2909/lx/lxtuples"
)

func TestPair_Basics(t *testing.T) {
	tests := []struct {
		name  string
		p     lxtuples.Pair[int, string]
		wantI int
		wantS string
	}{
		{name: "pair1", p: lxtuples.Pair[int, string]{First: 1, Second: "a"}, wantI: 1, wantS: "a"},
		{name: "pair2", p: lxtuples.Pair[int, string]{First: 2, Second: "b"}, wantI: 2, wantS: "b"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.p.First != tt.wantI || tt.p.Second != tt.wantS {
				t.Fatalf("Pair fields = (%v, %v); want (%v, %v)", tt.p.First, tt.p.Second, tt.wantI, tt.wantS)
			}
			// ensure zero value
			var zero lxtuples.Pair[int, string]
			if !reflect.DeepEqual(zero, lxtuples.Pair[int, string]{}) {
				t.Fatalf("zero Pair not empty: %v", zero)
			}
		})
	}
}
