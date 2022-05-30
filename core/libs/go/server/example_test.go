package server_test

import (
	"log"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/core/libs/go/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_ExampleNewEngine(t *testing.T) {
	ass := assert.New(t)

	routes := server.RoutingGroup{
		server.RoleRead: func(g *gin.RouterGroup) {
			g.GET("/read", func(c *gin.Context) {})
		},
		server.RoleWrite: func(g *gin.RouterGroup) {
			g.POST("/writer", func(c *gin.Context) {})
		},
	}

	srv, err := server.NewEngine(
		"test-read-feature-branch",
		routes,

		// Optional configuration override
		server.WithDebug(false),
		server.WithPushMetrics(true),
	)

	if err != nil {
		log.Fatal(err)
	}

	ass.Equal(server.EnvTest, srv.Context.Environment)
	ass.Equal(server.RoleRead, srv.Context.Role)
	ass.Equal("feature-branch", srv.Context.Tag)
}
