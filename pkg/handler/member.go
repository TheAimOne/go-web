package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
	memberModel "github.com/go-web/pkg/model/member"
)

func CreateEventMemberHandler(request interface{}) (*model.Response, error) {

	r := request.(*http.Request)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO send error
		log.Printf("Error reading body: %v", err)
		return nil, model.NewError(400, "invalid body")
	}

	requestEvent := eventModel.AddMemberToEventRequest{}

	err = json.Unmarshal(body, &requestEvent)
	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, "invalid body")
	}

	resp, err := EventMemberServiceImpl.CreateEventMember(&requestEvent)

	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, err.Error())
	}

	response := model.Response{
		Body: resp,
	}

	return response.Json(), nil
}

func GetEventMembers(request interface{}) (*model.Response, error) {

	r := request.(*http.Request)

	eventId := r.URL.Query().Get("eventId")

	getEventMemberRequest := &memberModel.GetEventMembersRequest{}
	getEventMemberRequest.EventId = eventId

	resp, err := EventMemberServiceImpl.GetEventMembers(getEventMemberRequest)

	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, err.Error())
	}

	response := model.Response{
		Body: resp,
	}

	return response.Json(), nil
}
