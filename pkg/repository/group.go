package repository

import (
	"fmt"

	"github.com/go-web/database/function"
	model "github.com/go-web/pkg/model/group"
	uuid "github.com/satori/go.uuid"
)

const groupTableName = "group"
const groupMemberTableName = "group_member"

var groupColumns = []string{
	"group_id", "name", "description", "size",
}

var groupMemberColumns = []string{
	"group_id", "member_id", "is_admin",
}

type GroupRepository interface {
	CreateGroup(group *model.Group) error
	AddMembersToGroup(groupMember []*model.GroupMember) error
	GetGroupById(groupId string) (*model.Group, error)
	GetMembersByGroupId(groupId string) ([]*model.GroupMember, error)
	GetGroupsByMemberId(memberId string) ([]*model.Group, error)
}

func NewGroupRepository(DB function.DBFunction) GroupRepository {
	return &groupRepoImpl{
		DB: DB,
	}
}

type groupRepoImpl struct {
	DB function.DBFunction
}

func (g *groupRepoImpl) CreateGroup(group *model.Group) error {
	if group.GroupId == uuid.Nil {
		group.GroupId = uuid.NewV4()
	}

	values := []interface{}{
		group.GroupId,
		group.Name,
		group.Description,
		group.Size,
	}

	err := g.DB.Insert(groupTableName, groupColumns, values)
	return err
}

func (g *groupRepoImpl) AddMembersToGroup(groupMembers []*model.GroupMember) error {
	for _, groupMember := range groupMembers {
		values := []interface{}{
			groupMember.GroupId,
			groupMember.MemberId,
			groupMember.IsAdmin,
		}

		err := g.DB.Insert(groupMemberTableName, groupMemberColumns, values)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *groupRepoImpl) GetGroupById(groupId string) (*model.Group, error) {
	row, err := g.DB.Select(groupTableName, fmt.Sprintf(" group_id = %s ", groupId), groupColumns)
	if err != nil {
		return nil, err
	}

	var group model.Group
	row.Scan(&group.GroupId, &group.Name, &group.Description, &group.Size)

	return &group, nil
}

// TO-DO
func (g *groupRepoImpl) GetMembersByGroupId(groupId string) ([]*model.GroupMember, error) {

	return nil, nil
}

// TO-DO
func (g *groupRepoImpl) GetGroupsByMemberId(memberId string) ([]*model.Group, error) {
	return nil, nil
}
