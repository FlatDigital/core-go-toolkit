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

func Test_NewErrWrappedFailDependecy(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedFailDependency("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrFailDependency{}, err.WrappedErr())
}

func Test_ReturnFailDependecyError(t *testing.T) {
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

func Test_ReturnFailDependecyErrorFromStatus(t *testing.T) {
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
