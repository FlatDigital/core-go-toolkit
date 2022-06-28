package rest

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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
	}

	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
	}

	restyClient := resty.New()
	restyClient.
		// Overrode default transport layer
		SetTransport(transport)

		// TODO: Add values to RequestConfig to support the following methods

		// SetRetryCount(3).
		// SetRetryWaitTime(100 * time.Millisecond).
		// SetRetryMaxWaitTime(2 * time.Second).
		// SetRetryAfter(nil).
		// AddRetryCondition(nil).
		// AddRetryAfterErrorCondition().
		// AddRetryHook(nil).

	if !rConfig.DisableTimeout {
		// Set client timeout as per your need
		restyClient.SetTimeout(rConfig.Timeout)
	}

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
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakePostRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

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
	client := service.restyClient
	client.SetTimeout(config.Timeout)

	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Get(url)

	if err != nil {
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		return r.StatusCode(), r.Body(), err
	}

	return r.StatusCode(), r.Body(), nil
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
