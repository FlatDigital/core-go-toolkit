package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedVersionNotSupported(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedVersionNotSupported("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrVersionNotSupported{}, err.WrappedErr())
}

func Test_ReturnVersionNotSupportedError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedVersionNotSupported("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusHTTPVersionNotSupported, rr.Code)
	ass.Equal("{\"error\":\"HTTPVersionNotSupportedApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetVersionNotSupportedStatusCode(t *testing.T) {
	err := error.NewErrWrappedVersionNotSupported("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusHTTPVersionNotSupported, statusCode)
}

func Test_VersionNotSupportedIsServerError(t *testing.T) {
	err := error.NewErrWrappedVersionNotSupported("forced for test")

	isClientError := error.IsServerError(err)

	assert.True(t, isClientError)
}

func Test_VersionNotSupportedIsClientError(t *testing.T) {
	err := error.NewErrWrappedVersionNotSupported("forced for test")

	isClientError := error.IsClientError(err)

	assert.False(t, isClientError)
}
