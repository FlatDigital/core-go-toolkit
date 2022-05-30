package error

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/core/libs/go/errors"
)

// Unavailable for legal reasons (451)

// ErrUnavailableForLegalReasons unavailable for legal reasons error
type ErrUnavailableForLegalReasons struct {
	s string
}

func (e ErrUnavailableForLegalReasons) Error() string {
	return e.s
}

// newErrUnavailableForLegalReasons returns a unavailable for legal reasons error
func newErrUnavailableForLegalReasons(text string, a ...interface{}) error {
	return ErrUnavailableForLegalReasons{fmt.Sprintf(text, a...)}
}

// NewErrWrappedUnavailableForLegalReasons returns a wrapped unavailable for legal reasons error
func NewErrWrappedUnavailableForLegalReasons(text string, a ...interface{}) Wrapper {
	return Wrap(newErrUnavailableForLegalReasons(text, a...))
}

// ReturnUnavailableForLegalReasonsError returns a unavailable for legal reasons error.
func ReturnUnavailableForLegalReasonsError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code:  gkErrors.UnavailableForLegalReasonsError,
		Cause: err.Error(),
	})
}
