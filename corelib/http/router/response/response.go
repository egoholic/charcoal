package response

import "bytes"

type Response struct {
	body       *bytes.Buffer
	headers    map[string][]string
	statusCode int
}

func New() *Response {
	return &Response{bytes.NewBuffer([]byte{}), map[string][]string{}, 0}
}

func (r *Response) SetHeader(hn string, hv ...string) {
	r.headers[hn] = hv
}

func (r *Response) Write(v []byte) (int, error) {
	return r.body.Write(v)
}

func (r *Response) SetStatus(s int) {
	r.statusCode = s
}

func (r *Response) Status() int {
	return r.statusCode
}

func (r *Response) Body() *bytes.Buffer {
	return r.body
}

func (r *Response) Headers() map[string][]string {
	return r.headers
}
