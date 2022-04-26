package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
)

// NotImplemented (501)

// ErrNotImplemented NotImplemented error
type ErrNotImplemented struct {
	s string
}

func (e ErrNotImplemented) Error() string {
	return e.s
}

// newErrNotImplemented returns a NotImplemented error.
func newErrNotImplemented(text string, a ...interface{}) error {
	return ErrNotImplemented{fmt.Sprintf(text, a...)}
}

// NewErrWrappedNotImplemented returns a wrapped NotImplemented error.
func NewErrWrappedNotImplemented(text string, a ...interface{}) Wrapper {
	return Wrap(newErrNotImplemented(text, a...))
}

// ReturnNotImplementedError returns a NotImplemented error
func ReturnNotImplementedError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusNotImplemented,
			Literal:   "NotImplementedApiError",
			Alertable: true,
		},
		Cause: err.Error(),
	})
}
