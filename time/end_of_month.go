package lxtime

import "time"

// EndOfMonth returns the end of the month for the given time.
// Returns the last day of the month at 23:59:59.999999999.
//
// Example:
//
//	t := time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)
//	end := lxtime.EndOfMonth(t)
//	// end: 2026-04-30 23:59:59.999999999 +0000 UTC
func EndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1).Add(24*time.Hour - 1*time.Nanosecond)
}
