package lxtime

import "time"

// IsTomorrow returns true if the given time is tomorrow.
// It compares the date portion of the time (Year, Month, Day) with tomorrow's date.
// The comparison is done in the time's local timezone.
//
// Example:
//
//	t := time.Now().AddDate(0, 0, 1)
//	if lxtime.IsTomorrow(t) {
//		// t is tomorrow
//	}
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().AddDate(0, 0, 1)
	y1, m1, d1 := t.Date()
	y2, m2, d2 := tomorrow.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
