package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/v2/core/libs/go/errors"
)

// TooManyRequests (429)

// ErrTooManyRequests upgrade required error
type ErrTooManyRequests struct {
	s string
}

func (e ErrTooManyRequests) Error() string {
	return e.s
}

// newErrTooManyRequests returns an upgrade required error
func newErrTooManyRequests(text string, a ...interface{}) error {
	return ErrTooManyRequests{fmt.Sprintf(text, a...)}
}

// NewErrWrappedTooManyRequests returns a wrapped upgrade required error
func NewErrWrappedTooManyRequests(text string, a ...interface{}) Wrapper {
	return Wrap(newErrTooManyRequests(text, a...))
}

// ReturnTooManyRequestsError returns a Upgrade Required error
func ReturnTooManyRequestsError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusTooManyRequests,
			Literal:   "TooManyRequestsApiError",
			Alertable: false,
		},
		Cause: err.Error(),
	})
}
