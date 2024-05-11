package handler

import (
	"github.com/go-web/database/function"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/repository/member"
	messageRepo "github.com/go-web/pkg/repository/message"
	"github.com/go-web/pkg/service"
)

var AuthServiceImpl service.AuthService
var EventServiceImpl service.EventService
var EventMemberServiceImpl service.EventMemberService
var VenueServiceImpl service.VenueService
var UserServiceImpl service.UserService
var GroupServiceImpl service.GroupService
var MessageServiceImpl service.MessageService

func InititializeService() {
	dbFunction := function.NewDBFunction()

	eventRepository := repository.NewEventRepository(dbFunction)
	eventMemberRepository := member.NewEventMemberRepository(dbFunction)
	venueRepository := repository.NewVenueRepository(dbFunction)
	userRepository := repository.NewMemberRepository(dbFunction)
	groupRepository := repository.NewGroupRepository(dbFunction)
	messageRepository := messageRepo.NewMessageRepo(dbFunction)

	AuthServiceImpl = service.NewAuthService(userRepository)
	EventMemberServiceImpl = service.NewEventMemberService(eventMemberRepository)
	EventServiceImpl = service.NewEventService(eventRepository, EventMemberServiceImpl)
	VenueServiceImpl = service.NewVenueService(venueRepository)
	UserServiceImpl = service.NewUserService(userRepository)
	GroupServiceImpl = service.NewGroupService(groupRepository)
	MessageServiceImpl = service.NewMessageService(messageRepository)
}
