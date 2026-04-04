package lxtime

import "time"

// IsToday returns true if the given time is today.
// It compares the date portion of the time (Year, Month, Day) with today's date.
// The comparison is done in the time's local timezone.
//
// Example:
//
//	t := time.Now()
//	if lxtime.IsToday(t) {
//		// t is today
//	}
func IsToday(t time.Time) bool {
	now := time.Now()
	y1, m1, d1 := t.Date()
	y2, m2, d2 := now.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
