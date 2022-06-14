package sqlmock

import (
	"context"
	"database/sql"

	"github.com/FlatDigital/core-go-toolkit/database/converter"
)

// SQLTxMock mock struct
type SQLTxMock struct {
	patchTxCommit         []outputForTxCommit
	patchTxRollback       []outputForTxRollback
	patchTxPrepareContext map[hash][]outputForTxPrepareContext
	converter.DBTxer
}

// NewTxMockService return mock for database/sql
func NewTxMockService() *SQLTxMock {
	sqlTxMock := SQLTxMock{
		patchTxCommit:         make([]outputForTxCommit, 0),
		patchTxRollback:       make([]outputForTxRollback, 0),
		patchTxPrepareContext: map[hash][]outputForTxPrepareContext{},
	}

	return &sqlTxMock
}

type (
	outputForTxCommit struct {
		outputError error
	}

	outputForTxRollback struct {
		outputError error
	}

	inputForTxPrepareContext struct {
		Context context.Context
		Query   string
	}

	outputForTxPrepareContext struct {
		stmt        converter.DBStmter
		outputError error
	}
)

// PatchCommit patches the funcion Commit
func (mock *SQLTxMock) PatchCommit(outputErr error) {
	output := outputForTxCommit{
		outputError: outputErr,
	}

	mock.patchTxCommit = append(mock.patchTxCommit, output)
}

// Commit mocks the real implementation of Commit for the database/sql/tx
func (mock *SQLTxMock) Commit() error {
	if len(mock.patchTxCommit) == 0 {
		panic("Mock not available for SQLTxMock.Commit")
	}

	output := mock.patchTxCommit[0]
	// dequeue
	mock.patchTxCommit = mock.patchTxCommit[1:]

	return output.outputError
}

// PatchRollback patches the funcion Rollback
func (mock *SQLTxMock) PatchRollback(outputErr error) {
	output := outputForTxRollback{
		outputError: outputErr,
	}

	mock.patchTxRollback = append(mock.patchTxRollback, output)
}

// Rollback mocks the real implementation of Rollback for the database/sql/tx
func (mock *SQLTxMock) Rollback() error {
	if len(mock.patchTxRollback) == 0 {
		panic("Mock not available for SQLTxMock.Rollback")
	}

	output := mock.patchTxRollback[0]
	// dequeue
	mock.patchTxRollback = mock.patchTxRollback[1:]

	return output.outputError
}

// PatchPrepareContext patches the funcion PrepareContext
func (mock *SQLTxMock) PatchPrepareContext(ctx context.Context, query string,
	stmt converter.DBStmter, outputErr error) {
	input := inputForTxPrepareContext{
		Context: ctx,
		Query:   query,
	}
	hash := toHash(input)

	output := outputForTxPrepareContext{
		stmt:        stmt,
		outputError: outputErr,
	}

	mock.patchTxPrepareContext[hash] = append(mock.patchTxPrepareContext[hash], output)
}

// PrepareContext mocks the real implementation of PrepareContext for the database/sql/tx
func (mock *SQLTxMock) PrepareContext(ctx context.Context, query string) (converter.DBStmter, error) {
	inputStruct := inputForTxPrepareContext{
		Context: ctx,
		Query:   query,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchTxPrepareContext[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLTxMock.PrepareContext")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchTxPrepareContext[hash] = mocksArr[1:]

	return output.stmt, output.outputError
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
