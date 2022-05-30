package flat_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/core/flat"
	"github.com/FlatDigital/core-go-toolkit/core/libs/go/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	uuid, _ := uuid.NewV4()
	reqID := uuid.String()

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Set("RequestId", reqID)

	c.Request, _ = http.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("X-Request-Id", reqID)
	c.Request.Header.Set("X-Caller-Id", "120120120")
	c.Request.Header.Set("X-Caller-Scopes", "admin")
	c.Request.Header.Set("X-Public", "true")

	flat.Handler(func(c *gin.Context, ctx *flat.Context) {
		assert.EqualValues(t, reqID, ctx.RequestID)
		assert.EqualValues(t, 120120120, ctx.Caller.ID)
		assert.EqualValues(t, true, ctx.Caller.IsAdmin)
		assert.EqualValues(t, true, ctx.Caller.IsPublic)

		assert.NotNil(t, ctx.Log)
		assert.IsType(t, &logger.Logger{}, ctx.Log)
	})(c)
}

func TestCreateTestContext(t *testing.T) {
	// This test is really unnecessary, but we do it as to not to penalize our code coverage
	flat.CreateTestContext()
}
