package service

import (
	"log"

	"github.com/go-web/pkg/constants"
	eventModel "github.com/go-web/pkg/model/event"
	memberModel "github.com/go-web/pkg/model/member"
	"github.com/go-web/pkg/repository/member"
)

type eventMemberImpl struct {
	EventMemberRepository member.EventMemberRepository
}

func (e *eventMemberImpl) CreateEventMember(request *eventModel.AddMemberToEventRequest) (*eventModel.AddMemberToEventResponse, error) {

	request.Status = "C"

	err := e.EventMemberRepository.AddEventMember(request)

	response := &eventModel.AddMemberToEventResponse{}

	response.Status = "CONFIRMED"

	return response, err
}

func (e *eventMemberImpl) GetEventMembers(request *memberModel.GetEventMembersRequest) (*memberModel.GetEventMembersResponse, error) {
	response := &memberModel.GetEventMembersResponse{}

	result, err := e.EventMemberRepository.GetEventMembers(request)

	if err != nil {
		log.Printf("Error getting event members from database %v", err)
		return nil, constants.ErrorGettingEventMembers
	}

	response.Members = result

	return response, err
}
