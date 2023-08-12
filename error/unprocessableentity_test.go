package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedUnprocessableEntity(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedUnprocessableEntity("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrUnprocessableEntity{}, err.WrappedErr())
}

func Test_ReturnUnprocessableEntityError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedUnprocessableEntity("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusUnprocessableEntity, rr.Code)
	ass.Equal("{\"error\":\"UnprocessableEntityApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
