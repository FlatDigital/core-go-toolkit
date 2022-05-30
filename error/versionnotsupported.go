package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/core/libs/go/errors"
)

// VersionNotSupported

// ErrVersionNotSupported version not supported error
type ErrVersionNotSupported struct {
	s string
}

func (e ErrVersionNotSupported) Error() string {
	return e.s
}

// newErrVersionNotSupported returns an version not supported error
func newErrVersionNotSupported(text string, a ...interface{}) error {
	return ErrVersionNotSupported{fmt.Sprintf(text, a...)}
}

// NewErrWrappedVersionNotSupported returns a wrapped version not supported error
func NewErrWrappedVersionNotSupported(text string, a ...interface{}) Wrapper {
	return Wrap(newErrVersionNotSupported(text, a...))
}

// ReturnVersionNotSupportedError returns a Http Version Not Supported
func ReturnVersionNotSupportedError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusHTTPVersionNotSupported,
			Literal:   "HTTPVersionNotSupportedApiError",
			Alertable: false,
		},
		Cause: err.Error(),
	})
}
