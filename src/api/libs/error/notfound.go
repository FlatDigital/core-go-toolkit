package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Not found (404)

// ErrNotFound not found error
type ErrNotFound struct {
	s string
}

func (e ErrNotFound) Error() string {
	return e.s
}

// NewErrNotFound returns a not found error.
func newErrNotFound(text string, a ...interface{}) error {
	return ErrNotFound{fmt.Sprintf(text, a...)}
}

// NewErrWrappedNotFound returns a wrapped not found error.
func NewErrWrappedNotFound(text string, a ...interface{}) Wrapper {
	return Wrap(newErrNotFound(text, a...))
}

// ReturnNotFoundError returns a Not Found error
func ReturnNotFoundError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.NotFoundApiError,
		Cause: err.Error(),
	})
}
