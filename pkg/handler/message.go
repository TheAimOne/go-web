package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/go-web/pkg/model"
	messageModel "github.com/go-web/pkg/model/message"
	uuid "github.com/satori/go.uuid"
)

func CreateMessageHandler(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO send error
		log.Printf("Error reading body: %v", err)
		return nil, model.NewError(400, "invalid body")
	}

	requestMessage := messageModel.SendMessageReq{}

	err = json.Unmarshal(body, &requestMessage)
	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, "invalid body")
	}

	resp, err := MessageServiceImpl.CreateMessage(&requestMessage)

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

func RetrieveMessageHandler(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)

	eventId := r.URL.Query().Get("eventId")
	groupId := r.URL.Query().Get("groupId")
	offsetString := r.URL.Query().Get("offset")

	offset := 0
	if len(offsetString) != 0 {
		o, err := strconv.Atoi(offsetString)
		if err != nil {
			return nil, model.NewError(400, err.Error())
		}
		offset = o
	}

	rq := messageModel.RetrieveMessageReq{
		EventId: uuid.FromStringOrNil(eventId),
		GroupId: uuid.FromStringOrNil(groupId),
		Offset:  offset,
	}

	resp, err := MessageServiceImpl.RetrieveMessageForEvent(&rq)
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
