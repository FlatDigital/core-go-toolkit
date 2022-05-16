package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// Failed Dependency (424)

// ErrFailDependency failed dependency error
type ErrFailDependency struct {
	s string
}

func (e ErrFailDependency) Error() string {
	return e.s
}

// newErrFailDependency returns a dependency fail error.
func newErrFailDependency(text string, a ...interface{}) error {
	return ErrFailDependency{fmt.Sprintf(text, a...)}
}

// NewErrWrappedFailDependency returns a wrapped dependency fail error.
func NewErrWrappedFailDependency(text string, a ...interface{}) Wrapper {
	return Wrap(newErrFailDependency(text, a...))
}

// ReturnFailDependencyError returns a dependency fail error
func ReturnFailDependencyError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusFailedDependency,
			Literal:   "FailDependecyError",
			Alertable: true,
		},
		Cause: err.Error(),
	})
}
