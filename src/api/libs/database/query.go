package database

// Deprecated: use database/Database interface instead

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	// QueryType defines the type of the query
	QueryType string

	// Query allows you to define a query, add params and then obtain the statement and the full params list
	Query struct {
		queryType       QueryType
		stmtBase        string
		stmtPlaceholder string
		stmtBegin       string
		stmtEnd         string
		stmtValue       string
		params          []interface{}
		valueLength     int
		paramStr        string
		paramsQty       int
		forUpdate       bool
	}
)

const (
	// QueryTypePlain initializes a plain query
	QueryTypePlain QueryType = "plain"

	// QueryTypePlaceholder replaces the values inside the query
	QueryTypePlaceholder QueryType = "placeholder"

	// QueryTypeBeginEnd Defines the begin and the end of the query
	QueryTypeBeginEnd QueryType = "begin_end"

	queryParamsConcatenate = ","
	queryEnd               = ";"
)

// NewQueryPlainMustWork returns a query of type QueryTypePlain
func NewQueryPlainMustWork(stmt string) *Query {
	// Build the query
	query, err := NewQueryPlain(stmt)
	if err != nil {
		// This must not fail, because we are initializing a new query with constants.
		// If it fails, it's like a RegexMustCompile panic.
		panic(err)
	}

	// done
	return query
}

// NewQueryPlain returns a query of type QueryTypePlain
func NewQueryPlain(stmt string) (*Query, error) {
	// Calculate value length
	valueLength, err := calculateValueLength(stmt)
	if err != nil {
		return nil, err
	}

	// done
	return &Query{
		queryType:   QueryTypePlain,
		stmtBase:    stmt,
		valueLength: valueLength,
	}, nil
}

// NewQueryPlaceholderMustWork returns a query of type QueryTypePlaceholder
func NewQueryPlaceholderMustWork(stmtBase string, stmtPlaceholder string, stmtValue string) *Query {
	// Build the query
	query, err := NewQueryPlaceholder(stmtBase, stmtPlaceholder, stmtValue)
	if err != nil {
		// This must not fail, because we are initializing a new query with constants.
		// If it fails, it's like a RegexMustCompile panic.
		panic(err)
	}

	// done
	return query
}

// NewQueryPlaceholder returns a query of type QueryTypePlaceholder
func NewQueryPlaceholder(stmtBase string, stmtPlaceholder string, stmtValue string) (*Query, error) {
	// Calculate value length
	valueLength, err := calculateValueLength(stmtValue)
	if err != nil {
		return nil, err
	}

	// How many times appears the placeholder into the base? (must be 1)
	cnt := strings.Count(stmtBase, stmtPlaceholder)
	if cnt != 1 {
		return nil, fmt.Errorf("the placeholder appears %d times into base query (only 1 is allowed)", cnt)
	}

	// done
	return &Query{
		queryType:       QueryTypePlaceholder,
		stmtBase:        stmtBase,
		stmtPlaceholder: stmtPlaceholder,
		stmtValue:       stmtValue,
		valueLength:     valueLength,
	}, nil
}

// NewQueryBeginEndMustWork returns a query of type QueryTypePlaceholder
func NewQueryBeginEndMustWork(stmtBase string, stmtPlaceholder string, stmtValue string) *Query {
	// Build the query
	query, err := NewQueryBeginEnd(stmtBase, stmtPlaceholder, stmtValue)
	if err != nil {
		// This must not fail, because we are initializing a new query with constants.
		// If it fails, it's like a RegexMustCompile panic.
		panic(err)
	}

	// done
	return query
}

// NewQueryBeginEnd returns a query of type QueryTypeBeginEnd
func NewQueryBeginEnd(stmtBegin string, stmtEnd string, stmtValue string) (*Query, error) {
	// Calculate value length
	valueLength, err := calculateValueLength(stmtValue)
	if err != nil {
		return nil, err
	}

	// Complete end with a semicolon (if empty)
	if stmtEnd == "" {
		stmtEnd = queryEnd
	}

	// done
	return &Query{
		queryType:   QueryTypeBeginEnd,
		stmtBegin:   stmtBegin,
		stmtEnd:     stmtEnd,
		stmtValue:   stmtValue,
		valueLength: valueLength,
	}, nil
}

// AddParams adds params to this query
func (q *Query) AddParams(values []interface{}) error {
	// We have a valid quantity of params?
	if q.valueLength != len(values) {
		return fmt.Errorf("the allowed length for this query is %d, but you have sent a %d params",
			q.valueLength, len(values))
	}

	// We have a plain query?
	if q.queryType == QueryTypePlain {
		// Control for plain queries
		if q.paramsQty > 0 {
			return fmt.Errorf("unable to add more than a group of params to a plain query")
		}
	} else {
		// Concatenate params string
		if q.paramsQty > 0 {
			q.paramStr = q.paramStr + queryParamsConcatenate
		}
		q.paramStr = q.paramStr + q.stmtValue
	}

	// Append to params
	q.params = append(q.params, values...)

	// Increment qty
	q.paramsQty++

	// done
	return nil
}

// Statement returns the statement to be used as the query string
func (q *Query) Statement() (string, error) {
	// We have added params?
	if q.paramsQty == 0 {
		return "", fmt.Errorf("you haven't added any params to this query")
	}

	// Init stmt
	var stmt string

	// Switch by query type
	switch q.queryType {
	case QueryTypeBeginEnd:
		stmt = q.stmtBegin + q.paramStr + q.stmtEnd
	case QueryTypePlaceholder:
		stmt = strings.Replace(q.stmtBase, q.stmtPlaceholder, q.paramStr, 1)
	case QueryTypePlain:
		stmt = q.stmtBase
	default:
		return "", fmt.Errorf("unimplemented queryType: %s", q.queryType)
	}

	// Add a "FOR UPDATE" at the end of the query if we have a true forUpdate flag.
	if q.IsForUpdate() {
		stmt = regexp.MustCompile(`(?i)(FOR UPDATE|)(;|)$`).ReplaceAllString(strings.Trim(stmt, " "), " FOR UPDATE$2")
	}

	// done
	return stmt, nil
}

// Params returns the params to be used to execute the query
func (q *Query) Params() ([]interface{}, error) {
	// We have added params?
	if q.paramsQty == 0 {
		return nil, fmt.Errorf("you haven't added any params to this query")
	}
	// done
	return q.params, nil
}

// GetStatementAndParams returns the statement to be used as the query string and the params
func (q *Query) GetStatementAndParams() (string, []interface{}, error) {
	// Get statement and params with different errors
	statement, err1 := q.Statement()
	params, err2 := q.Params()
	if err1 != nil || err2 != nil {
		return "", nil, fmt.Errorf("unable to obtain statement or params for query (statement error: %s / params error: %s)",
			err1, err2)
	}

	// done
	return statement, params, nil
}

// ForUpdate sets the forUpdate flag of the query
func (q *Query) ForUpdate(forUpdate bool) {
	q.forUpdate = forUpdate
}

// IsForUpdate returns the ForUpdate flag of the query
func (q *Query) IsForUpdate() bool {
	return q.forUpdate
}

// calculateValueLength validates the value, and returns the value length
func calculateValueLength(stmtValue string) (int, error) {
	valueLength := strings.Count(stmtValue, "?")
	if valueLength <= 0 {
		return 0, fmt.Errorf("unable to find any ? in the value placeholder")
	}
	return valueLength, nil
}
