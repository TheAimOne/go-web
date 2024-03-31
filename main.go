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

	// Events
	s.AddHandler(endpoint.Endpoint{
		Path:    "/events",
		Method:  "POST",
		Handler: handler.CreateEventHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/group/events",
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
		Path:    "/events/search",
		Method:  "POST",
		Handler: handler.SearchEventHandler,
	})

	// Venues
	s.AddHandler(endpoint.Endpoint{
		Path:    "/venues",
		Method:  "POST",
		Handler: handler.GetVenueHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/venue",
		Method:  "POST",
		Handler: handler.CreateVenueHandler,
	})

	// User
	s.AddHandler(endpoint.Endpoint{
		Path:    "/user",
		Method:  "POST",
		Handler: handler.CreateUserHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/user",
		Method:  "GET",
		Handler: handler.GetUserByIdHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/users",
		Method:  "GET",
		Handler: handler.GetUsersHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/users/search",
		Method:  "POST",
		Handler: handler.SearchUserHandler,
	})

	// Group
	s.AddHandler(endpoint.Endpoint{
		Path:    "/group",
		Method:  "POST",
		Handler: handler.CreateGroupWithMembershandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/group",
		Method:  "GET",
		Handler: handler.GetGroupById,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/group/members",
		Method:  "POST",
		Handler: handler.AddMembersToGroupHandler,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/group/members",
		Method:  "GET",
		Handler: handler.GetMembersByGroupId,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/member/group",
		Method:  "GET",
		Handler: handler.GetGroupsByMemberId,
	})
	s.AddHandler(endpoint.Endpoint{
		Path:    "/groups",
		Method:  "GET",
		Handler: handler.GetGroups,
	})

	// Authentication
	s.AddHandler(endpoint.Endpoint{
		Path:    "/user/authenticate",
		Method:  "POST",
		Handler: handler.CreateAuthenticationHandler,
	})

	connection.InitDB()

	handler.InititializeService()

	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
