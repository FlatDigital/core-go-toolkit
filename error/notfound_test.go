package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedNotFound(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedNotFound("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrNotFound{}, err.WrappedErr())
}

func Test_ReturnNotFoundError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedNotFound("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusNotFound, rr.Code)
	ass.Equal("{\"error\":\"NotFoundApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetNotFoundStatusCode(t *testing.T) {
	err := error.NewErrWrappedNotFound("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusNotFound, statusCode)
}

func Test_NotFoundIsServerError(t *testing.T) {
	err := error.NewErrWrappedNotFound("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_NotFoundIsClientError(t *testing.T) {
	err := error.NewErrWrappedNotFound("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
