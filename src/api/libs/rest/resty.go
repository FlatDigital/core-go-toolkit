package rest

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var (
// DefaultDialTimeout is the max interval of time the dialer will wait when
// executing the TCP handshake before returning a timeout error.
//
// This value is known and fixed within the internal network.
// DefaultDialTimeout = 300 * time.Millisecond

// DefaultKeepAliveProbeInterval is the interval at which the dialer sets the
// KeepAlive probe packet to be sent to assert the state of the connection.
// DefaultKeepAliveProbeInterval = 15 * time.Second
)

type restyService struct {
	restyClient *resty.Client
}

func NewRestyService() Rest {

	return &restyService{
		restyClient: resty.New(),
	}
}

func NewRestyServiceWithConfig(config ServiceConfig) Rest {
	rConfig := config.RequestConfig

	dialer := &net.Dialer{
		Timeout: rConfig.ConnectTimeout,
		// KeepAlive: DefaultKeepAliveProbeInterval,
		// DualStack: true,
	}

	transport := &http.Transport{
		DialContext: dialer.DialContext,
		// ForceAttemptHTTP2:     true,
		// IdleConnTimeout:       90 * time.Second,
		// MaxIdleConnsPerHost:   500,
		// Proxy:                 http.ProxyFromEnvironment,
		// ExpectContinueTimeout: 1 * time.Second,
		// TLSHandshakeTimeout:   10 * time.Second,
	}

	restyClient := resty.New()
	restyClient.
		// Set client timeout as per your need
		SetTimeout(rConfig.Timeout).
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		SetTransport(transport)

	return &restyService{
		restyClient: restyClient,
	}
}

func (service *restyService) MakeGetRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Get(url)

	if err != nil {
		// returns API error
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		// returns API error
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakePostRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Post(url)

	if err != nil {
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakePutRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	r, err := req.Put(url)

	if err != nil {
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakeDeleteRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Delete(url)

	if err != nil {
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakeGetRequestWithConfig(ctx *gin.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePostRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePutRequestWithConfig(ctx *gin.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeDeleteRequestWithConfig(ctx *gin.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeGetRequestWithTimeout(ctx *gin.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePostRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePutRequestWithTimeout(ctx *gin.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeDeleteRequestWithTimeout(ctx *gin.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}
