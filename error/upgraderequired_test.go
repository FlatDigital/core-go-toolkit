package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedUpgradeRequired(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedUpgradeRequired("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrUpgradeRequired{}, err.WrappedErr())
}

func Test_ReturnUpgradeRequiredError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedUpgradeRequired("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusUpgradeRequired, rr.Code)
	ass.Equal("{\"error\":\"UpgradeRequiredApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetUpgradeRequiredStatusCode(t *testing.T) {
	err := error.NewErrWrappedUpgradeRequired("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusUpgradeRequired, statusCode)
}

func Test_UpgradeRequiredIsServerError(t *testing.T) {
	err := error.NewErrWrappedUpgradeRequired("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_UpgradeRequiredIsClientError(t *testing.T) {
	err := error.NewErrWrappedUpgradeRequired("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
