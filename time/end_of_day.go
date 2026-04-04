package lxtime

import "time"

// EndOfDay returns the end of the day for the given time.
// Returns the time at 23:59:59.999999999 on the same date.
//
// Example:
//
//	t := time.Date(2026, 4, 4, 15, 30, 45, 123456789, time.UTC)
//	end := lxtime.EndOfDay(t)
//	// end: 2026-04-04 23:59:59.999999999 +0000 UTC
func EndOfDay(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour).Add(24*time.Hour - 1*time.Nanosecond)
}
