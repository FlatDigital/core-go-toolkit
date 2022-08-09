package rest

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/FlatDigital/core-go-toolkit/godog"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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

func (service *restyService) MakeGetRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	start := time.Now()
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Get(url)

	if err != nil {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeGetRequest", err)
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeGetRequest", errors.New(http.StatusText(r.StatusCode())))
		return r.StatusCode(), r.Body(), err
	}

	service.recordSuccessMetric(ctx, url, r.StatusCode(), "MakeGetRequest", start)
	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakePostRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	start := time.Now()
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	r, err := req.Post(url)

	if err != nil {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakePostRequest", err)
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakePostRequest", errors.New(http.StatusText(r.StatusCode())))
		return r.StatusCode(), r.Body(), err
	}

	service.recordSuccessMetric(ctx, url, r.StatusCode(), "MakePostRequest", start)
	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakePutRequest(ctx *gin.Context, url string, body interface{}, headers http.Header) (int, []byte, error) {
	start := time.Now()
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	r, err := req.Put(url)

	if err != nil {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakePutRequest", err)
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakePutRequest", errors.New(http.StatusText(r.StatusCode())))
		return r.StatusCode(), r.Body(), err
	}

	service.recordSuccessMetric(ctx, url, r.StatusCode(), "MakePutRequest", start)
	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakeDeleteRequest(ctx *gin.Context, url string, headers http.Header) (int, []byte, error) {
	start := time.Now()
	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Delete(url)

	if err != nil {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeDeleteRequest", err)
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeDeleteRequest", errors.New(http.StatusText(r.StatusCode())))
		return r.StatusCode(), r.Body(), err
	}

	service.recordSuccessMetric(ctx, url, r.StatusCode(), "MakeDeleteRequest", start)
	return r.StatusCode(), r.Body(), nil
}

func (service *restyService) MakeGetRequestWithConfig(ctx *gin.Context, url string, headers http.Header, config RequestConfig) (int, []byte, error) {
	start := time.Now()
	client := service.restyClient
	client.SetTimeout(config.Timeout)

	var r *resty.Response
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	r, err := req.Get(url)

	if err != nil {
		// Send metric to DD
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeGetRequestWithConfig", err)
		return r.StatusCode(), r.Body(), err
	}

	if r.StatusCode() != http.StatusOK {
		// Send metric to DD
		service.recordErrorMetric(ctx, url, r.StatusCode(), "MakeGetRequestWithConfig", errors.New(http.StatusText(r.StatusCode())))
		return r.StatusCode(), r.Body(), err
	}

	// Send metric to DD
	service.recordSuccessMetric(ctx, url, r.StatusCode(), "MakeGetRequestWithConfig", start)

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

func (service *restyService) recordSuccessMetric(c *gin.Context, url string, statusCode int, resource string, start time.Time) {
	// Metric
	tags := new(godog.Tags).
		Add("url", url).
		Add("status_code", fmt.Sprintf("%d", statusCode))

	godog.RecordSimpleMetric(
		fmt.Sprintf("resty.%s.%s.success", service.datadogMetricPrefix, resource),
		1,
		tags.ToArray()...)

	godog.RecordCompoundMetric(
		fmt.Sprintf("resty.%s.%s.elapsed_time", service.datadogMetricPrefix, resource),
		ElapsedSinceFloat(start),
		tags.ToArray()...)
}

func (service *restyService) recordErrorMetric(c *gin.Context, url string, statusCode int, resource string, err error) {
	// Metric
	tags := new(godog.Tags).
		Add("url", url).
		Add("status_code", fmt.Sprintf("%d", statusCode)).
		Add("error", err.Error())

	godog.RecordSimpleMetric(fmt.Sprintf("resty.%s.%s.error", service.datadogMetricPrefix, resource), 1, tags.ToArray()...)
}

// ElapsedSinceFloat returns elapsed time in ms as float64
func ElapsedSinceFloat(start time.Time) float64 {
	return float64(time.Since(start).Nanoseconds()) / 1000000.0
}
