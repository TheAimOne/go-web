package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
	authModel "github.com/go-web/pkg/model/auth"
)

func CreateAuthenticationHandler(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO send error
		log.Printf("Error reading body: %v", err)
		return nil, model.NewError(400, "invalid body")
	}

	requestEvent := authModel.AuthRequest{}

	err = json.Unmarshal(body, &requestEvent)
	if err != nil {
		// TODO
		fmt.Println("End Error", err)
		return nil, model.NewError(400, "invalid body")
	}
	resp, err := AuthServiceImpl.Authenticate(&requestEvent)
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
