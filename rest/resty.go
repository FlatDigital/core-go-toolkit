package rest

import (
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/FlatDigital/core-go-toolkit/v2/core/flat"
	"github.com/go-resty/resty/v2"
)

const (
	MakeGetRequest           string = "MakeGetRequest"
	MakePostRequest          string = "MakePostRequest"
	MakePutRequest           string = "MakePutRequest"
	MakePatchRequest         string = "MakePatchRequest"
	MakeDeleteRequest        string = "MakeDeleteRequest"
	MakeGetRequestWithConfig string = "MakeGetRequestWithConfig"
)

type restyService struct {
	restyClient         *resty.Client
	datadogMetricPrefix string
}

func NewRestyService(metricPrefix string) Rest {
	return &restyService{
		restyClient:         resty.New(),
		datadogMetricPrefix: metricPrefix,
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
		restyClient:         restyClient,
		datadogMetricPrefix: config.DatadogMetricPrefix,
	}
}

func (service *restyService) MakeGetRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Get(url)
	return service.evaluateResponse(ctx, url, response, MakeGetRequest, start, err)
}

func (service *restyService) MakePostRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Post(url)
	return service.evaluateResponse(ctx, url, response, MakePostRequest, start, err)
}

func (service *restyService) MakePutRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Put(url)
	return service.evaluateResponse(ctx, url, response, MakePutRequest, start, err)
}

func (service *restyService) MakePatchRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Patch(url)
	return service.evaluateResponse(ctx, url, response, MakePutRequest, start, err)
}

func (service *restyService) MakeDeleteRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Delete(url)
	return service.evaluateResponse(ctx, url, response, MakeDeleteRequest, start, err)
}

func (service *restyService) MakeGetRequestWithConfig(ctx *flat.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	start := time.Now()
	client := service.restyClient
	client.SetTimeout(config.Timeout)

	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Get(url)
	return service.evaluateResponse(ctx, url, response, MakeGetRequestWithConfig, start, err)
}

func (service *restyService) MakePostRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePutRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeDeleteRequestWithConfig(ctx *flat.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeGetRequestWithTimeout(ctx *flat.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePostRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakePutRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) MakeDeleteRequestWithTimeout(ctx *flat.Context, url string, headers http.Header, timeout time.Duration) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (service *restyService) evaluateResponse(ctx *flat.Context, url string, response *resty.Response, resource string,
	start time.Time, err error) (int, []byte, error) {

	if response == nil {
		return 0, nil, errResponseNotReceived
	}
	if err != nil {
		return response.StatusCode(), response.Body(), err
	}
	if !(response.StatusCode() >= http.StatusOK && response.StatusCode() <= http.StatusIMUsed) {
		return response.StatusCode(), response.Body(), errors.New(response.Status())
	}

	return response.StatusCode(), response.Body(), nil
}
