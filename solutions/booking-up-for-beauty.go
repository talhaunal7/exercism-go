// https://exercism.org/tracks/go/exercises/booking-up-for-beauty , time
package booking

import "time"

const (
	scheduleLayout    = "1/2/2006 15:04:05"
	hasPassedLayout   = "January 2, 2006 15:04:05"
	isAfternoonLayout = "Monday, January 2, 2006 15:04:05"
	appointmentLayout = "You have an appointment on Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse(scheduleLayout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse(hasPassedLayout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse(isAfternoonLayout, date)
	return (t.Hour() >= 12 && t.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse(scheduleLayout, date)
	return t.Format(appointmentLayout)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}