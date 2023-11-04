package sqlmock

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// SQLConnMock mock struct
type SQLConnMock struct {
	mock.Mock
}

// NewConnMockService return mock for database/sql
func NewConnMockService() *SQLConnMock {
	return &SQLConnMock{
		Mock: mock.Mock{},
	}
}

// PatchClose patches the funcion Close
func (mock *SQLConnMock) PatchClose(outputErr error) {
	mock.On("Close").Return(outputErr).Once()
}

// Close mocks the real implementation of Close for the database/sql/conn
func (mock *SQLConnMock) Close() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
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
