package service

import (
	"encoding/json"

	"github.com/go-web/pkg/constants"
	eventModel "github.com/go-web/pkg/model/event"
	"github.com/go-web/pkg/repository"
	uuid "github.com/satori/go.uuid"
)

type eventImpl struct {
	EventRepository repository.EventRepository
}

func (e *eventImpl) CreateEvent(requestEvent *eventModel.Event) (*eventModel.EventResponse, error) {

	requestEvent.EventId = uuid.NewV4()

	paramsString, err := json.Marshal(requestEvent.Params)
	if err != nil {
		return nil, constants.ErrorParsingParams
	}

	requestEvent.Params = string(paramsString)
	requestEvent.Status = constants.CREATED

	err = e.EventRepository.CreateEvent(requestEvent)

	if err != nil {
		return nil, constants.ErrorCreatingEvent
	}

	return &eventModel.EventResponse{
		EventId: requestEvent.EventId,
		Status:  requestEvent.Status,
	}, nil
}

func (e *eventImpl) GetEventsByGroupId(eventRequest *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error) {

	groupId := eventRequest.GroupId

	result, err := e.EventRepository.GetEventsByGroupId(groupId)

	if err != nil {
		return nil, err
	}

	response := &eventModel.GetEventResponse{}

	response.Data = result

	return response, nil
}
