package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/src/api/libs/core/libs/go/errors"
)

// UpgradeRequired

// ErrUpgradeRequired upgrade required error
type ErrUpgradeRequired struct {
	s string
}

func (e ErrUpgradeRequired) Error() string {
	return e.s
}

// newErrUpgradeRequired returns an upgrade required error
func newErrUpgradeRequired(text string, a ...interface{}) error {
	return ErrUpgradeRequired{fmt.Sprintf(text, a...)}
}

// NewErrWrappedUpgradeRequired returns a wrapped upgrade required error
func NewErrWrappedUpgradeRequired(text string, a ...interface{}) Wrapper {
	return Wrap(newErrUpgradeRequired(text, a...))
}

// ReturnUpgradeRequiredError returns a Upgrade Required error
func ReturnUpgradeRequiredError(c *gin.Context, err error) {
	// delegate this to GORDIK!
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusUpgradeRequired,
			Literal:   "UpgradeRequiredApiError",
			Alertable: false,
		},
		Cause: err.Error(),
	})
}
