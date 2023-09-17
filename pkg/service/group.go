package service

import (
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/group"
	"github.com/go-web/pkg/repository"
)

type GroupImpl struct {
	groupRepository repository.GroupRepository
}

func (g *GroupImpl) CreateGroupWithMembers(group *model.CreateGroupModel) (*model.CreateGroupModel, error) {
	err := g.groupRepository.CreateGroup(&group.GroupInfo)
	if err != nil {
		return nil, constants.ErrorCreatingGroup
	}

	if _, err := g.AddMembersToGroup(group.Members); err != nil {
		return nil, err
	}

	return group, nil
}

func (g *GroupImpl) AddMembersToGroup(groupMembers []*model.GroupMember) ([]*model.GroupMember, error) {
	if len(groupMembers) != 0 {
		err := g.groupRepository.AddMembersToGroup(groupMembers)
		if err != nil {
			return nil, constants.ErrorCreatingGroupMembers
		}
	}

	return groupMembers, nil
}

func (g *GroupImpl) GetGroup(groupId string) (*model.Group, error) {
	return nil, nil
}

func (g *GroupImpl) GetMembersByGroupId(groupId string) ([]*model.GroupMember, error) {
	return nil, nil
}

func (g *GroupImpl) GetGroupsByMemberId(memberId string) ([]*model.Group, error) {
	return nil, nil
}
