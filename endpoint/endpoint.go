package endpoint

import "github.com/go-web/model"

type Handler func(i interface{}) (model.Response, error)

type Endpoints struct {
	endpoints []Endpoint
}

type Endpoint struct {
	Method  string
	Handler Handler
	Path    string
}

var endpoint Endpoints

func Init() {
	endpoints := Endpoints{}
	endpoints.endpoints = make([]Endpoint, 0)
}

func CreateEndpoint(e Endpoint) {
	endpoint.endpoints = append(endpoint.endpoints, e)
}
