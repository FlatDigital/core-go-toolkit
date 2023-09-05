package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedNotImplemented(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedNotImplemented("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrNotImplemented{}, err.WrappedErr())
}

func Test_ReturnNotImplementedError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedNotImplemented("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusNotImplemented, rr.Code)
	ass.Equal("{\"error\":\"NotImplementedApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetNotImplementedStatusCode(t *testing.T) {
	err := error.NewErrWrappedNotImplemented("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusNotImplemented, statusCode)
}

func Test_NotImplementedIsServerError(t *testing.T) {
	err := error.NewErrWrappedNotImplemented("forced for test")

	isClientError := error.IsServerError(err)

	assert.True(t, isClientError)
}

func Test_NotImplementedIsClientError(t *testing.T) {
	err := error.NewErrWrappedNotImplemented("forced for test")

	isClientError := error.IsClientError(err)

	assert.False(t, isClientError)
}
