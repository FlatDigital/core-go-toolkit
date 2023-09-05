package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedBadGateway(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedBadGateway("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrBadGateway{}, err.WrappedErr())
}

func Test_ReturnBadGatewayError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedBadGateway("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusBadGateway, rr.Code)
	ass.Equal("{\"error\":\"BadGatewayApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetBadGatewayStatusCode(t *testing.T) {
	err := error.NewErrWrappedBadGateway("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusBadGateway, statusCode)
}

func Test_BadGatewayIsServerError(t *testing.T) {
	err := error.NewErrWrappedBadGateway("forced for test")

	isClientError := error.IsServerError(err)

	assert.True(t, isClientError)
}

func Test_BadGatewayIsClientError(t *testing.T) {
	err := error.NewErrWrappedBadGateway("forced for test")

	isClientError := error.IsClientError(err)

	assert.False(t, isClientError)
}
