package sqlmock

import (
	"context"
	"database/sql"

	"github.com/FlatDigital/core-go-toolkit/src/api/libs/database/converter"
)

// SQLStmtMock mock struct
type SQLStmtMock struct {
	patchStmtQueryContext map[hash][]outputForStmtQueryContext
	patchStmtQuery        map[hash][]outputForStmtQuery
	patchStmtClose        []outputForStmtClose
	patchStmtExec         map[hash][]outputForStmtExec
	patchStmtExecContext  map[hash][]outputForStmtExecContext
}

// NewStmtMockService return mock for database/sql
func NewStmtMockService() *SQLStmtMock {
	sqlStmtMock := SQLStmtMock{
		patchStmtQueryContext: map[hash][]outputForStmtQueryContext{},
		patchStmtQuery:        map[hash][]outputForStmtQuery{},
		patchStmtClose:        make([]outputForStmtClose, 0),
		patchStmtExec:         map[hash][]outputForStmtExec{},
		patchStmtExecContext:  map[hash][]outputForStmtExecContext{},
	}

	return &sqlStmtMock
}

type (
	inputForStmtQuery struct {
		Args []interface{}
	}

	outputForStmtQuery struct {
		rows        converter.DBRowser
		outputError error
	}

	outputForStmtClose struct {
		outputError error
	}

	inputForStmtExec struct {
		Args []interface{}
	}

	outputForStmtExec struct {
		result      sql.Result
		outputError error
	}

	inputForStmtExecContext struct {
		Context context.Context
		Args    []interface{}
	}

	outputForStmtExecContext struct {
		result      sql.Result
		outputError error
	}

	inputForStmtQueryContext struct {
		Context context.Context
		Args    []interface{}
	}

	outputForStmtQueryContext struct {
		rows        converter.DBRowser
		outputError error
	}
)

// PatchQuery patches the funcion Query
func (mock *SQLStmtMock) PatchQuery(args []interface{}, rows converter.DBRowser, outputErr error) {
	input := inputForStmtQuery{
		Args: args,
	}
	hash := toHash(input)

	output := outputForStmtQuery{
		rows:        rows,
		outputError: outputErr,
	}

	mock.patchStmtQuery[hash] = append(mock.patchStmtQuery[hash], output)
}

// Query mocks the real implementation of Query for the database/sql/stmt
func (mock *SQLStmtMock) Query(args ...interface{}) (converter.DBRowser, error) {
	inputStruct := inputForStmtQuery{
		Args: args,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchStmtQuery[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLStmtMock.Query")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchStmtQuery[hash] = mocksArr[1:]

	return output.rows, output.outputError
}

// PatchClose patches the funcion Close
func (mock *SQLStmtMock) PatchClose(outputErr error) {
	output := outputForStmtClose{
		outputError: outputErr,
	}

	mock.patchStmtClose = append(mock.patchStmtClose, output)
}

// Close mocks the real implementation of Close for the database/sql/stmt
func (mock *SQLStmtMock) Close() error {
	if len(mock.patchStmtClose) == 0 {
		panic("Mock not available for SQLStmtMock.Close")
	}

	output := mock.patchStmtClose[0]
	// dequeue
	mock.patchStmtClose = mock.patchStmtClose[1:]

	return output.outputError
}

// PatchExec patches the funcion Exec
func (mock *SQLStmtMock) PatchExec(args []interface{}, result sql.Result, outputErr error) {
	input := inputForStmtExec{
		Args: args,
	}
	hash := toHash(input)

	output := outputForStmtExec{
		result:      result,
		outputError: outputErr,
	}

	mock.patchStmtExec[hash] = append(mock.patchStmtExec[hash], output)
}

// Exec mocks the real implementation of Exec for the database/sql/stmt
func (mock *SQLStmtMock) Exec(args ...interface{}) (sql.Result, error) {
	inputStruct := inputForStmtExec{
		Args: args,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchStmtExec[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLStmtMock.Exec")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchStmtExec[hash] = mocksArr[1:]

	return output.result, output.outputError
}

// PatchExecContext patches the funcion ExecContext
func (mock *SQLStmtMock) PatchExecContext(ctx context.Context, args []interface{}, result sql.Result, outputErr error) {
	input := inputForStmtExecContext{
		Context: ctx,
		Args:    args,
	}
	hash := toHash(input)

	output := outputForStmtExecContext{
		result:      result,
		outputError: outputErr,
	}

	mock.patchStmtExecContext[hash] = append(mock.patchStmtExecContext[hash], output)
}

// ExecContext mocks the real implementation of ExecContext for the database/sql/stmt
func (mock *SQLStmtMock) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	inputStruct := inputForStmtExecContext{
		Context: ctx,
		Args:    args,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchStmtExecContext[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLStmtMock.ExecContext")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchStmtExecContext[hash] = mocksArr[1:]

	return output.result, output.outputError
}

// PatchQueryContext patches the funcion Query
func (mock *SQLStmtMock) PatchQueryContext(ctx context.Context, args []interface{},
	rows converter.DBRowser, outputErr error) {
	input := inputForStmtQueryContext{
		Context: ctx,
		Args:    args,
	}
	hash := toHash(input)

	output := outputForStmtQueryContext{
		rows:        rows,
		outputError: outputErr,
	}

	mock.patchStmtQueryContext[hash] = append(mock.patchStmtQueryContext[hash], output)
}

// QueryContext mocks the real implementation of QueryContext for the database/sql/stmt
func (mock *SQLStmtMock) QueryContext(ctx context.Context, args ...interface{}) (converter.DBRowser, error) {
	inputStruct := inputForStmtQueryContext{
		Context: ctx,
		Args:    args,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchStmtQueryContext[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLStmtMock.QueryContext")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchStmtQueryContext[hash] = mocksArr[1:]

	return output.rows, output.outputError
}

// QueryRow mocks the real implementation of QueryRow for the database/sql/stmt
func (mock *SQLStmtMock) QueryRow(args ...interface{}) *sql.Row {
	panic("TODO: Implement mock for sql.stmt.QueryRow")
}

// QueryRowContext mocks the real implementation of QueryRowContext for the database/sql/stmt
func (mock *SQLStmtMock) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row {
	panic("TODO: Implement mock for sql.stmt.QueryRowContext")
}
