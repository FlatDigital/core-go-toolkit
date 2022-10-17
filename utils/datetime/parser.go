package datetime

import "time"

const (
	// ISO8601ParseDateFormat defines ISI8601 date time format
	ISO8601ParseDateFormat = "2006-01-02T15:04:05.000Z07:00"
)

type (
	// Parser defines datime parser functionality
	Parser interface {
		Time(value string) (time.Time, error)
		String(time time.Time) string
	}

	parser struct {
		Parser
		parserDateFormat string
	}
)

// NewParser returns an new instance of parser functionality based on specific parser date format
func NewParser(parserDateFormat string) Parser {
	return &parser{
		parserDateFormat: parserDateFormat,
	}
}

// Time converts an datetime string value to a time based on parser date format
func (parser *parser) Time(value string) (time.Time, error) {
	datetime, err := time.Parse(parser.parserDateFormat, value)
	if err != nil {
		return ZeroDate(), err
	}
	datetime = datetime.UTC()
	return datetime, nil
}

// String converts an time to a datetime string value based on parser date format
func (parser *parser) String(time time.Time) string {
	return time.Format(parser.parserDateFormat)
}

// ZeroDate Returns a zero value timestamp
func ZeroDate() time.Time {
	return time.Time{}
}
