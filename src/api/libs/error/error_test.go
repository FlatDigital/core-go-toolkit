package error_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/src/api/libs/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ReturnError_DefaultError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	e := errors.New("forced for test")
	err := error.Wrap(e)

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusInternalServerError, rr.Code)
	ass.Equal("{\"error\":\"InternalServerApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
