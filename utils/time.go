package utils

import (
	"fmt"
	"time"

	"github.com/FlatDigital/core-go-toolkit/v2/utils/datetime"
)

var (
	utcZone               = time.FixedZone("UTC", 0)
	iso8601Format         = "2006-01-02T15:04:05.000Z07:00"
	sqlFormat             = "2006-01-02 15:04:05"
	dateOnlyFormat        = "2006-01-02"
	dateOnlyWithoutDashes = "20060102"
	dateOnlyMonthYear     = "02-2006"
	dateUSFormat          = "01-02-2006"
)

// Now Returns a timestamp
func Now() time.Time {
	return time.Now().In(utcZone)
}

// Today returns a datestamp
func Today() time.Time {
	return Date(Now())
}

// Date returns a datestamp from a timestamp
func Date(time time.Time) time.Time {
	year, month, day := time.Date()
	yearStr := fmt.Sprintf("%d", year)
	monthStr := fmt.Sprintf("%d", month)
	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}
	dayStr := fmt.Sprintf("%d", day)
	if len(dayStr) == 1 {
		dayStr = "0" + dayStr
	}

	date, _ := ParseDateOnly(fmt.Sprintf("%s-%s-%s", yearStr, monthStr, dayStr))
	return date
}

// ZeroTime Returns a zero value timestamp
func ZeroTime() time.Time {
	return time.Time{}
}

// IsZeroDate indicates if a date is zero date
func IsZeroDate(date time.Time) bool {
	return date == ZeroTime()
}

//

// FormatDate Returns a string representing a specific datetime formatted in ISO8601
func FormatDate(datetime time.Time) string {
	return datetime.Format(iso8601Format)
}

// FormatCustomDate Returns a string representing a specific datetime formatted in custom format
func FormatCustomDate(datetime time.Time, format string) string {
	return datetime.Format(format)
}

// FormatDateOnlyWithoutDashes Returns a string representing a specific date only string without dashes
func FormatDateOnlyWithoutDashes(datetime time.Time) string {
	return datetime.Format(dateOnlyWithoutDashes)
}

// FormatSQLDate returns a string representing a specific datetime formatted in sql date
func FormatSQLDate(datetime time.Time) string {
	return datetime.Format(sqlFormat)
}

// FormatDateOnly Returns a string representing a specific date only string
func FormatDateOnly(datetime time.Time) string {
	return datetime.Format(dateOnlyFormat)
}

// FormatMonthYear Returns a string representing a date with format "MM-YYYY
func FormatMonthYear(datetime time.Time) string {
	return datetime.Format(dateOnlyMonthYear)
}

// ParseDate Returns a timestamp for the date in string passed
func ParseDate(value string) (time.Time, error) {
	return ParseCustomTime(value, iso8601Format)
}

// ParseSQLDate returns a timestamp from a sql date string
func ParseSQLDate(value string) (time.Time, error) {
	return ParseCustomTime(value, sqlFormat)
}

// ParseDateOnly returns a timestamp from a date only string
func ParseDateOnly(value string) (time.Time, error) {
	return ParseCustomTime(value, dateOnlyFormat)
}

// DaysSinceTime returns the duration between the requested date and Now
func DaysSinceTime(datetime time.Time, clock datetime.Clock) uint {
	duration := clock.Now().Sub(datetime)
	daysDuration := int64(duration.Hours()) / 24
	return uint(daysDuration)
}

//

func ParseCustomTime(value string, format string) (time.Time, error) {
	datetime, err := time.Parse(format, value)
	if err != nil {
		return ZeroTime(), err
	}
	datetime = datetime.UTC()
	return datetime, nil
}

// ElapsedSince returns elapsed time in ms
func ElapsedSince(start time.Time) string {
	return fmt.Sprintf("%d ms", time.Since(start).Nanoseconds()/1000000)
}

// ElapsedSinceFloat returns elapsed time in ms as float64
func ElapsedSinceFloat(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds()) / 1000000.0
}

// DaysInMonth returns the count of days each month
func DaysInMonth(year, month int) int {
	// To compute the number of days in the month, we add a month and subtract a day.
	// This returns the max day of the next month
	firstDateOfCurrentMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	firstDateOfNextMonth := firstDateOfCurrentMonth.AddDate(0, 1, 0)
	lastDateOfCurrentMonth := firstDateOfNextMonth.AddDate(0, 0, -1)
	return lastDateOfCurrentMonth.Day()
}

// Age Calculation
// Age is shorthand for AgeAt(birthDate, time.Now()), and carries the same usage and limitations.
func Age(birthDate time.Time) int {
	return ageAt(birthDate, Now())
}

// AgeAt gets the age of an entity at a certain time.
func ageAt(birthDate time.Time, now time.Time) int {
	// Get the year number change since the player's birth.
	years := now.Year() - birthDate.Year()

	// If the date is before the date of birth, then not that many years have elapsed.
	birthDay := getAdjustedBirthDay(birthDate, now)
	if now.YearDay() < birthDay {
		years--
	}

	return years
}

// Gets the adjusted date of birth to work around leap year differences.
func getAdjustedBirthDay(birthDate time.Time, now time.Time) int {
	birthYearDay := birthDate.YearDay()
	currentDay := now.YearDay()
	if isLeap(birthDate) && !isLeap(now) && birthYearDay >= 60 {
		return birthYearDay - 1
	}
	if isLeap(now) && !isLeap(birthDate) && currentDay >= 60 {
		return birthYearDay + 1
	}
	return birthYearDay
}

// Works out if a time.Time is in a leap year.
func isLeap(date time.Time) bool {
	year := date.Year()
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

// ParseTimeUS Parses mm-dd-yyyy to Time-time in US format
func ParseTimeUS(value string) (time.Time, error) {
	return ParseCustomTime(value, dateUSFormat)
}
