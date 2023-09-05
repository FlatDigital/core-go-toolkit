package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedInternalServerError(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedInternalServerError("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrInternalServerError{}, err.WrappedErr())
}

func Test_ReturnInternalServerErrorError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedInternalServerError("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusInternalServerError, rr.Code)
	ass.Equal("{\"error\":\"InternalServerApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetInternalServerErrorStatusCode(t *testing.T) {
	err := error.NewErrWrappedInternalServerError("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusInternalServerError, statusCode)
}

func Test_InternalServerErrorIsServerError(t *testing.T) {
	err := error.NewErrWrappedInternalServerError("forced for test")

	isClientError := error.IsServerError(err)

	assert.True(t, isClientError)
}

func Test_InternalServerErrorIsClientError(t *testing.T) {
	err := error.NewErrWrappedInternalServerError("forced for test")

	isClientError := error.IsClientError(err)

	assert.False(t, isClientError)
}
