package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// UnprocessableEntity

// ErrUnprocessableEntity unprocessable entity error
type ErrUnprocessableEntity struct {
	s string
}

func (e ErrUnprocessableEntity) Error() string {
	return e.s
}

// newErrUnprocessableEntity returns an unprocessable entity error
func newErrUnprocessableEntity(text string, a ...interface{}) error {
	return ErrUnprocessableEntity{fmt.Sprintf(text, a...)}
}

// NewErrWrappedUnprocessableEntity returns a wrapped unprocessable entity error
func NewErrWrappedUnprocessableEntity(text string, a ...interface{}) Wrapper {
	return Wrap(newErrUnprocessableEntity(text, a...))
}

// ReturnUnprocessableEntityError returns a Unauthorized error
func ReturnUnprocessableEntityError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.UnprocessableEntityApiError,
		Cause: err.Error(),
	})
}
