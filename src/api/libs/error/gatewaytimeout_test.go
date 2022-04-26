package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedGatewayTimeout(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedGatewayTimeout("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrGatewayTimeout{}, err.WrappedErr())
}

func Test_ReturnGatewayTimeoutError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedGatewayTimeout("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusGatewayTimeout, rr.Code)
	ass.Equal("{\"error\":\"GatewayTimeoutApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
