package service

import (
	model "github.com/go-web/model"
	"github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
	memberModel "github.com/go-web/pkg/model/member"
	venueModel "github.com/go-web/pkg/model/venue"
	"github.com/go-web/pkg/repository"
	"github.com/go-web/pkg/repository/member"
)

type EventService interface {
	CreateEvent(eventRequest *eventModel.Event) (*eventModel.EventResponse, error)
	GetEventsByGroupId(eventRequest *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventImpl{
		eventRepository,
	}
}

type EventMemberService interface {
	CreateEventMember(eventRequest *eventModel.AddMemberToEventRequest) (*eventModel.AddMemberToEventResponse, error)
	GetEventMembers(request *memberModel.GetEventMembersRequest) (*memberModel.GetEventMembersResponse, error)
}

func NewEventMemberService(eventRepository member.EventMemberRepository) EventMemberService {
	return &eventMemberImpl{
		eventRepository,
	}
}

type VenueService interface {
	CreateVenue(venueRequest *venueModel.Venue) (*venueModel.Venue, error)
	GetVenues(request *model.Filter) (*venueModel.GetVenueResponse, error)
}

func NewVenueService() VenueService {
	return nil
}
