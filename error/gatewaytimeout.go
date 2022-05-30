package error

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	gkErrors "github.com/FlatDigital/core-go-toolkit/core/libs/go/errors"
)

// GatewayTimeout (504)

// ErrGatewayTimeout GatewayTimeout error
type ErrGatewayTimeout struct {
	s string
}

func (e ErrGatewayTimeout) Error() string {
	return e.s
}

// newErrGatewayTimeout returns a GatewayTimeout error.
func newErrGatewayTimeout(text string, a ...interface{}) error {
	return ErrGatewayTimeout{fmt.Sprintf(text, a...)}
}

// NewErrWrappedGatewayTimeout returns a wrapped GatewayTimeout error.
func NewErrWrappedGatewayTimeout(text string, a ...interface{}) Wrapper {
	return Wrap(newErrGatewayTimeout(text, a...))
}

// ReturnGatewayTimeoutError returns a GatewayTimeout error
func ReturnGatewayTimeoutError(c *gin.Context, err error) {
	gkErrors.ReturnError(c, &gkErrors.Error{
		Code: gkErrors.ErrorCode{
			Status:    http.StatusGatewayTimeout,
			Literal:   "GatewayTimeoutApiError",
			Alertable: true,
		},
		Cause: err.Error(),
	})
}
