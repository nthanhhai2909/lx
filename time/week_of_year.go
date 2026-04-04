package lxtime

import "time"

// WeekOfYear returns the week number of the year for the given time.
// The week is defined as Monday through Sunday.
// Returns a value from 1 to 53.
// The first week is the one containing the first Monday of the year.
// Days before the first Monday are considered part of week 0 (but we return 1 for consistency).
//
// Example:
//
//	t := time.Date(2026, 1, 10, 10, 30, 0, 0, time.UTC) // Saturday in week 2
//	week := lxtime.WeekOfYear(t)
//	// week: 2
//
//	t2 := time.Date(2026, 1, 5, 10, 30, 0, 0, time.UTC) // Monday, first week
//	week2 := lxtime.WeekOfYear(t2)
//	// week2: 1
func WeekOfYear(t time.Time) int {
	// Get the start of the year
	year := t.Year()
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())

	// Get the start of the current week
	weekStart := StartOfWeek(t)

	// Get the start of the first week of the year
	// First, find the first Monday of the year
	firstWeekStart := StartOfWeek(yearStart)

	// If the first Monday is in the previous year, move to the next Monday
	if firstWeekStart.Year() < year {
		firstWeekStart = firstWeekStart.AddDate(0, 0, 7)
	}

	// Calculate days difference between week start and first week start
	// Add 1 because we count weeks starting from 1
	daysDiff := int(weekStart.Sub(firstWeekStart).Hours() / 24)
	weekNumber := daysDiff/7 + 1

	return weekNumber
}
