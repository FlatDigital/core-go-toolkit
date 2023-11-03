//nolint:sqlclosecheck
package sqlmock_test

import (
	"context"
	"testing"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
	"github.com/stretchr/testify/assert"
)

func Test_Conn_PatchClose_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()

	// when
	mock.PatchClose(nil)
	err := mock.Close()

	// then
	assert.Nil(err)
}

func Test_Conn_PatchClose_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Close()
	})
}

func Test_Conn_PatchBeginTx_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.BeginTx(ctx, nil)
	})
}

func Test_Conn_PatchExecContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.ExecContext(ctx, "", nil)
	})
}

func Test_Conn_PatchPingContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.PingContext(ctx)
	})
}

func Test_Conn_PatchPrepareContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.PrepareContext(ctx, "")
	})
}

func Test_Conn_PatchQueryContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
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

func Test_Conn_PatchQueryRowContext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()
	ctx := context.Background()

	// when

	// then
	assert.Panics(func() {
		mock.QueryRowContext(ctx, "", nil)
	})
}

func Test_Conn_PatchRaw_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewConnMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Raw(nil)
	})
}
