package service

import (
	"errors"
	"log"

	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/group"
	"github.com/go-web/pkg/repository"
	uuid "github.com/satori/go.uuid"
)

type GroupImpl struct {
	groupRepository repository.GroupRepository
}

func (g *GroupImpl) CreateGroupWithMembers(group *model.CreateGroupModel) (*model.CreateGroupModel, error) {
	if group.GroupInfo == (model.Group{}) {
		return nil, constants.ErrorCreatingGroup
	}
	if group.GroupInfo.GroupId == uuid.Nil {
		group.GroupInfo.GroupId = uuid.NewV4()
	}
	err := g.groupRepository.CreateGroup(&group.GroupInfo)
	if err != nil {
		log.Println(err)
		return nil, constants.ErrorCreatingGroup
	}

	if len(group.Members) != 0 {
		if _, err := g.AddMembersToGroup(group.GroupInfo.GroupId, group.Members); err != nil {
			return nil, err
		}
	}

	return group, nil
}

func (g *GroupImpl) AddMembersToGroup(groupId uuid.UUID, groupMembers []*model.GroupMember) ([]*model.GroupMember, error) {
	if len(groupMembers) != 0 && groupId != uuid.Nil {
		for _, groupMember := range groupMembers {
			groupMember.GroupId = groupId
			groupMember.Status = model.ACTIVE
		}

		err := g.groupRepository.AddMembersToGroup(groupMembers)
		if err != nil {
			return nil, constants.ErrorCreatingGroupMembers
		}
		return groupMembers, nil
	} else {
		return nil, constants.ErrorCreatingGroupMembers
	}
}

func (g *GroupImpl) GetGroup(groupId string) (*model.Group, error) {
	if groupId == "" {
		return nil, errors.New("Invalid group ID")
	}

	group, err := g.groupRepository.GetGroupById(groupId)

	if err != nil {
		return nil, constants.ErrorFetchingGroup
	}
	return group, nil
}

func (g *GroupImpl) GetMembersByGroupId(groupId string) ([]*model.GroupMemberByIdResponse, error) {
	groupMembers, err := g.groupRepository.GetMembersByGroupId(groupId)
	if err != nil {
		return nil, constants.ErrorFetchingGroupMembers
	}

	return groupMembers, nil
}

func (g *GroupImpl) GetGroupsByMemberId(memberId string) (*model.GroupsByMemberResponse, error) {
	groups, err := g.groupRepository.GetGroupsByMemberId(memberId)

	if err != nil {
		return nil, constants.ErrorFetchingGroup
	}

	response := &model.GroupsByMemberResponse{
		Data: groups,
	}

	return response, nil
}

func (g *GroupImpl) GetGroups(request model.GroupsByNameRequest) (*model.GroupsByMemberResponse, error) {
	page := request.Page
	perPage := request.PerPage

	if request.Page == 0 {
		page = 1
	}

	if request.PerPage == 0 {
		perPage = 10
	}
	groups, err := g.groupRepository.GetGroups(request.Name, page, perPage)

	if err != nil {
		return nil, constants.ErrorFetchingGroup
	}

	response := &model.GroupsByMemberResponse{
		Data: groups,
	}

	return response, nil
}
