package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
)

func WriteJson(body interface{}, w http.ResponseWriter) {
	j, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Error occured while serializing JSON", err)
	}
	_, err = w.Write(j)

	if err != nil {
		log.Fatal(err)
	}
}

func ReadJson[T any](request interface{}, object T) (*T, error) {
	r := request.(*http.Request)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(400, "Invalid body")
	}
	err = json.Unmarshal(body, &object)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(400, "Invalid body")
	}
	return &object, nil
}

func GetResponse(resp interface{}) *model.Response {
	response := model.Response{
		Body: resp,
	}
	return response.Json()
}
