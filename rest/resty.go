package rest

import (
	"errors"
	"fmt"
	"github.com/FlatDigital/core-go-toolkit/v2/godog"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/FlatDigital/core-go-toolkit/v2/core/flat"
	"github.com/go-resty/resty/v2"
)

type logType string

const (
	MakeGetRequest           string = "get_request"
	MakePostRequest          string = "post_request"
	MakePutRequest           string = "put_request"
	MakePatchRequest         string = "patch_request"
	MakeDeleteRequest        string = "delete_request"
	MakeGetRequestWithConfig string = "get_with_config_request"

	logError   logType = "error"
	logSuccess logType = "success"
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

func (service *restyService) MakeGetRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, http.Header, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Get(url)
	return service.evaluateResponse(url, response, MakeGetRequest, start, err)
}

func (service *restyService) MakePostRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, http.Header, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Post(url)
	return service.evaluateResponse(url, response, MakePostRequest, start, err)
}

func (service *restyService) MakePutRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, http.Header, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Put(url)
	return service.evaluateResponse(url, response, MakePutRequest, start, err)
}

func (service *restyService) MakePatchRequest(ctx *flat.Context, url string, body interface{}, headers http.Header) (int, []byte, http.Header, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)
	req.SetBody(body)

	response, err := req.Patch(url)
	return service.evaluateResponse(url, response, MakePutRequest, start, err)
}

func (service *restyService) MakeDeleteRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, http.Header, error) {
	start := time.Now()
	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Delete(url)
	return service.evaluateResponse(url, response, MakeDeleteRequest, start, err)
}

func (service *restyService) MakeGetRequestWithConfig(ctx *flat.Context, url string, headers http.Header, config RequestConfig) (int, []byte, http.Header, error) {
	start := time.Now()
	client := service.restyClient
	client.SetTimeout(config.Timeout)

	req := service.restyClient.R()
	req.SetHeaderMultiValues(headers)

	response, err := req.Get(url)
	return service.evaluateResponse(url, response, MakeGetRequestWithConfig, start, err)
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

func (service *restyService) evaluateResponse(url string, response *resty.Response, method string, start time.Time, err error) (int, []byte, http.Header, error) {

	if response == nil {
		service.logMetric(logError, url, response.StatusCode(), method, start)
		return 0, nil, http.Header{}, errResponseNotReceived
	}

	if err != nil {
		service.logMetric(logError, url, response.StatusCode(), method, start)
		return response.StatusCode(), response.Body(), response.Header(), err
	}

	if !(response.StatusCode() >= http.StatusOK && response.StatusCode() <= http.StatusIMUsed) {
		err = errors.New(response.Status())
		service.logMetric(logError, url, response.StatusCode(), method, start)
		return response.StatusCode(), response.Body(), response.Header(), err
	}

	service.logMetric(logSuccess, url, response.StatusCode(), method, start)
	return response.StatusCode(), response.Body(), response.Header(), nil
}

func getSanitizedPathFromUrl(rawUrl string) string {
	var path string
	parsedUrl, _ := url.Parse(rawUrl)
	if parsedUrl != nil {
		path = parsedUrl.Path
	}

	// ids removal
	re := regexp.MustCompile(`/[0-9]+`)
	path = re.ReplaceAllString(path, "/_")

	// emails removal
	re = regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`)
	path = re.ReplaceAllString(path, "_")

	return path
}

func (service *restyService) logMetric(logType logType, rawUrl string, statusCode int, action string, start time.Time) {
	// Metric
	tags := new(godog.Tags).
		Add("resource", getSanitizedPathFromUrl(rawUrl)).
		Add("status_code", fmt.Sprintf("%d", statusCode)).
		Add("action", action)
	godog.RecordSimpleMetric(
		fmt.Sprintf("application.%s.rest.service.%s", service.datadogMetricPrefix, logType),
		1,
		tags.ToArray()...,
	)

	godog.RecordCompoundMetric(
		fmt.Sprintf("application.%s.rest.service.elapsed_time", service.datadogMetricPrefix),
		elapsedSinceFloat(start),
		tags.ToArray()...)
}

// elapsedSinceFloat returns elapsed time in ms as float64
func elapsedSinceFloat(start time.Time) float64 {
	return float64(time.Since(start).Milliseconds())
}
