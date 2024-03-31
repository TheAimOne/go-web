package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
	"github.com/go-web/pkg/util"
)

func CreateEventHandler(request interface{}) (*model.Response, error) {

	r := request.(*http.Request)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO send error
		log.Printf("Error reading body: %v", err)
		return nil, model.NewError(400, "invalid body")
	}

	requestEvent := eventModel.Event{}

	err = json.Unmarshal(body, &requestEvent)
	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, "invalid body")
	}

	resp, err := EventServiceImpl.CreateEvent(&requestEvent)

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

func GetEventByGroupIdHandler(request interface{}) (*model.Response, error) {

	r := request.(*http.Request)

	groupId := r.URL.Query().Get("groupId")
	getCountOfParticipants := r.URL.Query().Get("getCountOfParticipants")
	log.Println("GetEventByGroupIdHandler: groupId ", groupId)

	getEventRequest := &eventModel.GetEventRequest{}
	getEventRequest.GroupId = groupId
	getEventRequest.GetCountOfParticipants = false
	if strings.ToLower(getCountOfParticipants) == "true" {
		getEventRequest.GetCountOfParticipants = true
	}

	resp, err := EventServiceImpl.GetEventsByGroupId(getEventRequest)

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

func SearchEventHandler(request interface{}) (*model.Response, error) {
	filterRequest, err := util.ReadJson(request, model.Filter{})
	getCountOfParticipants := request.(*http.Request).URL.Query().Get("getCountOfParticipants")
	log.Println("getCountOfParticipants: ", getCountOfParticipants)

	eventFilterRequest := &eventModel.EventFilter{}
	eventFilterRequest.Filter = *filterRequest
	if strings.ToLower(getCountOfParticipants) == "true" {
		eventFilterRequest.GetCountOfParticipants = true
	}
	if err != nil {
		return nil, model.NewError(400, err.Error())
	}
	resp, err := EventServiceImpl.SearchEvent(eventFilterRequest)

	if err != nil {
		log.Println("Error: ", err)
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(resp), nil
}
