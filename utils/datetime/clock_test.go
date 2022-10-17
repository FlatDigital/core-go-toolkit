package datetime_test

import (
	"testing"
	"time"

	"github.com/FlatDigital/core-go-toolkit/utils/datetime"
	"github.com/stretchr/testify/assert"
)

func Test_AssertNowDatetime(t *testing.T) {
	assertion := assert.New(t)

	clock := datetime.NewClock()

	utcZone := time.FixedZone("UTC", 0)
	expectedNow := time.Now().In(utcZone)

	// Now clock date time interface and now date time runtime different should be less 100 milliseconds
	assertion.Equal(true, expectedNow.Sub(clock.Now()) < 100)
}
