package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
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
	fmt.Println("groupId 1 ", groupId)

	getEventRequest := &eventModel.GetEventRequest{}
	getEventRequest.GroupId = groupId

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
