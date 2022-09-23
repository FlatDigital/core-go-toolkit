package rest

import (
	"errors"
	"net/http"
	"time"

	"github.com/FlatDigital/core-go-toolkit/core/flat"
)

// Rest interface for rest service
type Rest interface {
	MakeGetRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error)
	MakePostRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
	MakePutRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
	MakePatchRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
	MakeDeleteRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error)

	MakeGetRequestWithConfig(ctx *flat.Context, url string, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakePostRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakePutRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header,
		config RequestConfig) (int, []byte, error)
	MakeDeleteRequestWithConfig(ctx *flat.Context, url string, headers http.Header,
		config RequestConfig) (int, []byte, error)

	MakeGetRequestWithTimeout(ctx *flat.Context, url string, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakePostRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakePutRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakeDeleteRequestWithTimeout(ctx *flat.Context, url string, headers http.Header,
		timeout time.Duration) (int, []byte, error)
	MakeGetRequestWithBody(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error)
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
	DisableTimeout bool
	Timeout        time.Duration
	ConnectTimeout time.Duration
}
