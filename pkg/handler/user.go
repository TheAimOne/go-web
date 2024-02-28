package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-web/pkg/model"
	userModel "github.com/go-web/pkg/model/user"
	"github.com/go-web/pkg/util"
)

func CreateUserHandler(request interface{}) (*model.Response, error) {
	requestUser, err := util.ReadJson(request, userModel.UserBase{})
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

func GetUsersHandler(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	page := r.URL.Query().Get("page")
	perPage := r.URL.Query().Get("perPage")
	log.Printf("page: %s, perPage: %s", page, perPage)
	// TODO validate pagination
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Invalid page")
		pageInt = 0
	}

	perPageInt, err := strconv.Atoi(perPage)
	if err != nil {
		log.Println("Invalid perPage")
		perPageInt = 0
	}

	request1 := userModel.GetUsersRequest{
		Page:    pageInt,
		PerPage: perPageInt,
	}

	user, err := UserServiceImpl.GetUsers(request1)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(user), nil
}

func SearchUserHandler(request interface{}) (*model.Response, error) {
	filterRequest, err := util.ReadJson(request, userModel.UserFilter{})
	if err != nil {
		return nil, model.NewError(400, err.Error())
	}

	users, err := UserServiceImpl.SearchUsers(*filterRequest)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(users), nil
}
