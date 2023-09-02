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

func InititializeService() {
	dbFunction := function.NewDBFunction()

	eventRepository := repository.NewEventRepository(dbFunction)
	eventMemberRepository := member.NewEventMemberRepository(dbFunction)

	EventServiceImpl = service.NewEventService(eventRepository)
	EventMemberServiceImpl = service.NewEventMemberService(eventMemberRepository)
}

func checkRequest(requestEvent eventModel.Event) {

}
