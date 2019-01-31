package helper

import (
	"fmt"
	"io"

	"github.com/egoholic/charcoal/corelib/http/router/params"
	"github.com/egoholic/charcoal/corelib/http/router/response"
)

func DumbHandler(p *params.Params, r *response.Response) {
	r.SetStatus(200)
	hdr := fmt.Sprintf("testHeader-%s", p.Param("header"))
	val := fmt.Sprintf("testValue-%s", p.Param("value"))
	r.SetHeader(hdr, val)
	body := p.Param("body")
	if len(body) > 0 {
		r.WriteBody([]byte(body[0]))
	} else {
		r.WriteBody([]byte("TestBody"))
	}
}

type TestReader struct {
	body []byte
}

func (r *TestReader) Read(p []byte) (int, error) {
	maxLen := len(p)
	dataLen := len(r.body)
	if maxLen > dataLen {
		maxLen = dataLen
	}

	for i := 0; i < maxLen; i++ {
		p[i] = r.body[i]
	}

	return maxLen, nil
}

type TestRequest struct {
	method string
	url    string
	body   io.Reader
}

type TestResponse struct {
	body       []byte
	statusCode int
	header     map[string][]string
}

func NewTestResponse() *TestResponse {
	header := response.Header{}
	return &TestResponse{[]byte{}, 000, header}
}

func NewTestRequest(method, url string) *TestRequest {
	reader := &TestReader{[]byte{}}
	return &TestRequest{method, url, reader}
}

func (r *TestResponse) Body() []byte {
	return r.body
}

func (r *TestResponse) StatusCode() int {
	return r.statusCode
}

func (r *TestResponse) Header() response.Header {
	return r.header
}

func (r *TestResponse) Write(v []byte) (int, error) {
	var i = 0

	for i, b := range v {
		r.body[i] = b
	}

	return i, nil
}

func (r *TestResponse) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}
