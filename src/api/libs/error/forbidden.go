package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Forbidden (403)

// ErrForbidden forbidden error
type ErrForbidden struct {
	s string
}

func (e ErrForbidden) Error() string {
	return e.s
}

// NewErrForbidden returns an forbidden error
func newErrForbidden(text string, a ...interface{}) error {
	return ErrForbidden{fmt.Sprintf(text, a...)}
}

// NewErrWrappedForbidden returns a wrapped forbidden error
func NewErrWrappedForbidden(text string, a ...interface{}) Wrapper {
	return Wrap(newErrForbidden(text, a...))
}

// ReturnForbiddenError returns a forbidden error
func ReturnForbiddenError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.ForbiddenApiError,
		Cause: err.Error(),
	})
}
