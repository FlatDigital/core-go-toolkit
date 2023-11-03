//nolint:sqlclosecheck
package sqlmock_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
)

func Test_Tx_PatchCommit_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when
	mock.PatchCommit(nil)
	err := mock.Commit()

	// then
	ass.Nil(err)
}

func Test_Tx_PatchCommit_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.Commit()
	})
}

func Test_Tx_PatchRollback_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when
	mock.PatchRollback(nil)
	err := mock.Rollback()

	// then
	ass.Nil(err)
}

func Test_Tx_PatchRollback_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.Rollback()
	})
}

func Test_Tx_PatchPrepareContext_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	stmt := sqlmock.NewStmtMockService()
	ctx := context.Background()

	// when
	mock.PatchPrepareContext(ctx, "", stmt, nil)
	out, err := mock.PrepareContext(ctx, "")

	// then
	ass.NotNil(out)
	ass.Nil(err)
}

func Test_Tx_PatchPrepareContext_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	ctx := context.Background()

	// when

	// then
	ass.Panics(func() {
		mock.PrepareContext(ctx, "")
	})
}

func Test_Tx_PatchExec_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.Exec("", nil)
	})
}

func Test_Tx_PatchExecContext_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	ctx := context.Background()

	// when

	// then
	ass.Panics(func() {
		mock.ExecContext(ctx, "", nil)
	})
}

func Test_Tx_PatchPrepare_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.Prepare("")
	})
}

func Test_Tx_PatchQuery_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		if rows, err := mock.Query("", nil); err != nil {
			_ = rows.Err()
			return
		}
	})
}

func Test_Tx_PatchQueryContext_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	ctx := context.Background()

	// when

	// then
	ass.Panics(func() {
		if rows, err := mock.QueryContext(ctx, "", nil); err != nil {
			_ = rows.Err()
			return
		}
	})
}

func Test_Tx_PatchQueryRow_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.QueryRow("", nil)
	})
}

func Test_Tx_PatchQueryRowContext_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	ctx := context.Background()

	// when

	// then
	ass.Panics(func() {
		mock.QueryRowContext(ctx, "", nil)
	})
}

func Test_Tx_PatchStmt_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()

	// when

	// then
	ass.Panics(func() {
		mock.Stmt(nil)
	})
}

func Test_Tx_PatchStmtContext_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewTxMockService()
	ctx := context.Background()

	// when

	// then
	ass.Panics(func() {
		mock.StmtContext(ctx, nil)
	})
}
