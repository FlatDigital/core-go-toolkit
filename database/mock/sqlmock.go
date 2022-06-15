package sqlmock

import (
	"context"
	"crypto/md5"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/FlatDigital/core-go-toolkit/database/converter"
)

// SQLMock mock struct
type SQLMock struct {
	patchConn           map[hash][]outputForConn
	patchPing           []outputForPing
	patchPingContext    map[hash][]outputForPingContext
	patchClose          []outputForClose
	patchBeginTx        map[hash][]outputForBeginTx
	patchPrepare        map[hash][]outputForPrepare
	patchStats          []outputForStats
	patchPrepareContext map[hash][]outputForPrepareContext
	converter.DBer
}

// NewMockService return mock for database/sql
func NewMockService() *SQLMock {
	sqlMock := SQLMock{
		patchConn:           map[hash][]outputForConn{},
		patchPing:           make([]outputForPing, 0),
		patchPingContext:    map[hash][]outputForPingContext{},
		patchClose:          make([]outputForClose, 0),
		patchBeginTx:        map[hash][]outputForBeginTx{},
		patchPrepare:        map[hash][]outputForPrepare{},
		patchStats:          make([]outputForStats, 0),
		patchPrepareContext: map[hash][]outputForPrepareContext{},
	}

	return &sqlMock
}

type (
	inputForConn struct {
		Context context.Context
	}

	outputForConn struct {
		conn        converter.DBConner
		outputError error
	}

	outputForPing struct {
		outputError error
	}

	inputForPingContext struct {
		Context context.Context
	}

	outputForPingContext struct {
		outputError error
	}

	outputForClose struct {
		outputError error
	}

	inputForBeginTx struct {
		Context context.Context
		Opts    *sql.TxOptions
	}

	outputForBeginTx struct {
		tx          converter.DBTxer
		outputError error
	}

	inputForPrepare struct {
		Query string
	}

	outputForPrepare struct {
		stmt        converter.DBStmter
		outputError error
	}

	outputForStats struct {
		outputStats sql.DBStats
	}

	inputForPrepareContext struct {
		Context context.Context
		Query   string
	}

	outputForPrepareContext struct {
		stmt        converter.DBStmter
		outputError error
	}

	hash [16]byte
)

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// PatchConn patches the funcion Conn
func (mock *SQLMock) PatchConn(ctx context.Context, outputConn converter.DBConner, outputErr error) {
	input := inputForConn{
		Context: ctx,
	}
	hash := toHash(input)

	output := outputForConn{
		conn:        outputConn,
		outputError: outputErr,
	}

	mock.patchConn[hash] = append(mock.patchConn[hash], output)
}

// Conn mocks the real implementation of Conn for the database/sql
func (mock *SQLMock) Conn(ctx context.Context) (converter.DBConner, error) {
	inputStruct := inputForConn{
		Context: ctx,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchConn[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLMock.Conn")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchConn[hash] = mocksArr[1:]

	return output.conn, output.outputError
}

// PatchPing patches the funcion Ping
func (mock *SQLMock) PatchPing(outputErr error) {
	output := outputForPing{
		outputError: outputErr,
	}

	mock.patchPing = append(mock.patchPing, output)
}

// Ping mocks the real implementation of Ping for the database/sql
func (mock *SQLMock) Ping() error {
	if len(mock.patchPing) == 0 {
		panic("Mock not available for SQLMock.Ping")
	}

	output := mock.patchPing[0]
	// dequeue
	mock.patchPing = mock.patchPing[1:]

	return output.outputError
}

// PatchPingContext patches the funcion PingContext
func (mock *SQLMock) PatchPingContext(ctx context.Context, outputErr error) {
	input := inputForPingContext{
		Context: ctx,
	}
	hash := toHash(input)

	output := outputForPingContext{
		outputError: outputErr,
	}

	mock.patchPingContext[hash] = append(mock.patchPingContext[hash], output)
}

// PingContext mocks the real implementation of PingContext for the database/sql
func (mock *SQLMock) PingContext(ctx context.Context) error {
	inputStruct := inputForPingContext{
		Context: ctx,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchPingContext[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLMock.PingContext")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchPingContext[hash] = mocksArr[1:]

	return output.outputError
}

// PatchClose patches the funcion Close
func (mock *SQLMock) PatchClose(outputErr error) {
	output := outputForClose{
		outputError: outputErr,
	}

	mock.patchClose = append(mock.patchClose, output)
}

// Close mocks the real implementation of Close for the database/sql
func (mock *SQLMock) Close() error {
	if len(mock.patchClose) == 0 {
		panic("Mock not available for SQLMock.Close")
	}

	output := mock.patchClose[0]
	// dequeue
	mock.patchClose = mock.patchClose[1:]

	return output.outputError
}

// PatchBeginTx patches the funcion BeginTx
func (mock *SQLMock) PatchBeginTx(ctx context.Context, opts *sql.TxOptions, tx converter.DBTxer, outputErr error) {
	input := inputForBeginTx{
		Context: ctx,
		Opts:    opts,
	}
	hash := toHash(input)

	output := outputForBeginTx{
		tx:          tx,
		outputError: outputErr,
	}

	mock.patchBeginTx[hash] = append(mock.patchBeginTx[hash], output)
}

// BeginTx mocks the real implementation of BeginTx for the database/sql
func (mock *SQLMock) BeginTx(ctx context.Context, opts *sql.TxOptions) (converter.DBTxer, error) {
	inputStruct := inputForBeginTx{
		Context: ctx,
		Opts:    opts,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchBeginTx[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLMock.BeginTx")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchBeginTx[hash] = mocksArr[1:]

	return output.tx, output.outputError
}

// PatchPrepare patches the funcion Prepare
func (mock *SQLMock) PatchPrepare(query string, stmt converter.DBStmter, outputErr error) {
	input := inputForPrepare{
		Query: query,
	}
	hash := toHash(input)

	output := outputForPrepare{
		stmt:        stmt,
		outputError: outputErr,
	}

	mock.patchPrepare[hash] = append(mock.patchPrepare[hash], output)
}

// Prepare mocks the real implementation of Prepare for the database/sql
func (mock *SQLMock) Prepare(query string) (converter.DBStmter, error) {
	inputStruct := inputForPrepare{
		Query: query,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchPrepare[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLMock.Prepare")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchPrepare[hash] = mocksArr[1:]

	return output.stmt, output.outputError
}

// PatchStats patches the funcion Stats
func (mock *SQLMock) PatchStats(stats sql.DBStats) {
	output := outputForStats{
		outputStats: stats,
	}

	mock.patchStats = append(mock.patchStats, output)
}

// Stats mocks the real implementation of Stats for the database/sql
func (mock *SQLMock) Stats() sql.DBStats {
	if len(mock.patchStats) == 0 {
		panic("Mock not available for SQLMock.Stats")
	}

	output := mock.patchStats[0]
	// dequeue
	mock.patchStats = mock.patchStats[1:]

	return output.outputStats
}

// PatchPrepareContext patches the funcion PrepareContext
func (mock *SQLMock) PatchPrepareContext(ctx context.Context, query string, stmt converter.DBStmter, outputErr error) {
	input := inputForPrepareContext{
		Context: ctx,
		Query:   query,
	}
	hash := toHash(input)

	output := outputForPrepareContext{
		stmt:        stmt,
		outputError: outputErr,
	}

	mock.patchPrepareContext[hash] = append(mock.patchPrepareContext[hash], output)
}

// PrepareContext mocks the real implementation of PrepareContext for the database/sql
func (mock *SQLMock) PrepareContext(ctx context.Context, query string) (converter.DBStmter, error) {
	inputStruct := inputForPrepareContext{
		Context: ctx,
		Query:   query,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchPrepareContext[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLMock.PrepareContext")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchPrepareContext[hash] = mocksArr[1:]

	return output.stmt, output.outputError
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
