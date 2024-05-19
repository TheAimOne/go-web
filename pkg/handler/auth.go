package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
	authModel "github.com/go-web/pkg/model/auth"
	"github.com/go-web/pkg/util"
)

func CreateAuthenticationHandler(request interface{}, rw http.ResponseWriter) (*model.Response, error) {
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
	refreshTokenCookie := http.Cookie{Name: "refresh-token", Value: resp.Data.RefreshToken, Expires: resp.Data.RefreshTokenExpiry}
	deviceIdCookie := http.Cookie{Name: "device-id", Value: resp.Data.DeviceId, Expires: resp.Data.RefreshTokenExpiry}
	rw.Header().Add("Set-Cookie", refreshTokenCookie.String())
	rw.Header().Add("Set-Cookie", deviceIdCookie.String())

	return response.Json(), nil
}

func GenerateAuthTokenHandler(request interface{}, rw http.ResponseWriter) (*model.Response, error) {
	r := request.(*http.Request)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, model.NewError(400, "invalid body")
	}
	requestEvent := authModel.AuthRequest{}

	err = json.Unmarshal(body, &requestEvent)
	if err != nil {
		return nil, model.NewError(400, "invalid body")
	}

	resp, err := AuthServiceImpl.GenerateToken(requestEvent)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		return nil, model.NewError(http.StatusUnauthorized, err.Error())
	}

	return util.GetResponse(resp), nil
}
