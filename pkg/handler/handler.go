package handler

import (
	"github.com/go-web/database/function"
	eventModel "github.com/go-web/pkg/model/event"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/repository/member"
	"github.com/go-web/pkg/service"
)

var AuthServiceImpl service.AuthService
var EventServiceImpl service.EventService
var EventMemberServiceImpl service.EventMemberService
var VenueServiceImpl service.VenueService
var UserServiceImpl service.UserService
var GroupServiceImpl service.GroupService

func InititializeService() {
	dbFunction := function.NewDBFunction()

	eventRepository := repository.NewEventRepository(dbFunction)
	eventMemberRepository := member.NewEventMemberRepository(dbFunction)
	venueRepository := repository.NewVenueRepository(dbFunction)
	userRepository := repository.NewMemberRepository(dbFunction)
	groupRepository := repository.NewGroupRepository(dbFunction)

	AuthServiceImpl = service.NewAuthService(userRepository)
	EventMemberServiceImpl = service.NewEventMemberService(eventMemberRepository)
	EventServiceImpl = service.NewEventService(eventRepository, EventMemberServiceImpl)
	VenueServiceImpl = service.NewVenueService(venueRepository)
	UserServiceImpl = service.NewUserService(userRepository)
	GroupServiceImpl = service.NewGroupService(groupRepository)
}

func checkRequest(requestEvent eventModel.Event) {

}
