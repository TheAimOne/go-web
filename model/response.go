package model

import (
	"encoding/json"
	"net/http"

	"github.com/go-web/constants"
)

type Response struct {
	Body        interface{}
	ContentType string
	Status      int
}

func (r Response) Json() Response {
	r.ContentType = constants.HEADER_APPLICATION_JSON
	r.Status = http.StatusOK
	return r
}

func (r Response) Write(rw http.ResponseWriter) {
	rw.Header().Set(constants.HEADER_CONTENT_TYPE_KEY, constants.HEADER_APPLICATION_JSON)

	b, _ := json.Marshal(r.Body)

	rw.Write([]byte(b))
}
