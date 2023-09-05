package error

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	goErr "github.com/go-errors/errors"
)

// Wrapper defines the basic functionality for the error wrapped
type Wrapper interface {

	// Details returns the error's message
	Details() string

	// Equal compares whether the error is equal to a given error
	Equal(originalError interface{}) bool

	// Stack returns the stack trace related to the error
	Stack() string

	// WrappedErr returns the wrapped error in the container
	WrappedErr() error
}

// errWrapper defines a struct with a container that implements Wrapper interface
type errWrapper struct {
	*container
}

// container defines an error container
type container struct {
	ext     *goErr.Error
	wrapped error
}

func (errContainer container) extErr() *goErr.Error {
	return errContainer.ext
}

func (errContainer container) wrappedErr() error {
	return errContainer.wrapped
}

func (errContainer container) isEqual(originalError interface{}) bool {
	return reflect.TypeOf(errContainer.wrappedErr()) == reflect.TypeOf(originalError)
}

// Implementations

// New returns a new err container from the given error
func New(format string, params ...interface{}) Wrapper {
	var err error
	if params == nil {
		err = fmt.Errorf(format)
	} else {
		err = fmt.Errorf(format, params...)
	}

	return Wrap(err)
}

// Wrap wraps an error in an error container
func Wrap(err error) Wrapper {
	if err == nil {
		return nil
	}

	return errWrapper{
		container: &container{
			ext:     goErr.New(err),
			wrapped: err,
		},
	}
}

// Details returns the error's message
func (errWrapper errWrapper) Details() string {
	return errWrapper.container.extErr().Error()
}

// Equal compares whether the error wrapped is equal to a given error
func (errWrapper errWrapper) Equal(originalError interface{}) bool {
	return errWrapper.container.isEqual(originalError)
}

// Stack returns the stack trace related to the error
func (errWrapper errWrapper) Stack() string {
	if errWrapper.container.extErr() != nil {
		stackTrace := errWrapper.container.extErr().ErrorStack()
		stackTrace = strings.Replace(stackTrace, "\n", " - ", -1)
		return stackTrace
	}
	return "no stack trace info in error"
}

// WrappedErr returns the error wrapped in the container
func (errWrapper errWrapper) WrappedErr() error {
	return errWrapper.container.wrappedErr()
}

//

// ReturnError returns the corresponding http error based on the wrapped err type
func ReturnError(c *gin.Context, errWrapped Wrapper) {
	err := errWrapped.WrappedErr()
	switch err.(type) {

	case ErrBadGateway:
		ReturnBadGatewayError(c, err)
	case ErrBadRequest:
		ReturnBadRequestError(c, err)
	case ErrConflict:
		ReturnConflictError(c, err)
	case ErrForbidden:
		ReturnForbiddenError(c, err)
	case ErrGatewayTimeout:
		ReturnGatewayTimeoutError(c, err)
	case ErrGone:
		ReturnGoneError(c, err)
	case ErrInternalServerError:
		ReturnInternalServerError(c, err)
	case ErrLocked:
		ReturnLockedError(c, err)
	case ErrNotFound:
		ReturnNotFoundError(c, err)
	case ErrNotImplemented:
		ReturnNotImplementedError(c, err)
	case ErrTooManyRequests:
		ReturnTooManyRequestsError(c, err)
	case ErrUnauthorized:
		ReturnUnauthorizedError(c, err)
	case ErrUnprocessableEntity:
		ReturnUnprocessableEntityError(c, err)
	case ErrUpgradeRequired:
		ReturnUpgradeRequiredError(c, err)
	case ErrVersionNotSupported:
		ReturnVersionNotSupportedError(c, err)
	case ErrFailDependency:
		ReturnFailDependencyError(c, err)
	case ErrUnavailableForLegalReasons:
		ReturnUnavailableForLegalReasonsError(c, err)

	default:
		ReturnInternalServerError(c, err)
	}
}

// ReturnWrappedErrorFromStatus returns a wrapped error from a status code
func ReturnWrappedErrorFromStatus(statusCode int, err error) Wrapper {
	switch statusCode {

	case http.StatusBadGateway:
		return NewErrWrappedBadGateway(err.Error())
	case http.StatusBadRequest:
		return NewErrWrappedBadRequest(err.Error())
	case http.StatusConflict:
		return NewErrWrappedConflict(err.Error())
	case http.StatusForbidden:
		return NewErrWrappedForbidden(err.Error())
	case http.StatusGatewayTimeout:
		return NewErrWrappedGatewayTimeout(err.Error())
	case http.StatusLocked:
		return NewErrWrappedLocked(err.Error())
	case http.StatusNotFound:
		return NewErrWrappedNotFound(err.Error())
	case http.StatusNotImplemented:
		return NewErrWrappedNotImplemented(err.Error())
	case http.StatusTooManyRequests:
		return NewErrWrappedTooManyRequests(err.Error())
	case http.StatusUnauthorized:
		return NewErrWrappedUnauthorized(err.Error())
	case http.StatusUnprocessableEntity:
		return NewErrWrappedUnprocessableEntity(err.Error())
	case http.StatusUpgradeRequired:
		return NewErrWrappedUpgradeRequired(err.Error())
	case http.StatusHTTPVersionNotSupported:
		return NewErrWrappedVersionNotSupported(err.Error())
	case http.StatusFailedDependency:
		return NewErrWrappedFailDependency(err.Error())

	default:
		return NewErrWrappedInternalServerError(err.Error())
	}
}

func GetStatusCode(errWrapped Wrapper) int {
	err := errWrapped.WrappedErr()
	switch err.(type) {

	case ErrBadGateway:
		return http.StatusBadGateway
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrConflict:
		return http.StatusConflict
	case ErrForbidden:
		return http.StatusForbidden
	case ErrGatewayTimeout:
		return http.StatusGatewayTimeout
	case ErrGone:
		return http.StatusGone
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrLocked:
		return http.StatusLocked
	case ErrNotFound:
		return http.StatusNotFound
	case ErrNotImplemented:
		return http.StatusNotImplemented
	case ErrTooManyRequests:
		return http.StatusTooManyRequests
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrUnprocessableEntity:
		return http.StatusUnprocessableEntity
	case ErrUpgradeRequired:
		return http.StatusUpgradeRequired
	case ErrVersionNotSupported:
		return http.StatusHTTPVersionNotSupported
	case ErrFailDependency:
		return http.StatusFailedDependency
	case ErrUnavailableForLegalReasons:
		return http.StatusUnavailableForLegalReasons

	default:
		return http.StatusInternalServerError
	}
}

func IsServerError(errWrapped Wrapper) bool {
	return GetStatusCode(errWrapped) >= http.StatusInternalServerError
}

func IsClientError(errWrapped Wrapper) bool {
	statusCode := GetStatusCode(errWrapped)

	return statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError
}
