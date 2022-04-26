package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Gone (410)

// ErrGone gone error
type ErrGone struct {
	s string
}

func (e ErrGone) Error() string {
	return e.s
}

// newErrGone returns a gone error.
func newErrGone(text string, a ...interface{}) error {
	return ErrGone{fmt.Sprintf(text, a...)}
}

// NewErrWrappedGone returns a wrapped gone error.
func NewErrWrappedGone(text string, a ...interface{}) Wrapper {
	return Wrap(newErrGone(text, a...))
}

// ReturnGoneError returns a Gone error
func ReturnGoneError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusGone,
			Literal:   "GoneApiError",
			Alertable: false,
		},
		Cause: err.Error(),
	})
}
