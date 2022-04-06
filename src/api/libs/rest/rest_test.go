package rest_test

// const (
// 	testURL        = "/test"
// 	serviceTimeout = 1500 * time.Millisecond
// 	requestTimeout = 500 * time.Millisecond
// )

// func Test_NewService(t *testing.T) {
// 	// given
// 	ass := assert.New(t)

// 	// when
// 	service := rest.NewService("test_base_url", 3)

// 	// then
// 	ass.NotNil(service)
// }

// func Test_MakeGetRequest(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewService("https://test_base_url.com", 3)

// 	// when
// 	mockRequest(http.MethodGet, 200)
// 	code, payload, err := service.MakeGetRequest(nil, testURL, http.Header{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodGet, 404)
// 	code, payload, err = service.MakeGetRequest(nil, testURL, http.Header{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakeGetRequestWithContext(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewService("https://test_base_url.com", 3)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodGet, 200)
// 	code, payload, err := service.MakeGetRequest(ctx, testURL, http.Header{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodGet, 404)
// 	code, payload, err = service.MakeGetRequest(ctx, testURL, http.Header{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakePostRequest(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewService("https://test_base_url.com", 3)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodPost, 200)
// 	code, payload, err := service.MakePostRequest(ctx, testURL, "", http.Header{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPost, 404)
// 	code, payload, err = service.MakePostRequest(ctx, testURL, "", http.Header{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakePutRequest(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewService("https://test_base_url.com", 3)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodPut, 200)
// 	code, payload, err := service.MakePutRequest(ctx, testURL, "", http.Header{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPut, 404)
// 	code, payload, err = service.MakePutRequest(ctx, testURL, "", http.Header{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakeDeleteRequest(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewService("https://test_base_url.com", 3)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodDelete, 200)
// 	code, payload, err := service.MakeDeleteRequest(ctx, testURL, http.Header{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodDelete, 404)
// 	code, payload, err = service.MakeDeleteRequest(ctx, testURL, http.Header{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_NewService_Timeout(t *testing.T) {
// 	// given
// 	ass := assert.New(t)

// 	// when
// 	service := rest.NewServiceWithTimeout("test_base_url", 3, serviceTimeout)

// 	// then
// 	ass.NotNil(service)
// }

// func Test_MakeGetRequest_Timeout(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewServiceWithTimeout("https://test_base_url.com", 3, serviceTimeout)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodGet, 200)
// 	code, payload, err := service.MakeGetRequestWithTimeout(ctx, testURL, http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodGet, 404)
// 	code, payload, err = service.MakeGetRequestWithTimeout(ctx, testURL, http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakePostRequest_Timeout(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewServiceWithTimeout("https://test_base_url.com", 3, serviceTimeout)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodPost, 200)
// 	code, payload, err := service.MakePostRequestWithTimeout(ctx, testURL, "", http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPost, 404)
// 	code, payload, err = service.MakePostRequestWithTimeout(ctx, testURL, "", http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakePutRequest_Timeout(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewServiceWithTimeout("https://test_base_url.com", 3, serviceTimeout)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodPut, 200)
// 	code, payload, err := service.MakePutRequestWithTimeout(ctx, testURL, "", http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPut, 404)
// 	code, payload, err = service.MakePutRequestWithTimeout(ctx, testURL, "", http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakeDeleteRequest_Timeout(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	service := rest.NewServiceWithTimeout("https://test_base_url.com", 3, serviceTimeout)
// 	ctx := gk.CreateTestContext()

// 	// when
// 	mockRequest(http.MethodDelete, 200)
// 	code, payload, err := service.MakeDeleteRequestWithTimeout(ctx, testURL, http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodDelete, 404)
// 	code, payload, err = service.MakeDeleteRequestWithTimeout(ctx, testURL, http.Header{}, requestTimeout)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func Test_MakeGetRequestWithConfig(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	ctx := gk.CreateTestContext()
// 	sConfig := rest.ServiceConfig{
// 		BaseURL:             "https://test_base_url.com",
// 		MaxIdleConnsPerHost: 3,
// 	}
// 	service := rest.NewRestService(sConfig)

// 	// when
// 	mockRequest(http.MethodGet, 200)
// 	code, payload, err := service.MakeGetRequest(ctx, testURL, nil)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodGet, 202)
// 	code, payload, err = service.MakeGetRequestWithConfig(ctx, testURL, http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(202, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodGet, 404)
// 	code, payload, err = service.MakeGetRequest(ctx, testURL, nil)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)

// 	mockRequest(http.MethodGet, 400)
// 	code, payload, err = service.MakeGetRequestWithConfig(ctx, testURL, http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(400, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "400 Bad Request")
// 	ass.NotNil(payload)
// }

// func Test_MakePostRequestWithConfig(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	ctx := gk.CreateTestContext()
// 	sConfig := rest.ServiceConfig{
// 		BaseURL:             "https://test_base_url.com",
// 		MaxIdleConnsPerHost: 3,
// 	}
// 	service := rest.NewRestService(sConfig)

// 	// when
// 	mockRequest(http.MethodPost, 200)
// 	code, payload, err := service.MakePostRequest(ctx, testURL, "", nil)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPost, 202)
// 	code, payload, err = service.MakePostRequestWithConfig(ctx, testURL, "", http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(202, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPost, 404)
// 	code, payload, err = service.MakePostRequest(ctx, testURL, "", nil)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPost, 400)
// 	code, payload, err = service.MakePostRequestWithConfig(ctx, testURL, "", http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(400, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "400 Bad Request")
// 	ass.NotNil(payload)
// }

// func Test_MakePutRequestWithConfig(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	ctx := gk.CreateTestContext()
// 	sConfig := rest.ServiceConfig{
// 		BaseURL:             "https://test_base_url.com",
// 		MaxIdleConnsPerHost: 3,
// 	}
// 	service := rest.NewRestService(sConfig)

// 	// when
// 	mockRequest(http.MethodPut, 200)
// 	code, payload, err := service.MakePutRequest(ctx, testURL, "", nil)

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPut, 202)
// 	code, payload, err = service.MakePutRequestWithConfig(ctx, testURL, "", http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(202, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPut, 404)
// 	code, payload, err = service.MakePutRequest(ctx, testURL, "", nil)

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodPut, 400)
// 	code, payload, err = service.MakePutRequestWithConfig(ctx, testURL, "", http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(400, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "400 Bad Request")
// 	ass.NotNil(payload)
// }

// func Test_MakeDeleteRequestWithConfig(t *testing.T) {
// 	// given
// 	restful.StartMockupServer()
// 	ass := assert.New(t)
// 	ctx := gk.CreateTestContext()
// 	sConfig := rest.ServiceConfig{
// 		BaseURL:             "https://test_base_url.com",
// 		MaxIdleConnsPerHost: 3,
// 	}
// 	service := rest.NewRestService(sConfig)

// 	// when
// 	mockRequest(http.MethodDelete, 200)
// 	code, payload, err := service.MakeDeleteRequestWithConfig(ctx, testURL, http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(200, code)
// 	ass.Nil(err)
// 	ass.NotNil(payload)

// 	// when
// 	mockRequest(http.MethodDelete, 404)
// 	code, payload, err = service.MakeDeleteRequestWithConfig(ctx, testURL, http.Header{}, rest.RequestConfig{})

// 	// then
// 	ass.Equal(404, code)
// 	ass.NotNil(err)
// 	ass.EqualError(err, "404 Not Found")
// 	ass.NotNil(payload)
// }

// func mockRequest(method string, statusCode int) {
// 	// status code 200
// 	httpMethodResponseMock(method,
// 		"https://test_base_url.com/test",
// 		"",
// 		`{"some_response_value": true}`,
// 		statusCode)
// }

// // HTTPMethodResponseMock mock method for rest calls
// func httpMethodResponseMock(method string, mockURL, reqBody, respBody string, responseCode int) {
// 	header := http.Header{}
// 	httpMethod := method

// 	mock := restful.Mock{
// 		URL:          mockURL,
// 		HTTPMethod:   httpMethod,
// 		ReqBody:      reqBody,
// 		RespHTTPCode: responseCode,
// 		RespHeaders:  header,
// 		RespBody:     respBody,
// 	}
// 	restful.AddMockups(&mock)
// }
