package utils

import (
	"testing"
	"time"

	"github.com/FlatDigital/core-go-toolkit/v2/utils/datetime"
	"github.com/stretchr/testify/assert"
)

func Test_TimeUtils(t *testing.T) {
	ass := assert.New(t)

	ass.NotEqual(time.Time{}, Today())
	ass.NotEqual(time.Time{}, Now())
	ass.Equal(time.Time{}, ZeroTime())

	now := Now().Round(time.Duration(1000000))
	parsedNow, err := ParseDate(FormatDate(now))
	ass.Equal(now.UTC(), parsedNow.UTC())
	ass.Nil(err)

	// error test
	_, err = ParseDate("Not a Date pattern")
	ass.NotNil(err)
	ass.Equal(`parsing time "Not a Date pattern" as "2006-01-02T15:04:05.000Z07:00": `+
		`cannot parse "Not a Date pattern" as "2006"`, err.Error())
}

func Test_ParseSQLDate(t *testing.T) {
	ass := assert.New(t)

	ass.NotEqual(time.Time{}, Today())
	ass.NotEqual(time.Time{}, Now())
	ass.Equal(time.Time{}, ZeroTime())

	now := Now().Round(time.Duration(1000000))
	parsedNow, err := ParseDate(FormatDate(now))
	ass.Equal(now.UTC(), parsedNow.UTC())
	ass.Nil(err)

	// error test
	_, err = ParseDate("Not a Date pattern")
	ass.NotNil(err)
	ass.Equal(`parsing time "Not a Date pattern" as "2006-01-02T15:04:05.000Z07:00": `+
		`cannot parse "Not a Date pattern" as "2006"`, err.Error())
}

func Test_DaysSinceTime(t *testing.T) {
	clockMock, _ := datetime.NewClockMock("2018-01-12T18:26:32.000-04:00")
	ass := assert.New(t)

	days := DaysSinceTime(clockMock.Now(), clockMock)
	ass.Equal(uint(0), days)

	days = DaysSinceTime(clockMock.Now().AddDate(0, 0, -5), clockMock)
	ass.Equal(uint(5), days)
}

func TestIsZeroDate(t *testing.T) {
	ass := assert.New(t)

	zeroTime := ZeroTime()

	ass.True(IsZeroDate(zeroTime))
}

func TestFormatDateOnly(t *testing.T) {
	ass := assert.New(t)

	datetimeExpected, err := ParseSQLDate("2020-01-02 15:04:05")
	ass.Nil(err)

	actualDateOnly := FormatDateOnly(datetimeExpected)
	ass.Equal("2020-01-02", actualDateOnly)
}

func TestFormatDateOnlyWithoutDashes(t *testing.T) {
	ass := assert.New(t)

	datetimeExpected, err := ParseSQLDate("2020-01-02 15:04:05")
	ass.Nil(err)

	actualDateOnly := FormatDateOnlyWithoutDashes(datetimeExpected)
	ass.Equal("20200102", actualDateOnly)
}

type AgeTestCandidate struct {
	BirthDate    time.Time
	CheckingTime time.Time
	ExpectedAge  int
}

var AgeTestCandidates = []AgeTestCandidate{
	{time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 10},
	{time.Date(2001, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 8},
	{time.Date(2004, 6, 18, 0, 0, 0, 0, time.UTC), time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 0},
}

func TestAgeAt(t *testing.T) {
	for _, candidate := range AgeTestCandidates {
		gotAge := ageAt(candidate.BirthDate, candidate.CheckingTime)
		if gotAge != candidate.ExpectedAge {
			t.Error(
				"For", candidate.BirthDate,
				"Expected", candidate.ExpectedAge,
				"Got", gotAge,
			)
		}
	}
}

func TestAge(t *testing.T) {
	for _, candidate := range AgeTestCandidates {
		gotAge := Age(candidate.BirthDate)
		if gotAge < candidate.ExpectedAge {
			t.Error(
				"For", candidate.BirthDate,
				"Expected", candidate.ExpectedAge,
				"Got", gotAge,
			)
		}
	}
}

func TestFormatMonthYear(t *testing.T) {
	ass := assert.New(t)

	datetimeExpected, err := ParseSQLDate("2020-01-02 15:04:05")
	ass.Nil(err)

	actualDateOnly := FormatMonthYear(datetimeExpected)
	ass.Equal("02-2020", actualDateOnly)
}

func Test_ParseTimeUS(t *testing.T) {
	ass := assert.New(t)

	// success
	datetimeExpected, err := ParseTimeUS("01-02-2020")
	ass.Nil(err)
	ass.Equal(datetimeExpected.Format(dateUSFormat), "01-02-2020")

	// not a date patter -> error
	datetimeExpected, err = ParseTimeUS("Not a Date pattern")
	ass.Equal(ZeroTime(), datetimeExpected)
	ass.NotNil(err)
	ass.Equal(`parsing time "Not a Date pattern" as "01-02-2006": cannot parse "Not a Date pattern" as "01"`, err.Error())

	// invalid US format -> error
	_, err = ParseTimeUS("24-12-2020")
	ass.NotNil(err)
	ass.Equal(`parsing time "24-12-2020": month out of range`, err.Error())

	// month not valid -> error
	_, err = ParseTimeUS("32-12-2020")
	ass.NotNil(err)
	ass.Equal(`parsing time "32-12-2020": month out of range`, err.Error())

	// day not valid -> error
	_, err = ParseTimeUS("10-33-2020")
	ass.NotNil(err)
	ass.Equal(`parsing time "10-33-2020": day out of range`, err.Error())

	// year not valid -> error
	_, err = ParseTimeUS("10-33-3")
	ass.NotNil(err)
	ass.Equal(`parsing time "10-33-3" as "01-02-2006": cannot parse "3" as "2006"`, err.Error())
}
