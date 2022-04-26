package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// gkErrors "github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
	gkErrors "github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
)

// BadGateway (502)

// ErrBadGateway BadGateway error
type ErrBadGateway struct {
	s string
}

func (e ErrBadGateway) Error() string {
	return e.s
}

// newErrBadGateway returns a BadGateway error.
func newErrBadGateway(text string, a ...interface{}) error {
	return ErrBadGateway{fmt.Sprintf(text, a...)}
}

// NewErrWrappedBadGateway returns a wrapped BadGateway error.
func NewErrWrappedBadGateway(text string, a ...interface{}) Wrapper {
	return Wrap(newErrBadGateway(text, a...))
}

// ReturnBadGatewayError returns a BadGateway error
func ReturnBadGatewayError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.BadGatewayApiError,
		Cause: err.Error(),
	})
}
