package sqlmock

import (
	"context"
	"database/sql"

	"github.com/FlatDigital/core-go-toolkit/v2/database/converter"
	"github.com/stretchr/testify/mock"
)

// SQLTxMock mock struct
type SQLTxMock struct {
	mock.Mock
	converter.DBTxer
}

// NewTxMockService return mock for database/sql
func NewTxMockService() *SQLTxMock {
	return &SQLTxMock{
		Mock: mock.Mock{},
	}
}

// PatchCommit patches the funcion Commit
func (mock *SQLTxMock) PatchCommit(outputErr error) {
	mock.On("Commit").Return(outputErr).Once()
}

// Commit mocks the real implementation of Commit for the database/sql/tx
func (mock *SQLTxMock) Commit() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchRollback patches the funcion Rollback
func (mock *SQLTxMock) PatchRollback(outputErr error) {
	mock.On("Rollback").Return(outputErr).Once()
}

// Rollback mocks the real implementation of Rollback for the database/sql/tx
func (mock *SQLTxMock) Rollback() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchPrepareContext patches the funcion PrepareContext
func (mock *SQLTxMock) PatchPrepareContext(ctx context.Context, query string,
	stmt converter.DBStmter, outputErr error) {
	mock.On("PrepareContext", ctx, query).Return(stmt, outputErr).Once()
}

// PrepareContext mocks the real implementation of PrepareContext for the database/sql/tx
func (mock *SQLTxMock) PrepareContext(ctx context.Context, query string) (converter.DBStmter, error) {
	args := mock.Called(ctx, query)
	stmt, _ := args.Get(0).(converter.DBStmter)
	err, _ := args.Get(1).(error)
	return stmt, err
}

// Exec mocks the real implementation of Exec for the database/sql/tx
func (mock *SQLTxMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("TODO: Implement mock for sql.tx.Exec")
}

// ExecContext mocks the real implementation of ExecContext for the database/sql/tx
func (mock *SQLTxMock) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	panic("TODO: Implement mock for sql.tx.ExecContext")
}

// Prepare mocks the real implementation of Prepare for the database/sql/tx
func (mock *SQLTxMock) Prepare(query string) (*sql.Stmt, error) {
	panic("TODO: Implement mock for sql.tx.Prepare")
}

// Query mocks the real implementation of Query for the database/sql/tx
func (mock *SQLTxMock) Query(query string, args ...interface{}) (*sql.Rows, error) {
	panic("TODO: Implement mock for sql.tx.Query")
}

// QueryContext mocks the real implementation of QueryContext for the database/sql/tx
func (mock *SQLTxMock) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	panic("TODO: Implement mock for sql.tx.QueryContext")
}

// QueryRow mocks the real implementation of QueryRow for the database/sql/tx
func (mock *SQLTxMock) QueryRow(query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.tx.QueryRow")
}

// QueryRowContext mocks the real implementation of QueryRowContext for the database/sql/tx
func (mock *SQLTxMock) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, error) {
	panic("TODO: Implement mock for sql.tx.QueryRowContext")
}

// Stmt mocks the real implementation of Stmt for the database/sql/tx
func (mock *SQLTxMock) Stmt(stmt *sql.Stmt) *sql.Stmt {
	panic("TODO: Implement mock for sql.tx.Stmt")
}

// StmtContext mocks the real implementation of StmtContext for the database/sql/tx
func (mock *SQLTxMock) StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt {
	panic("TODO: Implement mock for sql.tx.StmtContext")
}
