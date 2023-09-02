package member

import (
	"fmt"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/event"
	modelMember "github.com/go-web/pkg/model/member"
)

const tableName = "event_member"

var columns = []string{
	"event_id", "group_id", "member_id", "action", "status",
}

type EventMemberRepository interface {
	AddEventMember(*model.AddMemberToEventRequest) error
	GetEventMembers(addEventMember *modelMember.GetEventMembersRequest) ([]*modelMember.EventMember, error)
}

func NewEventMemberRepository(DB function.DBFunction) EventMemberRepository {
	return &eventMemberRepoImpl{
		DB,
	}
}

type eventMemberRepoImpl struct {
	DB function.DBFunction
}

func (e *eventMemberRepoImpl) AddEventMember(addEventMember *model.AddMemberToEventRequest) error {
	values := []interface{}{
		addEventMember.EventId,
		addEventMember.GroupId,
		addEventMember.MemberID,
		addEventMember.Action,
		addEventMember.Status,
	}

	err := e.DB.Insert(tableName, columns, values)
	return err
}

func (e *eventMemberRepoImpl) GetEventMembers(addEventMember *modelMember.GetEventMembersRequest) ([]*modelMember.EventMember, error) {
	result := make([]*modelMember.EventMember, 0)

	query := `
	select 
		event_id,
		group_id,
		member_id,
		action,
		status
	from event_member
	where
		event_id='%s'
	`

	rows, err := e.DB.SelectRaw(fmt.Sprintf(query, addEventMember.EventId))

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}

	for rows.Next() {
		var e modelMember.EventMember
		rows.Scan(&e.EventId, &e.GroupId, &e.MemberId, &e.Action, &e.Status)

		result = append(result, &e)
	}

	return result, nil
}
