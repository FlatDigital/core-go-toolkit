package rest

// func Test_GetRequestBuilderWithTimeout(t *testing.T) {
// 	ass := assert.New(t)
// 	// given
// 	timeOut := time.Duration(5) * time.Minute
// 	header := http.Header{}
// 	header.Add("Content-Type", "multipart/form-data")
// 	ctx := gk.CreateTestContext()

// 	// when
// 	service := service{}
// 	value := service.getRequestBuilder(ctx, header, RequestConfig{Timeout: timeOut})

// 	// then
// 	ass.Equal(3, int(value.ContentType)) // 3 refers to type multipart/form-data
// 	ass.Equal(5*time.Minute, value.Timeout)
// }

// func Test_GetRequestBuilderWithTimeout_OtherHeaderGetJson(t *testing.T) {
// 	ass := assert.New(t)
// 	// given
// 	timeOut := time.Duration(5) * time.Minute
// 	header := http.Header{}
// 	header.Add("Content-Type", "somethingElse")
// 	ctx := gk.CreateTestContext()

// 	// when
// 	service := service{}
// 	value := service.getRequestBuilder(ctx, header, RequestConfig{Timeout: timeOut})

// 	// then
// 	// 0 refers to type json
// 	ass.Equal(0, int(value.ContentType))
// 	ass.Equal(5*time.Minute, value.Timeout)
// }

// func Test_Evaluate_Response(t *testing.T) {
// 	ass := assert.New(t)

// 	service := service{}

// 	// Response nil
// 	status, _, err := service.evaluateResponse(nil)
// 	ass.Equal(0, status)
// 	ass.Equal("response not received", err.Error())

// 	// Error with response
// 	response := &rest.Response{
// 		Response: &http.Response{
// 			Status:     "Custom Error",
// 			StatusCode: 500,
// 		},
// 		Err: errors.New("custom error"),
// 	}
// 	status, _, err = service.evaluateResponse(response)
// 	ass.Equal(500, status)
// 	ass.NotNil(err)
// 	ass.Equal("custom error", err.Error())

// 	// Error without response
// 	response = &rest.Response{
// 		Err: errors.New("custom error"),
// 	}
// 	status, _, err = service.evaluateResponse(response)
// 	ass.Equal(0, status)
// 	ass.NotNil(err)
// 	ass.Equal("custom error", err.Error())
// }

// func Test_GetRequestBuilder_Change_Only_Retry_Policy(t *testing.T) {
// 	ass := assert.New(t)

// 	// given
// 	header := http.Header{}
// 	header.Add("Content-Type", "somethingElse")
// 	ctx := gk.CreateTestContext()

// 	// when
// 	service := service{}
// 	srs := retry.NewSimpleRetryStrategy(3, 4)
// 	value := service.getRequestBuilder(ctx, header, RequestConfig{RetryStrategy: &srs})

// 	// then
// 	ass.Equal(3*time.Second, value.Timeout)
// 	ass.NotNil(value.RetryStrategy)
// }
