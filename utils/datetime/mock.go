package datetime

import "time"

type mock struct {
	Clock
	nowDateTime time.Time
}

// NewClockMock returns a new instance of clock mock struct
func NewClockMock(nowDateTimeString string) (Clock, error) {
	nowDateTime, err := NewParser(ISO8601ParseDateFormat).Time(nowDateTimeString)
	if err != nil {
		return nil, err
	}
	mock := mock{
		nowDateTime: nowDateTime,
	}
	return &mock, nil
}

// Now returns now date time mocked value
func (clock *mock) Now() time.Time {
	return clock.nowDateTime
}
