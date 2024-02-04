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
	EventMemService EventMemberService
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

func GetCountOfEventParticipants(e *eventImpl, result []*eventModel.Event) (map[string]int64, error) {
	var ids []string
	for _, event := range result {
		ids = append(ids, event.EventId.String())
	}
	countResult, err := e.EventMemService.CountMembersByEventId(ids)
	if err != nil {
		return nil, err
	}

	eventIdCountMap := make(map[string]int64)
	for _, result := range countResult {
		eventIdCountMap[result.EventId.String()] = result.Count
	}

	return eventIdCountMap, nil
}

func (e *eventImpl) GetEventsByGroupId(eventRequest *eventModel.GetEventRequest) (*eventModel.GetEventResponse, error) {

	groupId := eventRequest.GroupId

	result, err := e.EventRepository.GetEventsByGroupId(groupId)

	if err != nil {
		return nil, err
	}

	response := &eventModel.GetEventResponse{}

	response.Data = result

	if eventRequest.GetCountOfParticipants && len(result) > 0 {
		eventIdCountMap, err := GetCountOfEventParticipants(e, result)
		if err != nil {
			return nil, err
		}
		for _, event := range response.Data {
			if val, ok := eventIdCountMap[event.EventId.String()]; ok {
				event.NoOfJoinedParticipants = val
			}
		}
	}

	return response, nil
}
