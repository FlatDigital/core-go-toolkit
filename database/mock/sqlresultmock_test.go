package sqlmock_test

import (
	"testing"

	sqlmock "github.com/FlatDigital/core-go-toolkit/v2/database/mock"
	"github.com/stretchr/testify/assert"
)

func Test_Result_PatchRowsAffected_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when
	mock.PatchRowsAffected(1, nil)
	outResult, err := mock.RowsAffected()

	// then
	ass.NotNil(outResult)
	ass.Nil(err)
}

func Test_Result_PatchRowsAffected_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when

	// then
	ass.Panics(func() {
		mock.RowsAffected()
	})
}

func Test_Result_PatchLastInsertId_Panic(t *testing.T) {
	// given
	ass := assert.New(t)

	mock := sqlmock.NewResultMockService()

	// when

	// then
	ass.Panics(func() {
		mock.LastInsertId()
	})
}
