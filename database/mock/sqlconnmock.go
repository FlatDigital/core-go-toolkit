package sqlmock

import (
	"context"
	"database/sql"
)

// SQLConnMock mock struct
type SQLConnMock struct {
	patchClose []outputForConnClose
}

// NewConnMockService return mock for database/sql
func NewConnMockService() *SQLConnMock {
	sqlConnMock := SQLConnMock{
		patchClose: make([]outputForConnClose, 0),
	}

	return &sqlConnMock
}

type (
	outputForConnClose struct {
		outputError error
	}
)

// PatchClose patches the funcion Close
func (mock *SQLConnMock) PatchClose(outputErr error) {
	output := outputForConnClose{
		outputError: outputErr,
	}

	mock.patchClose = append(mock.patchClose, output)
}

// Close mocks the real implementation of Close for the database/sql/conn
func (mock *SQLConnMock) Close() error {
	if len(mock.patchClose) == 0 {
		panic("Mock not available for SQLConnMock.Close")
	}

	output := mock.patchClose[0]
	// dequeue
	mock.patchClose = mock.patchClose[1:]

	return output.outputError
}

// BeginTx mocks the real implementation of BeginTx for the database/sql/conn
func (mock *SQLConnMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	panic("TODO: Implement mock for sql.conn.BeginTx")
}

// ExecContext mocks the real implementation of ExecContext for the database/sql/conn
func (mock *SQLConnMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("TODO: Implement mock for sql.conn.ExecContext")
}

// PingContext mocks the real implementation of PingContext for the database/sql/conn
func (mock *SQLConnMock) PingContext(ctx context.Context) error {
	panic("TODO: Implement mock for sql.conn.PingContext")
}

// PrepareContext mocks the real implementation of PrepareContext for the database/sql/conn
func (mock *SQLConnMock) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	panic("TODO: Implement mock for sql.conn.PrepareContext")
}

// QueryContext mocks the real implementation of QueryContext for the database/sql/conn
func (mock *SQLConnMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("TODO: Implement mock for sql.conn.QueryContext")
}

// QueryRowContext mocks the real implementation of QueryRowContext for the database/sql/conn
func (mock *SQLConnMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.conn.QueryRowContext")
}

// Raw mocks the real implementation of Raw for the database/sql/conn
func (mock *SQLConnMock) Raw(f func(driverConn interface{}) error) (err error) {
	panic("TODO: Implement mock for sql.conn.Raw")
}
