package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-web/pkg/model"
	groupModel "github.com/go-web/pkg/model/group"
	"github.com/go-web/pkg/util"
)

func CreateGroupWithMembershandler(request interface{}) (*model.Response, error) {
	requestGroup, err := util.ReadJson(request, &groupModel.CreateGroupModel{})
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	newGroup, err := GroupServiceImpl.CreateGroupWithMembers(*requestGroup)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(newGroup), nil
}

func AddMembersToGroupHandler(request interface{}) (*model.Response, error) {
	request, err := util.ReadJson(request, groupModel.AddMembersToGroupRequest{})
	addMembersToGroupRequest := request.(*groupModel.AddMembersToGroupRequest)
	if err != nil {
		return nil, model.NewError(400, err.Error())
	}
	newMembers, err := GroupServiceImpl.AddMembersToGroup(addMembersToGroupRequest.GroupId, addMembersToGroupRequest.Members)
	if err != nil {
		return nil, model.NewError(400, err.Error())
	}
	return util.GetResponse(newMembers), nil
}

func GetMembersByGroupId(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	groupId := r.URL.Query().Get("groupId")
	log.Println("groupId:", groupId)

	if groupId == "" {
		return nil, model.NewError(400, "Invalid Group id")
	}

	members, err := GroupServiceImpl.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(members), nil
}

func GetGroupsByMemberId(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	memberId := r.URL.Query().Get("memberId")
	log.Println("memberId:", memberId)

	if memberId == "" {
		return nil, model.NewError(400, "Invalid Member ID")
	}

	groups, err := GroupServiceImpl.GetGroupsByMemberId(memberId)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(groups), nil
}

func GetGroupById(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	groupId := r.URL.Query().Get("groupId")
	log.Println("groupId: ", groupId)

	if groupId == "" {
		return nil, model.NewError(400, "Invalid Group ID")
	}

	group, err := GroupServiceImpl.GetGroup(groupId)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(group), nil
}

func GetGroups(request interface{}) (*model.Response, error) {
	r := request.(*http.Request)
	name := r.URL.Query().Get("name")
	log.Println("name:", name)

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

	request1 := groupModel.GroupsByNameRequest{
		Name:    name,
		Page:    pageInt,
		PerPage: perPageInt,
	}

	groups, err := GroupServiceImpl.GetGroups(request1)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(groups), nil
}
