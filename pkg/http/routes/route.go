package routes

import (
	"time"

	"github.com/go-web/endpoint"
	"github.com/go-web/model"
)

func MakeCategoryEnpoint() endpoint.Endpoint {

	e := endpoint.Endpoint{}

	e.Path = "/category"
	e.Method = "POST"
	e.Handler = CategoryHandler

	return e
}

func CategoryHandler(request interface{}) (model.Response, error) {

	m := model.Message{
		Name: "Sandeep",
		Body: "jupp",
		Time: (time.Now().Unix()),
	}

	response := model.Response{
		Body: m,
	}

	return response.Json(), nil
}
