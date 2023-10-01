package handler

import (
	"log"
	"net/http"

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
	request, err := util.ReadJson[groupModel.AddMembersToGroupRequest](request, groupModel.AddMembersToGroupRequest{})
	addMembersToGroupRequest := request.(*groupModel.AddMembersToGroupRequest)
	if err != nil {
		return nil, model.NewError(500, err.Error())
	}
	newMembers, err := GroupServiceImpl.AddMembersToGroup(addMembersToGroupRequest.GroupId, addMembersToGroupRequest.Members)
	if err != nil {
		return nil, model.NewError(500, err.Error())
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
