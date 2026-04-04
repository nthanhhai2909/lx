package lxtime

import "time"

// EndOfQuarter returns the end of the quarter for the given time.
// Returns the last day of the quarter at 23:59:59.999999999 in the same timezone.
// Q1: March 31, Q2: June 30, Q3: September 30, Q4: December 31
//
// Example:
//
//	t := time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)
//	end := lxtime.EndOfQuarter(t)
//	// end: 2026-06-30 23:59:59.999999999 +0000 UTC (Q2)
func EndOfQuarter(t time.Time) time.Time {
	// Extract year and month
	year := t.Year()
	month := t.Month()

	// Determine the ending month of the quarter
	var endMonth time.Month
	switch {
	case month >= time.January && month <= time.March:
		endMonth = time.March
	case month >= time.April && month <= time.June:
		endMonth = time.June
	case month >= time.July && month <= time.September:
		endMonth = time.September
	default: // October, November, December
		endMonth = time.December
	}

	// Get the last day of the quarter's ending month
	// by going to the first day of the next month and subtracting 1 day
	nextMonth := endMonth + 1
	if nextMonth > time.December {
		nextMonth = time.January
		year++
	}

	firstOfNextMonth := time.Date(year, nextMonth, 1, 0, 0, 0, 0, t.Location())
	lastDayOfQuarter := firstOfNextMonth.AddDate(0, 0, -1)

	// Return with time set to end of day (23:59:59.999999999)
	return time.Date(lastDayOfQuarter.Year(), lastDayOfQuarter.Month(), lastDayOfQuarter.Day(),
		23, 59, 59, 999999999, t.Location())
}
