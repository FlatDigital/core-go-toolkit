package rest

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Rest interface for rest service
type Rest interface {
	MakeGetRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error)
	MakePostRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
	MakePutRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
	MakeDeleteRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error)

	MakeGetRequestWithConfig(ctx *gin.Context, url string, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakePostRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakePutRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakeDeleteRequestWithConfig(ctx *gin.Context, url string, headers http.Header,
		config RequestConfig) (int, []byte, error)

	MakeGetRequestWithTimeout(ctx *gin.Context, url string, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakePostRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakePutRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakeDeleteRequestWithTimeout(ctx *gin.Context, url string, headers http.Header,
		timeout time.Duration) (int, []byte, error)
}

var (
	// The changes here will affect to all requests
	defaultRequestConfig = RequestConfig{
		DisableTimeout: false,
		Timeout:        3 * time.Second,
		ConnectTimeout: 1500 * time.Millisecond,
	}
	errResponseNotReceived = errors.New("response not received")
)

type ServiceConfig struct {
	BaseURL             string
	MaxIdleConnsPerHost int
	RequestConfig       *RequestConfig
	DatadogMetricPrefix string
}

type RequestConfig struct {
	DisableTimeout bool // See @RequestBuilder.DisableTimeout: Disable timeout and default timeout = no timeout
	Timeout        time.Duration
	ConnectTimeout time.Duration
}

type service struct {
	baseURL string
}

func NewRestService(config ServiceConfig) Rest {

	if config.RequestConfig != nil {
		rConfig := config.RequestConfig
		defaultRequestConfig.DisableTimeout = rConfig.DisableTimeout
		if rConfig.Timeout > 0 {
			defaultRequestConfig.Timeout = rConfig.Timeout
		}
		if rConfig.ConnectTimeout > 0 {
			defaultRequestConfig.ConnectTimeout = rConfig.ConnectTimeout
		}
	}

	return &service{
		baseURL: config.BaseURL,
	}
}

func (service *service) MakeGetRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakePostRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakePutRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakeDeleteRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeGetRequestWithConfig(ctx *gin.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePostRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePutRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeDeleteRequestWithConfig(ctx *gin.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeGetRequestWithTimeout(ctx *gin.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePostRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePutRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeDeleteRequestWithTimeout(ctx *gin.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}
