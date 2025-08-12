package utils

import "time"

// DateDifference calculates the difference between two dates in years and months.
// It returns the number of full years and the number of full months remaining.
func DateDifference(startDate, endDate time.Time) (int, int) {
	// Ensure startDate is before endDate
	if startDate.After(endDate) {
		startDate, endDate = endDate, startDate
	}

	// Extract year, month, and day for easier calculation
	y1, m1, d1 := startDate.Date()
	y2, m2, d2 := endDate.Date()

	// Calculate the difference in years and months, ignoring the day
	yearDiff := y2 - y1
	monthDiff := int(m2 - m1)

	// If the end date's day is before the start date's day,
	// we haven't completed the last month yet.
	// For example, Jan 15 to Feb 14 is 0 months and 30 days.
	if d2 < d1 {
		monthDiff--
	}

	// If the month difference is negative, it means we have crossed a year boundary.
	// For example, from Oct 2023 to Feb 2024.
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	return yearDiff, monthDiff
}
