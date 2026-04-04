package lxtime

import "time"

// IsToday returns true if the given time is today (in UTC).
// This correctly handles timezones by comparing dates in UTC.
// Near midnight, if t is in a different timezone than the system, this ensures correct results.
//
// Example:
//
//	t := time.Now()
//	if lxtime.IsToday(t) {
//		// t is today
//	}
func IsToday(t time.Time) bool {
	now := time.Now()
	// Convert both to UTC to ensure consistent timezone comparison
	y1, m1, d1 := t.UTC().Date()
	y2, m2, d2 := now.UTC().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
