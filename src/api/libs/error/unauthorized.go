package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Unauthorized (401)

// ErrUnauthorized unauthorized error
type ErrUnauthorized struct {
	s string
}

func (e ErrUnauthorized) Error() string {
	return e.s
}

// NewErrUnauthorized returns an unauthorized error
func newErrUnauthorized(text string, a ...interface{}) error {
	return ErrUnauthorized{fmt.Sprintf(text, a...)}
}

// NewErrWrappedUnauthorized returns a wrapped unauthorized error
func NewErrWrappedUnauthorized(text string, a ...interface{}) Wrapper {
	return Wrap(newErrUnauthorized(text, a...))
}

// ReturnUnauthorizedError returns a Unauthorized error
func ReturnUnauthorizedError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.AuthorizationApiError,
		Cause: err.Error(),
	})
}
