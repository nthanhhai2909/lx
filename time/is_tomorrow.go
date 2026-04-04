package lxtime

import "time"

// IsTomorrow returns true if the given time is tomorrow (in UTC).
// This correctly handles timezones by comparing dates in UTC.
// Near midnight, if t is in a different timezone than the system, this ensures correct results.
//
// Example:
//
//	t := time.Now().AddDate(0, 0, 1)
//	if lxtime.IsTomorrow(t) {
//		// t is tomorrow
//	}
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().UTC().AddDate(0, 0, 1)
	// Convert both to UTC to ensure consistent timezone comparison
	y1, m1, d1 := t.UTC().Date()
	y2, m2, d2 := tomorrow.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
