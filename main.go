package main

import (
	"log"
	"net/http"

	"github.com/go-web/database/connection"
	"github.com/go-web/database/function"
	"github.com/go-web/endpoint"
	"github.com/go-web/pkg/handler"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/service"
	"github.com/go-web/server"
)

func main() {
	s := server.NewServer()
	s.AddHandler(endpoint.Endpoint{
		Path:    "/events",
		Method:  "POST",
		Handler: handler.EventHandler,
	})

	connection.InitDB()
	dbFunction := function.NewDBFunction()
	eventRepository := repository.NewEventRepository(dbFunction)

	service := service.NewEventService(eventRepository)

	handler.InititializeService(service)

	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
