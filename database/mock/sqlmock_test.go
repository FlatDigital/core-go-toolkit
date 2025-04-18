//nolint:sqlclosecheck
package sqlmock_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
)

func Test_PatchConn_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()
	conn := sqlmock.NewConnMockService()

	// when
	mock.PatchConn(ctx, conn, nil)
	outConn, err := mock.Conn(ctx)

	// then
	assert.NotNil(outConn)
	assert.Nil(err)
}

func Test_PatchConn_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.Conn(ctx)
	})
}

func Test_PatchPing_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when
	mock.PatchPing(nil)
	err := mock.Ping()

	// then
	assert.Nil(err)
}

func Test_PatchPing_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Ping()
	})
}

func Test_PatchPingContext_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when
	mock.PatchPingContext(ctx, nil)
	err := mock.PingContext(ctx)

	// then
	assert.Nil(err)
}

func Test_PatchPingContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.PingContext(ctx)
	})
}

func Test_PatchClose_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when
	mock.PatchClose(nil)
	err := mock.Close()

	// then
	assert.Nil(err)
}

func Test_PatchClose_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Close()
	})
}

func Test_PatchBeginTx_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()
	opts := &sql.TxOptions{}
	tx := sqlmock.NewTxMockService()

	// when
	mock.PatchBeginTx(ctx, opts, tx, nil)
	outTx, err := mock.BeginTx(ctx, opts)

	// then
	assert.NotNil(outTx)
	assert.Nil(err)
}

func Test_PatchBeginTx_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.BeginTx(ctx, nil)
	})
}

func Test_PatchPrepare_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	stmt := sqlmock.NewStmtMockService()

	// when
	mock.PatchPrepare("", stmt, nil)
	outStmt, err := mock.Prepare("")

	// then
	assert.NotNil(outStmt)
	assert.Nil(err)
}

func Test_PatchPrepare_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Prepare("")
	})
}

func Test_PatchStats_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	stats := sql.DBStats{}

	// when
	mock.PatchStats(stats)
	outStats := mock.Stats()

	// then
	assert.NotNil(outStats)
}

func Test_PatchStats_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Stats()
	})
}

func Test_PatchPrepareContext_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()
	stmt := sqlmock.NewStmtMockService()

	// when
	mock.PatchPrepareContext(ctx, "", stmt, nil)
	outStmt, err := mock.PrepareContext(ctx, "")

	// then
	assert.NotNil(outStmt)
	assert.Nil(err)
}

func Test_PatchPrepareContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.PrepareContext(ctx, "")
	})
}

func Test_PatchBegin_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Begin()
	})
}

func Test_PatchDriver_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Driver()
	})
}

func Test_PatchExec_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Exec("", nil)
	})
}

func Test_PatchExecContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.ExecContext(ctx, "", nil)
	})
}

func Test_PatchQuery_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		if rows, err := mock.Query("", nil); err != nil {
			_ = rows.Err()
			return
		}
	})
}

func Test_PatchQueryContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		if rows, err := mock.QueryContext(ctx, "", nil); err != nil {
			_ = rows.Err()
			return
		}
	})
}

func Test_PatchQueryRow_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.QueryRow("", nil)
	})
}

func Test_PatchQueryRowContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.QueryRowContext(ctx, "", nil)
	})
}

func Test_PatchSetConnLifeMaxTime_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.SetConnMaxLifetime(time.Duration(12000))
	})
}

func Test_PatchSetMaxIdleConns_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.SetMaxIdleConns(1)
	})
}

func Test_PatchSetMaxOpenConns_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewMockService()

	// when

	// then
	assert.Panics(func() {
		mock.SetMaxOpenConns(1)
	})
}
