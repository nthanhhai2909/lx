package lxtime

import "time"

// IsYesterday returns true if the given time is yesterday.
// It compares the date portion of the time (Year, Month, Day) with yesterday's date.
// The comparison is done in the time's local timezone.
//
// Example:
//
//	t := time.Now().AddDate(0, 0, -1)
//	if lxtime.IsYesterday(t) {
//		// t is yesterday
//	}
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return isSameDay(t, yesterday)
}
