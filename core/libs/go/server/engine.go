package server

import (
	"fmt"

	"github.com/atarantini/ginrequestid"
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// GroupPreffix is the preffix added to every exposed url
var GroupPreffix = "/flat"

// settings contains the relevant information that our server may use for initiating.
type settings struct {
	LogLevel    string
	PushMetrics bool
	Debug       bool
	AuthScopes  []string
	AppName     string
}

// Map with default server settings for each possible scope.
var envSettings = map[Environment]settings{
	EnvDevelop:    {LogLevel: "DEBUG", PushMetrics: true, Debug: true, AuthScopes: nil},
	EnvTest:       {LogLevel: "INFO", PushMetrics: true, Debug: true, AuthScopes: nil},
	EnvProduction: {LogLevel: "INFO", PushMetrics: true, Debug: false, AuthScopes: nil},
	EnvSandbox:    {LogLevel: "INFO", PushMetrics: true, Debug: false, AuthScopes: nil},
}

// RoutingGroup is a map of urls and functions for a given role.
type RoutingGroup map[Role]func(*gin.RouterGroup)

// Opt is a function for Server, it's used for optional modifiers used in package constructor
type Opt func(*Server)

// Server is our application main struct, it's basically a wrapper around a
// gin.Engine instance, with some functionality hidden for easier usage.
type Server struct {
	*gin.Engine
	Context ApplicationContext

	settings settings
}

// NewEngine configures the underlying gin.Engine struct of Server with a given scope, a
// RoutingGroup (exposed urls mapped to a valid Role) and accepts a list of options for
// specifying configuration options outside of the defaults for a given environment.
func NewEngine(scope string, routes RoutingGroup, opts ...Opt) (*Server, error) {

	// Tracer is used to send metrics to Datadog
	tracer.Start(
		tracer.WithDogstatsdAddress("datadog:8125"),
		tracer.WithRuntimeMetrics(),
		tracer.WithAgentAddr("datadog:8126"),
	)

	// Infer application context from scope
	ctx, err := ContextFromScopeString(scope)
	if err != nil {
		return nil, fmt.Errorf("error infering context from scope: %v", err)
	}

	// Check if the given routes are valid for the current application role
	if _, ok := routes[ctx.Role]; !ok {
		return nil, fmt.Errorf("given routes do not contain endpoints for the current application role")
	}

	// Get the Environment settings
	var envSet settings
	if envSet, err = getEnvironmentSettings(ctx); err != nil {
		return nil, err
	}

	// Create server with default configuration for current environment
	server := &Server{
		Context:  ctx,
		settings: envSet,
	}

	// Call option functions on instance before instantiating the server so that custom
	// options are taken into consideration.
	for _, opt := range opts {
		opt(server)
	}

	// Create a gin engine with debug or release config depending on the given settings
	if server.settings.Debug {
		gin.SetMode(gin.DebugMode)
		server.Engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		server.Engine = gin.New()
	}

	// Global server configuration, common to all environments
	server.NoRoute(NoRouteHandler())
	server.RedirectFixedPath = false
	server.RedirectTrailingSlash = false

	// Setup health check handler
	server.GET("/ping", HealthCheckHandler)

	// Using gintrace middleware for datadog tracing
	if len(server.settings.AppName) != 0 {
		server.Use(gintrace.Middleware(server.settings.AppName))
	}

	// Call the current Role group function with the current group as param
	// so that it loads the active urls.
	group := server.Group(GroupPreffix)

	group.Use(ginrequestid.RequestId())

	// Add support for test header used by API rules to all endpoints
	// group.Use(FlatAPIRules())

	fn := routes[ctx.Role]
	fn(group)

	return server, nil
}

func getEnvironmentSettings(ctx ApplicationContext) (settings, error) {
	value, ok := envSettings[ctx.Environment]
	if !ok {
		return settings{}, fmt.Errorf("given environment do not exist for the current application context")
	}
	return value, nil
}

// WithAuthScopes func sets up required application authentication scopes.
func WithAuthScopes(authScopes []string) Opt {
	return func(s *Server) {
		s.settings.AuthScopes = authScopes
	}
}

// WithLogLevel func sets application log level.
func WithLogLevel(logLevel string) Opt {
	return func(s *Server) {
		s.settings.LogLevel = logLevel
	}
}

// WithPushMetrics func sets push metrics options.
func WithPushMetrics(pushMetrics bool) Opt {
	return func(s *Server) {
		s.settings.PushMetrics = pushMetrics
	}
}

// WithDebug func sets debugging for webserver.
func WithDebug(debug bool) Opt {
	return func(s *Server) {
		s.settings.Debug = debug
	}
}

func WithAppName(name string) Opt {
	return func(s *Server) {
		s.settings.AppName = name
	}
}
