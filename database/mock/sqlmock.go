package sqlmock

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/FlatDigital/core-go-toolkit/v2/database/converter"
	"github.com/stretchr/testify/mock"
)

// SQLMock mock struct
type SQLMock struct {
	mock.Mock
	converter.DBer
}

// NewMockService return mock for database/sql
func NewMockService() *SQLMock {
	return &SQLMock{
		Mock: mock.Mock{},
	}
}

// PatchConn patches the funcion Conn
func (mock *SQLMock) PatchConn(ctx context.Context, outputConn converter.DBConner, outputErr error) {
	mock.On("Conn", ctx).Return(outputConn, outputErr).Once()
}

// Conn mocks the real implementation of Conn for the database/sql
func (mock *SQLMock) Conn(ctx context.Context) (converter.DBConner, error) {
	args := mock.Called(ctx)
	conn, _ := args.Get(0).(converter.DBConner)
	err, _ := args.Get(1).(error)
	return conn, err
}

// PatchPing patches the funcion Ping
func (mock *SQLMock) PatchPing(outputErr error) {
	mock.On("Ping").Return(outputErr).Once()
}

// Ping mocks the real implementation of Ping for the database/sql
func (mock *SQLMock) Ping() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchPingContext patches the funcion PingContext
func (mock *SQLMock) PatchPingContext(ctx context.Context, outputErr error) {
	mock.On("PingContext", ctx).Return(outputErr).Once()
}

// PingContext mocks the real implementation of PingContext for the database/sql
func (mock *SQLMock) PingContext(ctx context.Context) error {
	args := mock.Called(ctx)
	err, _ := args.Get(0).(error)
	return err
}

// PatchClose patches the funcion Close
func (mock *SQLMock) PatchClose(outputErr error) {
	mock.On("Close").Return(outputErr).Once()
}

// Close mocks the real implementation of Close for the database/sql
func (mock *SQLMock) Close() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchBeginTx patches the funcion BeginTx
func (mock *SQLMock) PatchBeginTx(ctx context.Context, opts *sql.TxOptions, tx converter.DBTxer, outputErr error) {
	mock.On("BeginTx", ctx, opts).Return(tx, outputErr).Once()
}

// BeginTx mocks the real implementation of BeginTx for the database/sql
func (mock *SQLMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (converter.DBTxer, error) {
	args := mock.Called(ctx, opts)
	tx, _ := args.Get(0).(converter.DBTxer)
	err, _ := args.Get(1).(error)
	return tx, err
}

// PatchPrepare patches the funcion Prepare
func (mock *SQLMock) PatchPrepare(query string, stmt converter.DBStmter, outputErr error) {
	mock.On("Prepare", query).Return(stmt, outputErr).Once()
}

// Prepare mocks the real implementation of Prepare for the database/sql
func (mock *SQLMock) Prepare(query string) (converter.DBStmter, error) {
	args := mock.Called(query)
	stmt, _ := args.Get(0).(converter.DBStmter)
	err, _ := args.Get(1).(error)
	return stmt, err
}

// PatchStats patches the funcion Stats
func (mock *SQLMock) PatchStats(stats sql.DBStats) {
	mock.On("Stats").Return(stats).Once()
}

// Stats mocks the real implementation of Stats for the database/sql
func (mock *SQLMock) Stats() sql.DBStats {
	args := mock.Called()
	stats, _ := args.Get(0).(sql.DBStats)
	return stats
}

// PatchPrepareContext patches the funcion PrepareContext
func (mock *SQLMock) PatchPrepareContext(ctx context.Context, query string, stmt converter.DBStmter, outputErr error) {
	mock.On("PrepareContext", ctx, query).Return(stmt, outputErr).Once()
}

// PrepareContext mocks the real implementation of PrepareContext for the database/sql
func (mock *SQLMock) PrepareContext(ctx context.Context, query string) (converter.DBStmter, error) {
	args := mock.Called(ctx, query)
	stmt, _ := args.Get(0).(converter.DBStmter)
	err, _ := args.Get(1).(error)
	return stmt, err
}

// Begin mocks the real implementation of Begin for the database/sql
func (mock *SQLMock) Begin() (*sql.Tx, error) {
	panic("TODO: Implement mock for sql.db.Begin")
}

// Driver mocks the real implementation of Driver for the database/sql
func (mock *SQLMock) Driver() driver.Driver {
	panic("TODO: Implement mock for sql.db.Driver")
}

// Exec mocks the real implementation of Exec for the database/sql
func (mock *SQLMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("TODO: Implement mock for sql.db.Exec")
}

// ExecContext mocks the real implementation of ExecContext for the database/sql
func (mock *SQLMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("TODO: Implement mock for sql.db.ExecContext")
}

// Query mocks the real implementation of Query for the database/sql
func (mock *SQLMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	panic("TODO: Implement mock for sql.db.Query")
}

// QueryContext mocks the real implementation of QueryContext for the database/sql
func (mock *SQLMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("TODO: Implement mock for sql.db.QueryContext")
}

// QueryRow mocks the real implementation of QueryRow for the database/sql
func (mock *SQLMock) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.db.QueryRow")
}

// QueryRowContext mocks the real implementation of QueryRowContext for the database/sql
func (mock *SQLMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.db.QueryRowContext")
}

// SetConnMaxLifetime mocks the real implementation of SetConnMaxLifetime for the database/sql
func (mock *SQLMock) SetConnMaxLifetime(d time.Duration) {
	panic("TODO: Implement mock for sql.db.SetConnMaxLifetime")
}

// SetMaxIdleConns mocks the real implementation of SetMaxIdleConns for the database/sql
func (mock *SQLMock) SetMaxIdleConns(n int) {
	panic("TODO: Implement mock for sql.db.SetMaxIdleConns")
}

// SetMaxOpenConns mocks the real implementation of SetMaxOpenConns for the database/sql
func (mock *SQLMock) SetMaxOpenConns(n int) {
	panic("TODO: Implement mock for sql.db.SetMaxOpenConns")
}
