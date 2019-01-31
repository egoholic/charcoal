package response

// copy of net/http's http.Header
type Header map[string][]string

// copy of net/http's http.ResponseWriter
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

type Response struct {
	responseWriter ResponseWriter
}

func New(w ResponseWriter) *Response {
	return &Response{w}
}

func (r *Response) SetHeader(hn string, hv ...string) {
	r.responseWriter.Header()[hn] = hv
}

func (r *Response) WriteBody(v []byte) (int, error) {
	return r.responseWriter.Write(v)
}

func (r *Response) SetStatus(s int) {
	r.responseWriter.WriteHeader(s)
}
