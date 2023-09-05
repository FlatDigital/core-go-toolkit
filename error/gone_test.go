package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedGone(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedGone("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrGone{}, err.WrappedErr())
}

func Test_ReturnGoneError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedGone("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusGone, rr.Code)
	ass.Equal("{\"error\":\"GoneApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetGoneStatusCode(t *testing.T) {
	err := error.NewErrWrappedGone("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusGone, statusCode)
}

func Test_GoneIsServerError(t *testing.T) {
	err := error.NewErrWrappedGone("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_GoneIsClientError(t *testing.T) {
	err := error.NewErrWrappedGone("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
