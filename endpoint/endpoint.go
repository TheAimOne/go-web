package endpoint

import (
	"net/http"

	"github.com/go-web/pkg/model"
)

type Handler func(i interface{}) (*model.Response, error)
type ResponseWriterHandler func(i interface{}, rw http.ResponseWriter) (*model.Response, error)

type GenericHandler interface {
	Handler | ResponseWriterHandler
}

type Endpoints struct {
	endpoints []Endpoint
}

type Endpoint struct {
	Method                string
	Handler               Handler
	ResponseWriterHandler ResponseWriterHandler
	Path                  string
}

var endpoint Endpoints

func Init() {
	endpoints := Endpoints{}
	endpoints.endpoints = make([]Endpoint, 0)
}

func CreateEndpoint(e Endpoint) {
	endpoint.endpoints = append(endpoint.endpoints, e)
}
