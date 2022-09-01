package flat

import (
	"net/http"
	"net/mail"
	"reflect"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/FlatDigital/core-go-toolkit/core/libs/go/logger"
)

// Caller is the type that contains the information inside a request that
// represents the user that generated it.
type Caller struct {
	ID       string
	IsAdmin  bool
	IsPublic bool
	Scopes   []string
}

// Context contains all the resources we use during a given request
type Context struct {
	ClientID    string
	ClientEmail string
	Caller      Caller
	RequestID   string
	Log         *logger.Logger
}

// HandlerFunc defines the signature of our http handlers
type HandlerFunc func(*gin.Context, *Context)

// Handler receives a HandlerFunc and allows it to be called from inside gin
// where a gin.HandlerFunc is expected.
func Handler(f HandlerFunc) gin.HandlerFunc {
	// Get caller function name so that we can rename newrelic transaction
	callerName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()

	return func(c *gin.Context) {
		context := TransformGinContextToGK(c, callerName)

		f(c, context)
	}
}

// CreateTestContext returns a MPCS Context ready to use for testing purposes. The
// context is only populated with a functioning logger and a valid request id.
// If more information is required, then the user should add it in its end.
func CreateTestContext() *Context {
	reqID, _ := uuid.NewV4()

	return &Context{
		ClientEmail: "test@test.com",
		RequestID:   reqID.String(),
		Log: &logger.Logger{
			Attributes: logger.Attrs{"request_id": reqID},
		},
	}
}

// TransformGinContextToGK recives a gin Context and transform it to gk Context
func TransformGinContextToGK(c *gin.Context, callerName string) *Context {
	rawCallerID := GetCaller(c.Request)
	clientID := GetClientId(c.Request)
	clientEmail := GetClientEmail(c.Request)

	// If we can't parse callerID then it remains 0
	// callerID, _ := strconv.ParseUint(rawCallerID, 10, 64)
	callerID := rawCallerID

	reqID := c.GetString("RequestId")

	context := &Context{
		Caller: Caller{
			ID:       callerID,
			IsAdmin:  IsCallerAdmin(c.Request),
			IsPublic: IsPublic(c.Request),
			Scopes:   GetCallerScopes(c.Request),
		},
		ClientID:    clientID,
		ClientEmail: clientEmail,
		RequestID:   reqID,
		Log: &logger.Logger{
			Attributes: logger.Attrs{"request_id": reqID},
		},
	}

	return context
}

//

func GetCaller(request *http.Request) string {
	if callerId := request.Header.Get("X-Caller-Id"); callerId != "" {
		return callerId
	}
	return request.URL.Query().Get("caller.id")
}

func GetClientId(request *http.Request) string {
	if clientId := request.Header.Get("X-Client-Id"); clientId != "" {
		return clientId
	}
	return request.URL.Query().Get("client.id")
}

func GetClientEmail(request *http.Request) string {
	email := request.Header.Get("X-Client-Email")
	if email == "" {
		email = request.URL.Query().Get("client.email")
		if email == "" {
			return ""
		}
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}

	return email
}

func IsPublic(request *http.Request) bool {
	return strings.ToLower(request.Header.Get("X-Public")) == "true"
}

func IsCallerAdmin(request *http.Request) bool {
	if scopes := GetCallerScopes(request); len(scopes) > 0 {
		for i := 0; i < len(scopes); i++ {
			if strings.ToLower(scopes[i]) == "admin" {
				return true
			}
		}
	}

	return false
}

func GetCallerScopes(request *http.Request) []string {
	if callerScopes := request.Header.Get("X-Caller-Scopes"); callerScopes != "" {
		return strings.Split(callerScopes, ",")

	} else {
		if callerScopes := request.URL.Query().Get("caller.scopes"); callerScopes != "" {
			return strings.Split(callerScopes, ",")

		} else {
			return make([]string, 0)
		}
	}
}
