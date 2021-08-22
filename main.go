package main

import (
	"log"
	"net/http"

	"github.com/go-web/endpoint"
	"github.com/go-web/pkg/http/routes"
	"github.com/go-web/server"
)

func main() {
	s := server.NewServer()
	s.AddHandler(endpoint.Endpoint{
		Path:    "/category",
		Method:  "POST",
		Handler: routes.CategoryHandler,
	})
	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
