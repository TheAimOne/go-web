package handler

import (
	"github.com/go-web/database/function"
	eventModel "github.com/go-web/pkg/model/event"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/repository/member"
	"github.com/go-web/pkg/service"
)

var EventServiceImpl service.EventService
var EventMemberServiceImpl service.EventMemberService
var VenueServiceImpl service.VenueService
var UserServiceImpl service.UserService

func InititializeService() {
	dbFunction := function.NewDBFunction()

	eventRepository := repository.NewEventRepository(dbFunction)
	eventMemberRepository := member.NewEventMemberRepository(dbFunction)
	venueRepository := repository.NewVenueRepository(dbFunction)
	userRepository := repository.NewMemberRepository(dbFunction)

	EventServiceImpl = service.NewEventService(eventRepository)
	EventMemberServiceImpl = service.NewEventMemberService(eventMemberRepository)
	VenueServiceImpl = service.NewVenueService(venueRepository)
	UserServiceImpl = service.NewUserService(userRepository)
}

func checkRequest(requestEvent eventModel.Event) {

}
