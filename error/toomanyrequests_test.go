package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedTooManyRequests(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedTooManyRequests("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrTooManyRequests{}, err.WrappedErr())
}

func Test_ReturnTooManyRequestsError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedTooManyRequests("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusTooManyRequests, rr.Code)
	ass.Equal("{\"error\":\"TooManyRequestsApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
