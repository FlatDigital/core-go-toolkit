package sqlmock_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
)

func Test_Rows_PatchColumns_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when
	mock.PatchColumns([]string{"a"}, nil)
	outResult, err := mock.Columns()

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Rows_PatchColumns_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Columns()
	})
}

func Test_Rows_PatchClose_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when
	mock.PatchClose(nil)
	err := mock.Close()

	// then
	assert.Nil(err)
}

func Test_Rows_PatchClose_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Close()
	})
}

func Test_Rows_PatchNext_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when
	mock.PatchNext(false)
	outResult := mock.Next()

	// then
	assert.NotNil(outResult)
}

func Test_Rows_PatchNext_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Next()
	})
}

func Test_Rows_PatchScan_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()
	dest := make([]interface{}, 0)

	// when
	mock.PatchScan(dest, nil)
	err := mock.Scan(dest...)

	// then
	assert.Nil(err)
}

func Test_Rows_PatchScan_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Scan(nil)
	})
}

func Test_Rows_PatchColumnTypes_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.ColumnTypes()
	})
}

func Test_Rows_PatchErr_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.Err()
	})
}

func Test_Rows_PatchNextResultSet_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewRowsMockService()

	// when

	// then
	assert.Panics(func() {
		mock.NextResultSet()
	})
}
