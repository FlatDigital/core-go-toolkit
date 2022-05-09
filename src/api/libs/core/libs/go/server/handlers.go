package server

import (
	"fmt"
	"net/http"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/core/libs/go/errors"
	"github.com/gin-gonic/gin"
)

// NoRouteHandler is a default handler that's usually used in conjunction with
// gins NoRoute method.
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		errors.ReturnError(c, &errors.Error{
			Code:    errors.NotFoundApiError,
			Message: fmt.Sprintf("Resource %s not found.", c.Request.URL.Path),
			Values: map[string]string{
				"resource": c.Request.URL.Path,
			},
		})

		c.Abort()
	}
}

// HealthCheckHandler is a default handler that's used for checking if a
// given application instance is accepting requests.
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
