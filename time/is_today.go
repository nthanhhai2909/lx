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
	return isSameDay(t, now)
}

// isSameDay is a helper function that checks if two times are on the same day
func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
