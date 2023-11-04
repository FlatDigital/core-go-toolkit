package sqlmock_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
)

func Test_Stmt_PatchQuery_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	params := make([]interface{}, 0)
	rows := sqlmock.NewRowsMockService()

	// when
	mock.PatchQuery(params, rows, nil)
	outResult, err := mock.Query(params...)

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Stmt_PatchQuery_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Query(nil)
	})
}

func Test_Stmt_PatchClose_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()

	// when
	mock.PatchClose(nil)
	err := mock.Close()

	// then
	assert.Nil(err)
}

func Test_Stmt_PatchClose_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Close()
	})
}

func Test_Stmt_PatchExec_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	params := make([]interface{}, 0)
	result := sqlmock.NewResultMockService()

	// when
	mock.PatchExec(params, result, nil)
	outResult, err := mock.Exec(params...)

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Stmt_PatchExec_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Exec(nil)
	})
}

func Test_Stmt_PatchExecContext_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	params := make([]interface{}, 0)
	result := sqlmock.NewResultMockService()
	ctx := context.Background()

	// when
	mock.PatchExecContext(ctx, params, result, nil)
	outResult, err := mock.ExecContext(ctx, params...)

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Stmt_PatchExecContex_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.ExecContext(ctx, nil)
	})
}

func Test_Stmt_PatchQueryContext_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	params := make([]interface{}, 0)
	rows := sqlmock.NewRowsMockService()
	ctx := context.Background()

	// when
	mock.PatchQueryContext(ctx, params, rows, nil)
	outResult, err := mock.QueryContext(ctx, params...)

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Stmt_PatchQueryContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.QueryContext(ctx, nil)
	})
}

func Test_Stmt_PatchQueryRow_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()

	// when

	// then
	assert.Panics(func() {
		mock.QueryRow(nil)
	})
}

func Test_Stmt_PatchQueryRowContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewStmtMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.QueryRowContext(ctx, "")
	})
}
