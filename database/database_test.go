package database

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
)

const (
	selectStmt  string = "SELECT * FROM test FOR UPDATE"
	selectStmt2 string = "SELECT * FROM test"
	insertStmt  string = "INSERT INTO test ('created_on', 'email',) VALUES (NOW(), $1) RETURNING id"
)

func Test_DBRow_Equals_True(t *testing.T) {
	ass := assert.New(t)

	row1 := DBRow{
		columns: map[string]DBColumn{
			"column_1": {
				name:  "column_1",
				field: "value_1",
			},
			"column_2": {
				name:  "column_1",
				field: 2,
			},
		},
	}

	row2 := DBRow{
		columns: map[string]DBColumn{
			"column_2": {
				name:  "column_1",
				field: 2,
			},
			"column_1": {
				name:  "column_1",
				field: "value_1",
			},
		},
	}

	areEquals := row1.Equals(&row2)
	ass.True(areEquals)

	areEquals = row2.Equals(&row1)
	ass.True(areEquals)
}

func Test_DBRow_Equals_False(t *testing.T) {
	ass := assert.New(t)

	row1 := DBRow{
		columns: map[string]DBColumn{
			"column_1": {
				name:  "column_1",
				field: "value_1",
			},
			"column_2": {
				name:  "column_1",
				field: 5,
			},
		},
	}

	row2 := DBRow{
		columns: map[string]DBColumn{
			"column_2": {
				name:  "column_1",
				field: 2,
			},
			"column_1": {
				name:  "column_1",
				field: "value_1",
			},
		},
	}

	areEquals := row1.Equals(&row2)
	ass.False(areEquals)
}

func Test_NewService_ReadTimeoutSet_Success(t *testing.T) {
	ass := assert.New(t)

	readTimeout := time.Second

	config := ServiceConfig{
		ConnReadTimeout: &readTimeout,
	}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_NewService_WriteTimeoutSet_Success(t *testing.T) {
	ass := assert.New(t)

	writeTimeout := time.Second

	config := ServiceConfig{
		ConnWriteTimeout: &writeTimeout,
	}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_NewService_TimeoutSet_Success(t *testing.T) {
	ass := assert.New(t)

	timeout := time.Second

	config := ServiceConfig{
		ConnTimeout: &timeout,
	}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_NewService_SetWriteAndTimeout_Success(t *testing.T) {
	ass := assert.New(t)

	timeout := time.Second
	writeTimeout := time.Second

	config := ServiceConfig{
		ConnTimeout:      &timeout,
		ConnWriteTimeout: &writeTimeout,
	}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_NewService_SetDbPort_Success(t *testing.T) {
	ass := assert.New(t)
	dbPort := 5555
	config := ServiceConfig{
		DBPort: dbPort,
	}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_NewService_DefaultPort_Success(t *testing.T) {
	ass := assert.New(t)
	config := ServiceConfig{}
	service, err := NewService(config)
	ass.Nil(err)
	ass.NotNil(service)
}

func Test_Connection_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	dbCtx, err := service.Connection()

	// then
	ass.NotNil(dbCtx)
	ass.Nil(err)
}

func Test_Connection_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()

	// when
	sqlMock.PatchConn(ctx, sqlConn, errors.New("test_conn_err"))
	dbCtx, err := service.Connection()

	// then
	ass.Nil(dbCtx)
	ass.NotNil(err)
}

func Test_Connection_PingContext_Fail(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, errors.New("test_conn_err"))
	dbCtx, err := service.Connection()

	// then
	ass.Nil(dbCtx)
	ass.NotNil(err)
}

func Test_TestConnection_Ctx_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	// when
	sqlMock.PatchPing(nil)
	err := service.TestConnection(nil)

	// then
	ass.Nil(err)
}

func Test_Close_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	sqlConnMock := newDBConnMock()

	dbc := DBContext{
		tx:     nil,
		dbConn: sqlConnMock,
	}

	// when
	sqlConnMock.PatchClose(nil)
	err := service.Close(&dbc)

	// then
	ass.Nil(err)
}

func Test_Close_Dbc_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	// when
	err := service.Close(nil)

	// then
	ass.NotNil(err)
}

func Test_Close_Tx_Not_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()

	dbc := DBContext{
		tx:     txMock,
		dbConn: sqlConnMock,
	}

	// when
	sqlConnMock.PatchClose(nil)
	err := service.Close(&dbc)

	// then
	ass.NotNil(err)
}

func Test_Close_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	sqlConnMock := newDBConnMock()

	dbc := DBContext{
		tx:     nil,
		dbConn: sqlConnMock,
	}

	// when
	sqlConnMock.PatchClose(errors.New("test_close_err"))
	err := service.Close(&dbc)

	// then
	ass.NotNil(err)
}

func Test_Begin_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	outputTx := sqlmock.NewTxMockService()

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, outputTx, nil)
	dbCtx, err := service.Begin(nil)

	// then
	ass.NotNil(dbCtx)
	ass.Nil(err)
}

func Test_Begin_Connection_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	outputTx := sqlmock.NewTxMockService()

	// when
	sqlMock.PatchConn(ctx, sqlConn, errors.New("test_begin_err"))
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, outputTx, nil)
	dbCtx, err := service.Begin(nil)

	// then
	ass.Nil(dbCtx)
	ass.NotNil(err)
}

func Test_Begin_Tx_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	outputTx := sqlmock.NewTxMockService()
	dbc := &DBContext{
		tx:     nil,
		dbConn: nil,
	}

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, outputTx, nil)
	dbCtx, err := service.Begin(dbc)

	// then
	ass.NotNil(dbCtx)
	ass.Nil(err)
}

func Test_Begin_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	outputTx := sqlmock.NewTxMockService()

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, outputTx, errors.New("test_begin_err"))
	dbCtx, err := service.Begin(nil)

	// then
	ass.Nil(dbCtx)
	ass.NotNil(err)
}

func Test_Commit_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchCommit(nil)
	sqlConnMock.PatchClose(nil)
	err := service.Commit(dbc)

	// then
	ass.Nil(err)
}

func Test_Commit_Dbc_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	// when
	err := service.Commit(nil)

	// then
	ass.NotNil(err)
}

func Test_Commit_Nesting_Level_Zero(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 0,
	}

	// when
	err := service.Commit(dbc)

	// then
	ass.NotNil(err)
}

func Test_Commit_Tx_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	dbc := &DBContext{
		tx:           nil,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	err := service.Commit(dbc)

	// then
	ass.NotNil(err)
}

func Test_Commit_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchCommit(errors.New("test_commit_err"))
	sqlConnMock.PatchClose(nil)
	err := service.Commit(dbc)

	// then
	ass.NotNil(err)
}

func Test_Commit_Close_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchCommit(nil)
	sqlConnMock.PatchClose(errors.New("test_commit_err"))
	err := service.Commit(dbc)

	// then
	ass.NotNil(err)
}

func Test_Commit_Nesting_Level_Two(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 2,
	}

	// when
	txMock.PatchCommit(nil)
	sqlConnMock.PatchClose(nil)
	err := service.Commit(dbc)

	// then
	ass.Nil(err)
}

func Test_Rollback_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchRollback(nil)
	sqlConnMock.PatchClose(nil)
	err := service.Rollback(dbc)

	// then
	ass.Nil(err)
}

func Test_Rollback_Dbc_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	// when
	err := service.Rollback(nil)

	// then
	ass.NotNil(err)
}

func Test_Rollback_Nesting_Level_Zero(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 0,
	}

	// when
	err := service.Rollback(dbc)

	// then
	ass.Nil(err)
}

func Test_Rollback_Tx_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	dbc := &DBContext{
		tx:           nil,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	err := service.Rollback(dbc)

	// then
	ass.NotNil(err)
}

func Test_Rollback_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchRollback(errors.New("test_rollback_err"))
	sqlConnMock.PatchClose(nil)
	err := service.Rollback(dbc)

	// then
	ass.NotNil(err)
}

func Test_Rollback_Close_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)

	sqlConnMock := newDBConnMock()
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:           txMock,
		dbConn:       sqlConnMock,
		nestingLevel: 1,
	}

	// when
	txMock.PatchRollback(nil)
	sqlConnMock.PatchClose(errors.New("test_rollback_err"))
	err := service.Rollback(dbc)

	// then
	ass.NotNil(err)
}

func Test_WithTransaction_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	txMock := sqlmock.NewTxMockService()

	txFn := func(dbc *DBContext) error { return nil }

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, txMock, nil)
	txMock.PatchCommit(nil)
	sqlConn.PatchClose(nil)
	err := service.WithTransaction(txFn)

	// then
	ass.Nil(err)
}

func Test_WithTransaction_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	txMock := sqlmock.NewTxMockService()

	testErr := errors.New("test error")
	txFn := func(dbc *DBContext) error { return testErr }

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, txMock, nil)
	txMock.PatchRollback(nil)
	sqlConn.PatchClose(nil)
	err := service.WithTransaction(txFn)

	// then
	ass.NotNil(err)
	ass.EqualValues(testErr, err)
}

func Test_WithTransaction_Commit_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	txMock := sqlmock.NewTxMockService()

	txFn := func(dbc *DBContext) error { return nil }

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, txMock, nil)
	testErr := errors.New("test error")
	txMock.PatchCommit(testErr)
	sqlConn.PatchClose(nil)
	err := service.WithTransaction(txFn)

	// then
	ass.NotNil(err)
	ass.EqualValues(testErr, err)
}

func Test_WithTransaction_Rollback_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()
	txMock := sqlmock.NewTxMockService()

	firstErr := errors.New("txFn error")
	txFn := func(dbc *DBContext) error { return firstErr }

	// when
	sqlMock.PatchConn(ctx, sqlConn, nil)
	sqlMock.PatchPing(nil)
	sqlMock.PatchPingContext(ctx, nil)
	sqlMock.PatchBeginTx(ctx, nil, txMock, nil)
	secondErr := errors.New("rollback error")
	txMock.PatchRollback(secondErr)
	sqlConn.PatchClose(nil)
	err := service.WithTransaction(txFn)

	// then
	expectedErr := errors.New("error rollbacking transaction: rollback error, txFn error")
	ass.NotNil(err)
	ass.EqualValues(expectedErr, err)
}

func Test_Select_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(queryForUpdate, stmtMock, nil)
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Select_Success_With_Tx(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:  txMock,
		ctx: ctx,
	}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQueryContext(ctx, params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	txMock.PatchPrepareContext(ctx, queryForUpdate, stmtMock, nil)
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Select_Success_With_Conn(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQueryContext(ctx, params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, queryForUpdate, stmtMock, nil)
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Select_Success_With_Conn_PrepareContext_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQueryContext(ctx, params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, queryForUpdate, stmtMock, errors.New("test_query_err"))
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Select_Success_With_Conn_QueryContext_Err(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQueryContext(ctx, params, rowsMock, errors.New("test_query_err"))
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, queryForUpdate, stmtMock, nil)
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Select_Prepare_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	params := make([]interface{}, 0)
	params = append(params, 3)

	// when
	sqlMock.PatchPrepare(queryForUpdate, nil, errors.New("test_select_err"))
	dbResult, err := service.Select(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_SelectUniqueValue_Prepare_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	params := make([]interface{}, 0)
	params = append(params, 3)

	// when
	sqlMock.PatchPrepare(queryForUpdate, nil, errors.New("test_select_err"))
	dbResult, err := service.SelectUniqueValue(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_SelectUniqueValue_No_Result(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(queryForUpdate, stmtMock, nil)
	dbResult, err := service.SelectUniqueValue(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.Nil(err)
}

func Test_SelectUniqueValue_Many_Result(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(queryForUpdate, stmtMock, nil)
	dbResult, err := service.SelectUniqueValue(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_SelectUniqueValueNonEmpty_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(queryForUpdate, stmtMock, nil)
	dbResult, err := service.SelectUniqueValueNonEmpty(dbc, query, true, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_SelectUniqueValueNonEmptyWithInsertStmt_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	query := insertStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 1)
	columns := []string{"id"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(errors.New("test error"))
	// One more for defer()
	rowsMock.PatchClose(errors.New("test error"))
	rowsMock.PatchNext(true)
	rowsMock.PatchScan(columnPointers, nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(query, stmtMock, nil)
	dbResult, err := service.SelectUniqueValueNonEmpty(nil, query, false, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_SelectUniqueValueNonEmpty_Prepare_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	params := make([]interface{}, 0)
	params = append(params, 3)

	// when
	sqlMock.PatchPrepare(queryForUpdate, nil, errors.New("test_select_err"))
	dbResult, err := service.SelectUniqueValueNonEmpty(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_SelectUniqueValueNonEmpty_No_Result(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	queryForUpdate := selectStmt
	stmtMock := newDBStmtMock()
	rowsMock := newDBRowsMock()
	params := make([]interface{}, 0)
	params = append(params, 3)
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}

	// when
	rowsMock.PatchColumns(columns, nil)
	rowsMock.PatchClose(nil)
	rowsMock.PatchNext(false)
	stmtMock.PatchQuery(params, rowsMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(queryForUpdate, stmtMock, nil)
	dbResult, err := service.SelectUniqueValueNonEmpty(dbc, query, true, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExec(params, resultMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Execute_Prepare_Err(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExec(params, resultMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(query, stmtMock, errors.New("test_execute_err"))
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_Stmt_Execute_Err(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExec(params, resultMock, errors.New("test_execute_err"))
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_With_Tx(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:  txMock,
		ctx: ctx,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, nil)
	stmtMock.PatchClose(nil)
	txMock.PatchPrepareContext(ctx, query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Execute_With_Tx_Prepare_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:  txMock,
		ctx: ctx,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, nil)
	stmtMock.PatchClose(nil)
	txMock.PatchPrepareContext(ctx, query, stmtMock, errors.New("test_execute_err"))
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_With_Tx_ExecContext_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, _ := newMockService(config)
	txMock := newDBTxMock()
	dbc := &DBContext{
		tx:  txMock,
		ctx: ctx,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, errors.New("test_execute_err"))
	stmtMock.PatchClose(nil)
	txMock.PatchPrepareContext(ctx, query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_With_Conn(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.NotNil(dbResult)
	ass.Nil(err)
}

func Test_Execute_With_Conn_PrepareContext_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, query, stmtMock, errors.New("test_execute_err"))
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_With_Conn_Stmt_ExecuteContext_Err(t *testing.T) {
	// given
	ass := assert.New(t)

	ctx := context.Background()
	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	connMock := newDBConnMock()
	dbc := &DBContext{
		tx:     nil,
		ctx:    ctx,
		dbConn: connMock,
	}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, nil)
	stmtMock.PatchExecContext(ctx, params, resultMock, errors.New("test_execute_err"))
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepareContext(ctx, query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_Execute_RowsAffected_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)
	dbc := &DBContext{}
	query := selectStmt2
	stmtMock := newDBStmtMock()
	params := make([]interface{}, 0)
	params = append(params, 3)

	resultMock := newDBResultMock()

	// when
	resultMock.PatchRowsAffected(3, errors.New("test_execute_err"))
	stmtMock.PatchExec(params, resultMock, nil)
	stmtMock.PatchClose(nil)
	sqlMock.PatchPrepare(query, stmtMock, nil)
	dbResult, err := service.Execute(dbc, query, params...)

	// then
	ass.Nil(dbResult)
	ass.NotNil(err)
}

func Test_MySQLError(t *testing.T) {
	// given
	ass := assert.New(t)

	config := ServiceConfig{
		MaxConnectionRetries: 1,
	}
	service, sqlMock := newMockService(config)

	ctx := context.Background()
	sqlConn := newDBConnMock()

	// when
	sqlMock.PatchConn(ctx, sqlConn, &pq.Error{
		Code:    pq.ErrorCode("1234"),
		Message: "postgres-error",
	})
	dbCtx, err := service.Connection()

	// then
	ass.Nil(dbCtx)
	ass.NotNil(err)
}

func TestService_SelectOnDbLinkView_Success(t *testing.T) {
	service, sqlMock := newMockService(ServiceConfig{
		MaxConnectionRetries: 1,
	})
	dblinkConnMock, _ := NewDbLinkConnection("test", "127.0.0.1", uint(1234), "usrtest", "pass123", "db_test")
	//dblink open connection
	stmtMock := newDBStmtMock()
	resultDbLinkOpenConnMock := newDBResultMock()
	resultDbLinkOpenConnMock.PatchRowsAffected(1, nil)
	stmtMock.PatchExec(nil, resultDbLinkOpenConnMock, nil)
	stmtMock.PatchClose(nil)
	dbLinkOpenConn := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1234 dbname=db_test user=usrtest password=pass123')"
	sqlMock.PatchPrepare(dbLinkOpenConn, stmtMock, nil)
	//dblink query
	stDbLinkMock := newDBStmtMock()
	rowsDblinkMock := newDBRowsMock()
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}
	rowsDblinkMock.PatchColumns(columns, nil)
	rowsDblinkMock.PatchClose(nil)
	rowsDblinkMock.PatchNext(true)
	rowsDblinkMock.PatchScan(columnPointers, nil)
	rowsDblinkMock.PatchNext(false)
	stDbLinkMock.PatchQuery(nil, rowsDblinkMock, nil)
	stDbLinkMock.PatchClose(nil)
	queryDbLink := "SELECT * FROM test"
	sqlMock.PatchPrepare(queryDbLink, stDbLinkMock, nil)

	//dblink close connection
	stmtCloseConnMock := newDBStmtMock()
	resultDbLinkOpenConnMock2 := newDBResultMock()
	resultDbLinkOpenConnMock2.PatchRowsAffected(1, nil)
	stmtCloseConnMock.PatchExec(nil, resultDbLinkOpenConnMock2, nil)
	stmtCloseConnMock.PatchClose(nil)
	dbLinkCloseConn := "SELECT dblink_disconnect('test')"
	sqlMock.PatchPrepare(dbLinkCloseConn, stmtCloseConnMock, nil)

	queryDbLinkMock := "SELECT * FROM test"
	dbResult, err := service.SelectOnDbLinkView(dblinkConnMock, nil, queryDbLinkMock)
	assert.NotNil(t, dbResult)
	assert.NoError(t, err)
}

func TestService_SelectOnDbLinkView_SuccessWithParams(t *testing.T) {
	service, sqlMock := newMockService(ServiceConfig{
		MaxConnectionRetries: 1,
	})
	dblinkConnMock, _ := NewDbLinkConnection("test", "127.0.0.1", uint(1234), "usrtest", "pass123", "db_test")
	//dblink open connection
	stmtMock := newDBStmtMock()
	resultDbLinkOpenConnMock := newDBResultMock()
	resultDbLinkOpenConnMock.PatchRowsAffected(1, nil)
	stmtMock.PatchExec(nil, resultDbLinkOpenConnMock, nil)
	stmtMock.PatchClose(nil)
	dbLinkOpenConn := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1234 dbname=db_test user=usrtest password=pass123')"
	sqlMock.PatchPrepare(dbLinkOpenConn, stmtMock, nil)
	//dblink query
	stDbLinkMock := newDBStmtMock()
	rowsDblinkMock := newDBRowsMock()
	columns := []string{"columnA", "columnB", "columnC"}
	columnsAux := make([]interface{}, len(columns))
	columnPointers := make([]interface{}, len(columns))
	for i := range columnsAux {
		columnPointers[i] = &columnsAux[i]
	}
	rowsDblinkMock.PatchColumns(columns, nil)
	rowsDblinkMock.PatchClose(nil)
	rowsDblinkMock.PatchNext(true)
	rowsDblinkMock.PatchScan(columnPointers, nil)
	rowsDblinkMock.PatchNext(false)
	queryParams := []interface{}{1}
	stDbLinkMock.PatchQuery(queryParams, rowsDblinkMock, nil)
	stDbLinkMock.PatchClose(nil)
	queryDbLink := "SELECT * FROM test where id=?"
	sqlMock.PatchPrepare(queryDbLink, stDbLinkMock, nil)

	//dblink close connection
	stmtCloseConnMock := newDBStmtMock()
	resultDbLinkOpenConnMock2 := newDBResultMock()
	resultDbLinkOpenConnMock2.PatchRowsAffected(1, nil)
	stmtCloseConnMock.PatchExec(nil, resultDbLinkOpenConnMock2, nil)
	stmtCloseConnMock.PatchClose(nil)
	dbLinkCloseConn := "SELECT dblink_disconnect('test')"
	sqlMock.PatchPrepare(dbLinkCloseConn, stmtCloseConnMock, nil)

	queryDbLinkMock := "SELECT * FROM test where id=?"
	dbResult, err := service.SelectOnDbLinkView(dblinkConnMock, nil, queryDbLinkMock, queryParams...)
	assert.NotNil(t, dbResult)
	assert.NoError(t, err)
}

func TestService_SelectOnDbLinkView_DblinkOpenConnError(t *testing.T) {
	service, sqlMock := newMockService(ServiceConfig{
		MaxConnectionRetries: 1,
	})
	dblinkConnMock, _ := NewDbLinkConnection("test", "127.0.0.1", uint(1234), "usrtest", "pass123", "db_test")
	//dblink open connection
	stmtMock := newDBStmtMock()
	resultDbLinkOpenConnMock := newDBResultMock()
	resultDbLinkOpenConnMock.PatchRowsAffected(1, nil)
	stmtMock.PatchExec(nil, resultDbLinkOpenConnMock, nil)
	stmtMock.PatchClose(nil)
	dbLinkOpenConn := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1234 dbname=db_test user=usrtest password=pass123')"
	sqlMock.PatchPrepare(dbLinkOpenConn, stmtMock, fmt.Errorf("error on dblink connect"))

	//dblink close connection
	stmtCloseConnMock := newDBStmtMock()
	resultDbLinkOpenConnMock2 := newDBResultMock()
	resultDbLinkOpenConnMock2.PatchRowsAffected(1, nil)
	stmtCloseConnMock.PatchExec(nil, resultDbLinkOpenConnMock2, nil)
	stmtCloseConnMock.PatchClose(nil)
	dbLinkCloseConn := "SELECT dblink_disconnect('test')"
	sqlMock.PatchPrepare(dbLinkCloseConn, stmtCloseConnMock, nil)

	queryDbLinkMock := "SELECT * FROM test"
	dbResult, err := service.SelectOnDbLinkView(dblinkConnMock, nil, queryDbLinkMock)
	assert.Nil(t, dbResult)
	assert.Error(t, err)
}

func TestService_SelectOnDbLinkView_ResultError(t *testing.T) {
	service, sqlMock := newMockService(ServiceConfig{
		MaxConnectionRetries: 1,
	})
	dblinkConnMock, _ := NewDbLinkConnection("test", "127.0.0.1", uint(1234), "usrtest", "pass123", "db_test")
	//dblink open connection
	stmtMock := newDBStmtMock()
	resultDbLinkOpenConnMock := newDBResultMock()
	resultDbLinkOpenConnMock.PatchRowsAffected(1, nil)
	stmtMock.PatchExec(nil, resultDbLinkOpenConnMock, nil)
	stmtMock.PatchClose(nil)
	dbLinkOpenConn := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1234 dbname=db_test user=usrtest password=pass123')"
	sqlMock.PatchPrepare(dbLinkOpenConn, stmtMock, nil)
	//dblink query
	stDbLinkMock := newDBStmtMock()
	queryDbLink := "SELECT * FROM test"
	sqlMock.PatchPrepare(queryDbLink, stDbLinkMock, fmt.Errorf("error on query"))

	//dblink close connection
	stmtCloseConnMock := newDBStmtMock()
	resultDbLinkOpenConnMock2 := newDBResultMock()
	resultDbLinkOpenConnMock2.PatchRowsAffected(1, nil)
	stmtCloseConnMock.PatchExec(nil, resultDbLinkOpenConnMock2, nil)
	stmtCloseConnMock.PatchClose(nil)
	dbLinkCloseConn := "SELECT dblink_disconnect('test')"
	sqlMock.PatchPrepare(dbLinkCloseConn, stmtCloseConnMock, nil)

	queryDbLinkMock := "SELECT * FROM test"
	dbResult, err := service.SelectOnDbLinkView(dblinkConnMock, nil, queryDbLinkMock)
	assert.Nil(t, dbResult)
	assert.Error(t, err)
}

func TestService_SelectOnDbLinkView_ErrorWithParams(t *testing.T) {
	service, sqlMock := newMockService(ServiceConfig{
		MaxConnectionRetries: 1,
	})
	dblinkConnMock, _ := NewDbLinkConnection("test", "127.0.0.1", uint(1234), "usrtest", "pass123", "db_test")
	//dblink open connection
	stmtMock := newDBStmtMock()
	resultDbLinkOpenConnMock := newDBResultMock()
	resultDbLinkOpenConnMock.PatchRowsAffected(1, nil)
	stmtMock.PatchExec(nil, resultDbLinkOpenConnMock, nil)
	stmtMock.PatchClose(nil)
	dbLinkOpenConn := "SELECT * FROM dblink_connect('test', 'host=127.0.0.1 port=1234 dbname=db_test user=usrtest password=pass123')"
	sqlMock.PatchPrepare(dbLinkOpenConn, stmtMock, nil)
	//dblink query
	stDbLinkMock := newDBStmtMock()
	queryParams := []interface{}{1}
	queryDbLink := "SELECT * FROM test where id=?"
	sqlMock.PatchPrepare(queryDbLink, stDbLinkMock, fmt.Errorf("error on database"))

	//dblink close connection
	stmtCloseConnMock := newDBStmtMock()
	resultDbLinkOpenConnMock2 := newDBResultMock()
	resultDbLinkOpenConnMock2.PatchRowsAffected(1, nil)
	stmtCloseConnMock.PatchExec(nil, resultDbLinkOpenConnMock2, nil)
	stmtCloseConnMock.PatchClose(nil)
	dbLinkCloseConn := "SELECT dblink_disconnect('test')"
	sqlMock.PatchPrepare(dbLinkCloseConn, stmtCloseConnMock, nil)

	queryDbLinkMock := "SELECT * FROM test where id=?"
	dbResult, err := service.SelectOnDbLinkView(dblinkConnMock, nil, queryDbLinkMock, queryParams...)
	assert.Nil(t, dbResult)
	assert.Error(t, err)
}

func newMockService(config ServiceConfig) (service, *sqlmock.SQLMock) {
	sqlMock := sqlmock.NewMockService()

	return service{
		db:                   sqlMock,
		maxConnectionRetries: config.MaxConnectionRetries,
		datadogMetricPrefix:  config.DatadogMetricPrefix,
	}, sqlMock
}

func newDBConnMock() *sqlmock.SQLConnMock {
	return sqlmock.NewConnMockService()
}

func newDBTxMock() *sqlmock.SQLTxMock {
	return sqlmock.NewTxMockService()
}

func newDBStmtMock() *sqlmock.SQLStmtMock {
	return sqlmock.NewStmtMockService()
}

func newDBRowsMock() *sqlmock.SQLRowsMock {
	return sqlmock.NewRowsMockService()
}

func newDBResultMock() *sqlmock.SQLResultMock {
	return sqlmock.NewResultMockService()
}
