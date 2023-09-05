package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedUnauthorized(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedUnauthorized("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrUnauthorized{}, err.WrappedErr())
}

func Test_ReturnUnauthorizedError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedUnauthorized("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusUnauthorized, rr.Code)
	ass.Equal("{\"error\":\"AuthorizationApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetUnauthorizedStatusCode(t *testing.T) {
	err := error.NewErrWrappedUnauthorized("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusUnauthorized, statusCode)
}

func Test_UnauthorizedIsServerError(t *testing.T) {
	err := error.NewErrWrappedUnauthorized("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_UnauthorizedIsClientError(t *testing.T) {
	err := error.NewErrWrappedUnauthorized("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
