package helper

import (
	"net/http"

	"github.com/egoholic/charcoal/corelib/http/router/params"
)

func TestHandlerFunc(w http.ResponseWriter, r *http.Request, p *params.Params) {
	w.WriteHeader(200)
	w.Header().Add("TEST-HEADER", p.Param("header")[0])
	w.Write([]byte("hello!"))
	var body = []byte{}
	r.Body.Read(body)
	w.Write(body)
}
