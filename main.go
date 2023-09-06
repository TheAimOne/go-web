package main

import (
	"log"
	"net/http"

	"github.com/go-web/database/connection"
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
	s.AddHandler(endpoint.Endpoint{
		Path:    "/venues",
		Method:  "GET",
		Handler: handler.GetVenueHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/venue",
		Method:  "POST",
		Handler: handler.CreateVenueHandler,
	})

	connection.InitDB()

	handler.InititializeService()

	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
