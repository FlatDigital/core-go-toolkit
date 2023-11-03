package sqlmock_test

import (
	"testing"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
	"github.com/stretchr/testify/assert"
)

func Test_Result_PatchRowsAffected_Success(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when
	mock.PatchRowsAffected(1, nil)
	outResult, err := mock.RowsAffected()

	// then
	assert.NotNil(outResult)
	assert.Nil(err)
}

func Test_Result_PatchRowsAffected_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when

	// then
	assert.Panics(func() {
		mock.RowsAffected()
	})
}

func Test_Result_PatchLastInsertId_Panic(t *testing.T) {
	// given
	assert := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when

	// then
	assert.Panics(func() {
		mock.LastInsertId()
	})
}
