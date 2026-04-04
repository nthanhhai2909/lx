package lxtime_test

import (
	"testing"
	"time"

	"github.com/hgapdvn/lx/time"
)

func TestEndOfDay_BasicCases(t *testing.T) {
	tests := []struct {
		name  string
		check func() bool
	}{
		{
			name: "returns end of day",
			check: func() bool {
				input := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				result := lxtime.EndOfDay(input)
				return result.Hour() == 23 && result.Minute() == 59 && result.Second() == 59
			},
		},
		{
			name: "preserves date",
			check: func() bool {
				input := time.Date(2026, 4, 4, 12, 0, 0, 0, time.UTC)
				result := lxtime.EndOfDay(input)
				return result.Year() == 2026 && result.Month() == time.April && result.Day() == 4
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check() {
				t.Errorf("EndOfDay() check failed")
			}
		})
	}
}
