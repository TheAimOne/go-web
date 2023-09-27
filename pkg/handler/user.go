package handler

import (
	"log"
	"net/http"

	"github.com/go-web/pkg/model"
	userModel "github.com/go-web/pkg/model/user"
	"github.com/go-web/pkg/util"
)

func CreateUserHandler(request interface{}) (*model.Response, error) {
	requestUser, err := util.ReadJson[userModel.UserBase](request, userModel.UserBase{})
	if err != nil {
		return nil, err
	}

	newUser, err := UserServiceImpl.CreateUser(requestUser)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(newUser), nil
}

func GetUserByIdHandler(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	memberId := r.URL.Query().Get("memberId")
	log.Println("memberId:", memberId)
	if memberId == "" {
		return nil, model.NewError(400, "Invalid Member id")
	}

	user, err := UserServiceImpl.GetUserByMemberId(memberId)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(user), nil
}
