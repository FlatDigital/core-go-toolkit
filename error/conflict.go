package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/v2/core/libs/go/errors"
)

// Conflict (409)

// ErrConflict conflict error
type ErrConflict struct {
	s string
}

func (e ErrConflict) Error() string {
	return e.s
}

// newErrConflict returns a conflict error.
func newErrConflict(text string, a ...interface{}) error {
	return ErrConflict{fmt.Sprintf(text, a...)}
}

// NewErrWrappedConflict returns a wrapped conflict error.
func NewErrWrappedConflict(text string, a ...interface{}) Wrapper {
	return Wrap(newErrConflict(text, a...))
}

// ReturnConflictError returns a conflct error
func ReturnConflictError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.ResourceConflictApiError,
		Cause: err.Error(),
	})
}
