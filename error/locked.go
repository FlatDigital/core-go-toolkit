package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/v2/core/libs/go/errors"
)

// Locked (423)

// ErrLocked Locked error
type ErrLocked struct {
	s string
}

func (e ErrLocked) Error() string {
	return e.s
}

// newErrLocked returns a Locked error.
func newErrLocked(text string, a ...interface{}) error {
	return ErrLocked{fmt.Sprintf(text, a...)}
}

// NewErrWrappedLocked returns a wrapped Locked error.
func NewErrWrappedLocked(text string, a ...interface{}) Wrapper {
	return Wrap(newErrLocked(text, a...))
}

// ReturnLockedError returns a Locked error
func ReturnLockedError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusLocked,
			Literal:   "LockedApiError",
			Alertable: false,
		},
		Cause: err.Error(),
	})
}
