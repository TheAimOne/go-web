package service

import (
	eventModel "github.com/go-web/pkg/model/event"
	"github.com/go-web/pkg/repository"
)

type EventService interface {
	CreateEvent(eventRequest *eventModel.Event) (*eventModel.EventResponse, error)
}

func NewEventService(eventRepository repository.EventRepository) EventService {
	return &eventImpl{
		eventRepository,
	}
}
