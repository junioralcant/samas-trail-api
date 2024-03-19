package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeHttpRequest(c *gin.Context, u url.Values, body interface{}, params gin.Params) {

	jsonbytes, _ := json.Marshal(body)

	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = params
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
	c.Request.URL.RawQuery = u.Encode()
}
