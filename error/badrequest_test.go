package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedBadRequest(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedBadRequest("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrBadRequest{}, err.WrappedErr())
}

func Test_ReturnBadRequestError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedBadRequest("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusBadRequest, rr.Code)
	ass.Equal("{\"error\":\"BadRequestApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
