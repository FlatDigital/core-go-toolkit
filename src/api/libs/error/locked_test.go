package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/src/api/libs/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrWrappedLocked(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedLocked("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrLocked{}, err.WrappedErr())
}

func Test_ReturnLockedError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedLocked("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusLocked, rr.Code)
	ass.Equal("{\"error\":\"LockedApiError\",\"cause\":\"forced for test\"}", rr.Body.String())
}
