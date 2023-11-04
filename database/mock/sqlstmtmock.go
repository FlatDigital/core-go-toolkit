package sqlmock

import (
	"context"
	"database/sql"

	"github.com/FlatDigital/core-go-toolkit/v2/database/converter"
	"github.com/stretchr/testify/mock"
)

// SQLStmtMock mock struct
type SQLStmtMock struct {
	mock.Mock
}

// NewStmtMockService return mock for database/sql
func NewStmtMockService() *SQLStmtMock {
	return &SQLStmtMock{
		Mock: mock.Mock{},
	}
}

// PatchQuery patches the funcion Query
func (mock *SQLStmtMock) PatchQuery(args []interface{}, rows converter.DBRowser, outputErr error) {
	mock.On("Query", args).Return(rows, outputErr).Once()
}

// Query mocks the real implementation of Query for the database/sql/stmt
func (mock *SQLStmtMock) Query(args ...interface{}) (converter.DBRowser, error) {
	argsMock := mock.Called(args)
	rows, _ := argsMock.Get(0).(converter.DBRowser)
	err, _ := argsMock.Get(1).(error)
	return rows, err
}

// PatchClose patches the funcion Close
func (mock *SQLStmtMock) PatchClose(outputErr error) {
	mock.On("Close").Return(outputErr).Once()
}

// Close mocks the real implementation of Close for the database/sql/stmt
func (mock *SQLStmtMock) Close() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchExec patches the funcion Exec
func (mock *SQLStmtMock) PatchExec(args []interface{}, result sql.Result, outputErr error) {
	mock.On("Exec", args).Return(result, outputErr).Once()
}

// Exec mocks the real implementation of Exec for the database/sql/stmt
func (mock *SQLStmtMock) Exec(args ...interface{}) (sql.Result, error) {
	argsMock := mock.Called(args)
	result, _ := argsMock.Get(0).(sql.Result)
	err, _ := argsMock.Get(1).(error)
	return result, err
}

// PatchExecContext patches the funcion ExecContext
func (mock *SQLStmtMock) PatchExecContext(ctx context.Context, args []interface{}, result sql.Result, outputErr error) {
	mock.On("ExecContext", ctx, args).Return(result, outputErr).Once()
}

// ExecContext mocks the real implementation of ExecContext for the database/sql/stmt
func (mock *SQLStmtMock) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	argsMock := mock.Called(ctx, args)
	result, _ := argsMock.Get(0).(sql.Result)
	err, _ := argsMock.Get(1).(error)
	return result, err
}

// PatchQueryContext patches the funcion Query
func (mock *SQLStmtMock) PatchQueryContext(ctx context.Context, args []interface{},
	rows converter.DBRowser, outputErr error) {
	mock.On("QueryContext", ctx, args).Return(rows, outputErr).Once()
}

// QueryContext mocks the real implementation of QueryContext for the database/sql/stmt
func (mock *SQLStmtMock) QueryContext(ctx context.Context, args ...interface{}) (converter.DBRowser, error) {
	argsMock := mock.Called(ctx, args)
	rows, _ := argsMock.Get(0).(converter.DBRowser)
	err, _ := argsMock.Get(1).(error)
	return rows, err
}

// QueryRow mocks the real implementation of QueryRow for the database/sql/stmt
func (mock *SQLStmtMock) QueryRow(args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.stmt.QueryRow")
}

// QueryRowContext mocks the real implementation of QueryRowContext for the database/sql/stmt
func (mock *SQLStmtMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.stmt.QueryRowContext")
}
