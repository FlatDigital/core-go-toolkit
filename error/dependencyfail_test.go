package error_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedFailDependency(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedFailDependency("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrFailDependency{}, err.WrappedErr())
}

func Test_ReturnFailDependencyError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedFailDependency("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusFailedDependency, rr.Code)
	ass.Equal("{\"error\":\"FailDependecyError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_ReturnFailDependencyErrorFromStatus(t *testing.T) {
	ass := assert.New(t)

	// given
	statusCode := http.StatusFailedDependency
	err := errors.New("forced for test")

	// when
	wrappedError := error.ReturnWrappedErrorFromStatus(statusCode, err)

	// then
	ass.NotNil(wrappedError)
	ass.IsType(error.ErrFailDependency{}, wrappedError.WrappedErr())
	ass.EqualError(wrappedError.WrappedErr(), err.Error())
}

func Test_GetFailDependencyStatusCode(t *testing.T) {
	err := error.NewErrWrappedFailDependency("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusFailedDependency, statusCode)
}

func Test_FailDependencyIsServerError(t *testing.T) {
	err := error.NewErrWrappedFailDependency("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_FailDependencyIsClientError(t *testing.T) {
	err := error.NewErrWrappedFailDependency("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
