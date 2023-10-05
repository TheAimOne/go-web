package repository

import (
	"fmt"
	"log"

	"github.com/go-web/database/function"
	model "github.com/go-web/pkg/model/group"
)

const groupTableName = `"group"`
const groupMemberTableName = "group_member"

var groupColumns = []string{
	"group_id", "name", "description", "size",
}

var groupMemberColumns = []string{
	"group_id", "member_id", "is_admin", "status",
}

type GroupRepository interface {
	CreateGroup(group *model.Group) error
	AddMembersToGroup(groupMember []*model.GroupMember) error
	GetGroupById(groupId string) (*model.Group, error)
	GetMembersByGroupId(groupId string) ([]*model.GroupMemberByIdResponse, error)
	GetGroupsByMemberId(memberId string) ([]*model.Group, error)
	GetGroups(name string, page, perPage int) ([]*model.Group, error)
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

	values := []interface{}{
		group.GroupId,
		group.Name,
		group.Description,
		group.Size,
	}
	log.Println(values...)

	err := g.DB.Insert(groupTableName, groupColumns, values)
	return err
}

func (g *groupRepoImpl) AddMembersToGroup(groupMembers []*model.GroupMember) error {
	for _, groupMember := range groupMembers {
		values := []interface{}{
			groupMember.GroupId,
			groupMember.MemberId,
			groupMember.IsAdmin,
			groupMember.Status,
		}

		err := g.DB.Insert(groupMemberTableName, groupMemberColumns, values)
		if err != nil {
			log.Println(err)
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
func (g *groupRepoImpl) GetMembersByGroupId(groupId string) ([]*model.GroupMemberByIdResponse, error) {
	query := `SELECT u.member_id, u."name", u.short_name, u.email, u.mobile, gm.status, gm.is_admin 
	FROM "user" u INNER JOIN group_member gm ON gm.member_id = u.member_id 
	INNER JOIN "group" g ON g.group_id = gm.group_id where g.group_id = '%s'`

	rows, err := g.DB.SelectRaw(fmt.Sprintf(query, groupId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*model.GroupMemberByIdResponse, 0)

	for rows.Next() {
		var groupMember model.GroupMemberByIdResponse
		rows.Scan(&groupMember.MemberId, &groupMember.Name, &groupMember.ShortName, &groupMember.Email,
			&groupMember.Mobile, &groupMember.Status, &groupMember.IsAdmin)
		result = append(result, &groupMember)
	}

	return result, nil
}

// TO-DO
func (g *groupRepoImpl) GetGroupsByMemberId(memberId string) ([]*model.Group, error) {
	query := `select g.group_id, g.name, g.description, g.size
	from "group" g inner join group_member gm 
	on gm.group_id = g.group_id where gm.member_id = '%s' `

	rows, err := g.DB.SelectRaw(fmt.Sprintf(query, memberId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*model.Group, 0)

	for rows.Next() {
		var group model.Group
		rows.Scan(&group.GroupId, &group.Name, &group.Description, &group.Size)
		result = append(result, &group)
	}

	return result, nil
}

func (g *groupRepoImpl) GetGroups(name string, page, perPage int) ([]*model.Group, error) {
	limit := perPage
	offset := page - 1

	query := `select group_id, g.name, g.description, g.size
	from "group" g`

	if name != "" {
		query = fmt.Sprintf("%s where name=%s", query, name)
	}

	query = fmt.Sprintf("%s limit %d offset %d", query, limit, offset)

	rows, err := g.DB.SelectRaw(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]*model.Group, 0)

	for rows.Next() {
		var group model.Group
		rows.Scan(&group.GroupId, &group.Name, &group.Description, &group.Size)
		result = append(result, &group)
	}

	return result, nil
}
