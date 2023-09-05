package error_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/error"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_NewErrUnavailableForLegalReasons(t *testing.T) {
	// given
	ass := assert.New(t)
	err := error.NewErrWrappedUnavailableForLegalReasons("forced for test")

	// then
	ass.NotNil(err)
	ass.Error(err.WrappedErr())
	ass.Equal("forced for test", err.Details())
	ass.IsType(error.ErrUnavailableForLegalReasons{}, err.WrappedErr())
}

func Test_ReturnUnavailableForLegalReasonsError(t *testing.T) {
	// given
	ass := assert.New(t)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	err := error.NewErrWrappedUnavailableForLegalReasons("forced for test")

	// when
	error.ReturnError(c, err)

	// then
	ass.Equal(http.StatusUnavailableForLegalReasons, rr.Code)
	ass.Equal("{\"error\":\"UnavailableForLegalReasonsError\",\"cause\":\"forced for test\"}", rr.Body.String())
}

func Test_GetUnavailableForLegalReasonsStatusCode(t *testing.T) {
	err := error.NewErrWrappedUnavailableForLegalReasons("forced for test")

	statusCode := error.GetStatusCode(err)

	assert.Equal(t, http.StatusUnavailableForLegalReasons, statusCode)
}

func Test_UnavailableForLegalReasonsIsServerError(t *testing.T) {
	err := error.NewErrWrappedUnavailableForLegalReasons("forced for test")

	isClientError := error.IsServerError(err)

	assert.False(t, isClientError)
}

func Test_UnavailableForLegalReasonsIsClientError(t *testing.T) {
	err := error.NewErrWrappedUnavailableForLegalReasons("forced for test")

	isClientError := error.IsClientError(err)

	assert.True(t, isClientError)
}
