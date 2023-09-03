package main

import (
	"log"
	"net/http"

	"github.com/go-web/database/connection"xs
	"github.com/go-web/endpoint"
	"github.com/go-web/pkg/handler"
	"github.com/go-web/server"
)

func main() {
	s := server.NewServer()
	s.AddHandler(endpoint.Endpoint{
		Path:    "/events",
		Method:  "POST",
		Handler: handler.CreateEventHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/groups/events",
		Method:  "GET",
		Handler: handler.GetEventByGroupIdHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/events/members",
		Method:  "POST",
		Handler: handler.CreateEventMemberHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/events/members",
		Method:  "GET",
		Handler: handler.GetEventMembers,
	})

	connection.InitDB()

	handler.InititializeService()

	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
