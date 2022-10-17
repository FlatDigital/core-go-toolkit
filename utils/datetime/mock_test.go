package datetime_test

import (
	"testing"

	"github.com/FlatDigital/core-go-toolkit/utils/datetime"
	"github.com/stretchr/testify/assert"
)

func Test_AssertANowDatimeClockMockValueShouldBeEqualsThanExpectedDatimeValue(t *testing.T) {
	assertion := assert.New(t)

	clockMock, clockMockError := datetime.NewClockMock("2018-01-12T18:26:32.000-04:00")
	assertion.NoError(clockMockError)
	expectedNowDateTime, _ := datetime.NewParser(datetime.ISO8601ParseDateFormat).Time("2018-01-12T18:26:32.000-04:00")

	assertion.Equal(expectedNowDateTime, clockMock.Now())
}

func Test_AssertAnInvalidDateDatimeClockMockValueShouldReturnError(t *testing.T) {
	assertion := assert.New(t)

	clockMock, clockMockError := datetime.NewClockMock("-01-12T18:26:32.000-04:00")
	assertion.Error(clockMockError)
	assertion.Nil(clockMock)
}
