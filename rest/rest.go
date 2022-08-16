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

func (service *service) MakeGetRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakePostRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakePutRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}
func (service *service) MakeDeleteRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeGetRequestWithConfig(ctx *flat.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePostRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePutRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeDeleteRequestWithConfig(ctx *flat.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeGetRequestWithTimeout(ctx *flat.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePostRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakePutRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *service) MakeDeleteRequestWithTimeout(ctx *flat.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}
