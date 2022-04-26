package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Bad request (400)

// ErrBadRequest bad request error
type ErrBadRequest struct {
	s string
}

func (e ErrBadRequest) Error() string {
	return e.s
}

// newErrBadRequest returns a bad request error.
func newErrBadRequest(text string, a ...interface{}) error {
	return ErrBadRequest{fmt.Sprintf(text, a...)}
}

// NewErrWrappedBadRequest returns a wrapped bad request error.
func NewErrWrappedBadRequest(text string, a ...interface{}) Wrapper {
	return Wrap(newErrBadRequest(text, a...))
}

// ReturnBadRequestError returns a Bad Request error
func ReturnBadRequestError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.BadRequestApiError,
		Cause: err.Error(),
	})
}
