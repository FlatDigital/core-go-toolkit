package rest

import (
	"crypto/md5"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/FlatDigital/core-go-toolkit/core/flat"
)

// Mock struct for Rest mock
type Mock struct {
	mux sync.Mutex

	makeGetRequestMockStack    map[hash][]outputForMakeGetRequest
	makePostRequestMockStack   map[hash][]outputForMakePostRequest
	makePutRequestMockStack    map[hash][]outputForMakePutRequest
	makeDeleteRequestMockStack map[hash][]outputForMakeDeleteRequest

	makeGetRequestWithConfigMockStack    map[hash][]outputForMakeGetRequestWithConfig
	makePostRequestWithConfigMockStack   map[hash][]outputForMakePostRequestWithConfig
	makePutRequestWithConfigMockStack    map[hash][]outputForMakePutRequestWithConfig
	makeDeleteRequestWithConfigMockStack map[hash][]outputForMakeDeleteRequestWithConfig

	makeGetRequestWithTimeoutMockStack    map[hash][]outputForMakeGetRequestWithTimeout
	makePostRequestWithTimeoutMockStack   map[hash][]outputForMakePostRequestWithTimeout
	makePutRequestWithTimeoutMockStack    map[hash][]outputForMakePutRequestWithTimeout
	makeDeleteRequestWithTimeoutMockStack map[hash][]outputForMakeDeleteRequestWithTimeout
}

// NewMock Rest Mock
func NewMock() *Mock {
	return &Mock{
		makeGetRequestMockStack:    map[hash][]outputForMakeGetRequest{},
		makePostRequestMockStack:   map[hash][]outputForMakePostRequest{},
		makePutRequestMockStack:    map[hash][]outputForMakePutRequest{},
		makeDeleteRequestMockStack: map[hash][]outputForMakeDeleteRequest{},

		makeGetRequestWithConfigMockStack:    map[hash][]outputForMakeGetRequestWithConfig{},
		makePostRequestWithConfigMockStack:   map[hash][]outputForMakePostRequestWithConfig{},
		makePutRequestWithConfigMockStack:    map[hash][]outputForMakePutRequestWithConfig{},
		makeDeleteRequestWithConfigMockStack: map[hash][]outputForMakeDeleteRequestWithConfig{},

		makeGetRequestWithTimeoutMockStack:    map[hash][]outputForMakeGetRequestWithTimeout{},
		makePostRequestWithTimeoutMockStack:   map[hash][]outputForMakePostRequestWithTimeout{},
		makePutRequestWithTimeoutMockStack:    map[hash][]outputForMakePutRequestWithTimeout{},
		makeDeleteRequestWithTimeoutMockStack: map[hash][]outputForMakeDeleteRequestWithTimeout{},
	}
}

type inputForMakeGetRequest struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
}

type outputForMakeGetRequest struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeGetRequest patch for MakeGetRequest function
func (mock *Mock) PatchMakeGetRequest(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequest{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
	}

	inputHash := toHash(input)

	output := outputForMakeGetRequest{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeGetRequestMockStack[inputHash] = append(mock.makeGetRequestMockStack[inputHash], output)
}

// MakeGetRequest mock for MakeGetRequest function
func (mock *Mock) MakeGetRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequest{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeGetRequestMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeGetRequest")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeGetRequestMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakePostRequest struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
}

type outputForMakePostRequest struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePostRequest patch for MakePostRequest function
func (mock *Mock) PatchMakePostRequest(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePostRequest{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
	}

	inputHash := toHash(input)

	output := outputForMakePostRequest{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePostRequestMockStack[inputHash] = append(mock.makePostRequestMockStack[inputHash], output)
}

// MakePostRequest mock for MakePostRequest function
func (mock *Mock) MakePostRequest(ctx *flat.Context, url string, body interface{},
	headers http.Header) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	cleanMultipartHeaders(headers)

	input := inputForMakePostRequest{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePostRequestMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePostRequest")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePostRequestMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakePutRequest struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
}

type outputForMakePutRequest struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePutRequest patch for MakePutRequest function
func (mock *Mock) PatchMakePutRequest(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequest{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
	}

	inputHash := toHash(input)

	output := outputForMakePutRequest{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePutRequestMockStack[inputHash] = append(mock.makePutRequestMockStack[inputHash], output)
}

// MakePutRequest mock for MakePutRequest function
func (mock *Mock) MakePutRequest(ctx *flat.Context, url string, body interface{},
	headers http.Header) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequest{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePutRequestMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePutRequest")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePutRequestMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakeDeleteRequest struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
}

type outputForMakeDeleteRequest struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeDeleteRequest patch for MakeDeleteRequest function
func (mock *Mock) PatchMakeDeleteRequest(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequest{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
	}

	inputHash := toHash(input)

	output := outputForMakeDeleteRequest{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeDeleteRequestMockStack[inputHash] = append(mock.makeDeleteRequestMockStack[inputHash], output)
}

// MakeDeleteRequest mock for MakeDeleteRequest function
func (mock *Mock) MakeDeleteRequest(ctx *flat.Context, url string, headers http.Header) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequest{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeDeleteRequestMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeDeleteRequest")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeDeleteRequestMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

// Config

type inputForMakeGetRequestWithConfig struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
	InputConfig  RequestConfig
}

type outputForMakeGetRequestWithConfig struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeGetRequestWithConfig patch for MakeGetRequestWithConfig function
func (mock *Mock) PatchMakeGetRequestWithConfig(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	inputConfig RequestConfig, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequestWithConfig{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
		InputConfig:  inputConfig,
	}

	inputHash := toHash(input)

	output := outputForMakeGetRequestWithConfig{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeGetRequestWithConfigMockStack[inputHash] =
		append(mock.makeGetRequestWithConfigMockStack[inputHash], output)
}

// MakeGetRequestWithConfig mock for MakeGetRequestWithConfig function
func (mock *Mock) MakeGetRequestWithConfig(ctx *flat.Context, url string, headers http.Header,
	config RequestConfig) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequestWithConfig{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
		InputConfig:  config,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeGetRequestWithConfigMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeGetRequestWithConfig")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeGetRequestWithConfigMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakePostRequestWithConfig struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
	InputConfig  RequestConfig
}

type outputForMakePostRequestWithConfig struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePostRequestWithConfig patch for MakePostRequestWithConfig function
func (mock *Mock) PatchMakePostRequestWithConfig(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, inputConfig RequestConfig, outputStatusCode int, outputResponse []byte,
	outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePostRequestWithConfig{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
		InputConfig:  inputConfig,
	}

	inputHash := toHash(input)

	output := outputForMakePostRequestWithConfig{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePostRequestWithConfigMockStack[inputHash] =
		append(mock.makePostRequestWithConfigMockStack[inputHash], output)
}

// MakePostRequestWithConfig mock for MakePostRequestWithConfig function
func (mock *Mock) MakePostRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header,
	config RequestConfig) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	cleanMultipartHeaders(headers)

	input := inputForMakePostRequestWithConfig{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
		InputConfig:  config,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePostRequestWithConfigMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePostRequestWithConfig")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePostRequestWithConfigMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakePutRequestWithConfig struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
	InputConfig  RequestConfig
}

type outputForMakePutRequestWithConfig struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePutRequestWithConfig patch for MakePutRequestWithConfig function
func (mock *Mock) PatchMakePutRequestWithConfig(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, inputConfig RequestConfig, outputStatusCode int,
	outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequestWithConfig{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
		InputConfig:  inputConfig,
	}

	inputHash := toHash(input)

	output := outputForMakePutRequestWithConfig{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePutRequestWithConfigMockStack[inputHash] =
		append(mock.makePutRequestWithConfigMockStack[inputHash], output)
}

// MakePutRequestWithConfig mock for MakePutRequestWithConfig function
func (mock *Mock) MakePutRequestWithConfig(ctx *flat.Context, url string, body interface{}, headers http.Header,
	config RequestConfig) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequestWithConfig{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
		InputConfig:  config,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePutRequestWithConfigMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePutRequestWithConfig")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePutRequestWithConfigMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakeDeleteRequestWithConfig struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
	InputConfig  RequestConfig
}

type outputForMakeDeleteRequestWithConfig struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeDeleteRequestWithConfig patch for MakeDeleteRequestWithConfig function
func (mock *Mock) PatchMakeDeleteRequestWithConfig(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	inputConfig RequestConfig, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequestWithConfig{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
		InputConfig:  inputConfig,
	}

	inputHash := toHash(input)

	output := outputForMakeDeleteRequestWithConfig{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeDeleteRequestWithConfigMockStack[inputHash] =
		append(mock.makeDeleteRequestWithConfigMockStack[inputHash], output)
}

// MakeDeleteRequestWithConfig mock for MakeDeleteRequestWithConfig function
func (mock *Mock) MakeDeleteRequestWithConfig(ctx *flat.Context, url string, headers http.Header,
	config RequestConfig) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequestWithConfig{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
		InputConfig:  config,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeDeleteRequestWithConfigMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeDeleteRequestWithConfig")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeDeleteRequestWithConfigMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

// Timeout

type inputForMakeGetRequestWithTimeout struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
	InputTimeout time.Duration
}

type outputForMakeGetRequestWithTimeout struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeGetRequestWithTimeout patch for MakeGetRequestWithTimeout function
func (mock *Mock) PatchMakeGetRequestWithTimeout(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	inputTimeout time.Duration, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequestWithTimeout{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
		InputTimeout: inputTimeout,
	}

	inputHash := toHash(input)

	output := outputForMakeGetRequestWithTimeout{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeGetRequestWithTimeoutMockStack[inputHash] =
		append(mock.makeGetRequestWithTimeoutMockStack[inputHash], output)
}

// MakeGetRequestWithTimeout mock for MakeGetRequestWithTimeout function
func (mock *Mock) MakeGetRequestWithTimeout(ctx *flat.Context, url string, headers http.Header,
	timeout time.Duration) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeGetRequestWithTimeout{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
		InputTimeout: timeout,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeGetRequestWithTimeoutMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeGetRequestWithTimeout")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeGetRequestWithTimeoutMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakePostRequestWithTimeout struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
	InputTimeout time.Duration
}

type outputForMakePostRequestWithTimeout struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePostRequestWithTimeout patch for MakePostRequestWithTimeout function
func (mock *Mock) PatchMakePostRequestWithTimeout(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, inputTimeout time.Duration, outputStatusCode int, outputResponse []byte,
	outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePostRequestWithTimeout{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
		InputTimeout: inputTimeout,
	}

	inputHash := toHash(input)

	output := outputForMakePostRequestWithTimeout{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePostRequestWithTimeoutMockStack[inputHash] =
		append(mock.makePostRequestWithTimeoutMockStack[inputHash], output)
}

// MakePostRequestWithTimeout mock for MakePostRequestWithTimeout function
func (mock *Mock) MakePostRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header,
	timeout time.Duration) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	cleanMultipartHeaders(headers)

	input := inputForMakePostRequestWithTimeout{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
		InputTimeout: timeout,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePostRequestWithTimeoutMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePostRequestWithTimeout")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePostRequestWithTimeoutMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

func cleanMultipartHeaders(headers http.Header) {
	// For test purposes, remove boundary if content type is multipart
	contentTypeKey := headers.Get("Content-Type")
	if len(contentTypeKey) >= 19 && contentTypeKey[:19] == "multipart/form-data" {
		headers.Set("Content-Type", "multipart/form-data")
	}
}

type inputForMakePutRequestWithTimeout struct {
	InputCTX     *flat.Context
	InputURL     string
	InputBody    interface{}
	InputHeaders http.Header
	InputTimeout time.Duration
}

type outputForMakePutRequestWithTimeout struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakePutRequestWithTimeout patch for MakePutRequestWithTimeout function
func (mock *Mock) PatchMakePutRequestWithTimeout(inputCTX *flat.Context, inputURL string, inputBody interface{},
	inputHeaders http.Header, inputTimeout time.Duration, outputStatusCode int, outputResponse []byte,
	outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequestWithTimeout{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputBody:    inputBody,
		InputHeaders: inputHeaders,
		InputTimeout: inputTimeout,
	}

	inputHash := toHash(input)

	output := outputForMakePutRequestWithTimeout{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makePutRequestWithTimeoutMockStack[inputHash] =
		append(mock.makePutRequestWithTimeoutMockStack[inputHash], output)
}

// MakePutRequestWithTimeout mock for MakePutRequestWithTimeout function
func (mock *Mock) MakePutRequestWithTimeout(ctx *flat.Context, url string, body interface{}, headers http.Header,
	timeout time.Duration) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakePutRequestWithTimeout{
		InputCTX:     ctx,
		InputURL:     url,
		InputBody:    body,
		InputHeaders: headers,
		InputTimeout: timeout,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makePutRequestWithTimeoutMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakePutRequestWithTimeout")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makePutRequestWithTimeoutMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type inputForMakeDeleteRequestWithTimeout struct {
	InputCTX     *flat.Context
	InputURL     string
	InputHeaders http.Header
	InputTimeout time.Duration
}

type outputForMakeDeleteRequestWithTimeout struct {
	OutputStatusCode int
	OutputResponse   []byte
	OutputError      error
}

// PatchMakeDeleteRequestWithTimeout patch for MakeDeleteRequestWithTimeout function
func (mock *Mock) PatchMakeDeleteRequestWithTimeout(inputCTX *flat.Context, inputURL string, inputHeaders http.Header,
	inputTimeout time.Duration, outputStatusCode int, outputResponse []byte, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequestWithTimeout{
		InputCTX:     inputCTX,
		InputURL:     inputURL,
		InputHeaders: inputHeaders,
		InputTimeout: inputTimeout,
	}

	inputHash := toHash(input)

	output := outputForMakeDeleteRequestWithTimeout{
		OutputStatusCode: outputStatusCode,
		OutputResponse:   outputResponse,
		OutputError:      outputError,
	}

	mock.makeDeleteRequestWithTimeoutMockStack[inputHash] =
		append(mock.makeDeleteRequestWithTimeoutMockStack[inputHash], output)
}

// MakeDeleteRequestWithTimeout mock for MakeDeleteRequestWithTimeout function
func (mock *Mock) MakeDeleteRequestWithTimeout(ctx *flat.Context, url string, headers http.Header,
	timeout time.Duration) (int, []byte, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := inputForMakeDeleteRequestWithTimeout{
		InputCTX:     ctx,
		InputURL:     url,
		InputHeaders: headers,
		InputTimeout: timeout,
	}

	inputHash := toHash(input)
	arrOutput, exists := mock.makeDeleteRequestWithTimeoutMockStack[inputHash]

	if !exists || len(arrOutput) == 0 {
		panic("Mock not available for MakeDeleteRequestWithTimeout")
	}

	output := arrOutput[0]
	arrOutput = arrOutput[1:]

	mock.makeDeleteRequestWithTimeoutMockStack[inputHash] = arrOutput
	return output.OutputStatusCode, output.OutputResponse, output.OutputError
}

type hash [16]byte

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}
