package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// InternalServerError (500)

// ErrInternalServerError InternalServerError error
type ErrInternalServerError struct {
	s string
}

func (e ErrInternalServerError) Error() string {
	return e.s
}

// newErrInternalServerError returns a InternalServerError error.
func newErrInternalServerError(text string, a ...interface{}) error {
	return ErrInternalServerError{fmt.Sprintf(text, a...)}
}

// NewErrWrappedInternalServerError returns a wrapped InternalServerError error.
func NewErrWrappedInternalServerError(text string, a ...interface{}) Wrapper {
	return Wrap(newErrInternalServerError(text, a...))
}

// ReturnInternalServerError returns a InternalServerError error
func ReturnInternalServerError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.InternalServerApiError,
		Cause: err.Error(),
	})
}
