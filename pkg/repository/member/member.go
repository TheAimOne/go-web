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
	GetEventMembers(addEventMember *modelMember.GetEventMembersRequest) ([]*modelMember.EventMemberDetail, error)
	CountEventMemberByEventIds(ids string) ([]*modelMember.CountMembersByEventId, error)
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

func (e *eventMemberRepoImpl) GetEventMembers(addEventMember *modelMember.GetEventMembersRequest) ([]*modelMember.EventMemberDetail, error) {
	result := make([]*modelMember.EventMemberDetail, 0)

	query := `
	select 
		e.event_id,
		e.group_id,
		e.member_id,
		e.action,
		e.status,
		u.name,
		u.email,
		u.mobile
	from event_member e
	join "user" u
		on u.member_id = e.member_id
	where
		event_id='%s'
	`

	rows, err := e.DB.SelectRaw(fmt.Sprintf(query, addEventMember.EventId))

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}

	for rows.Next() {
		var e modelMember.EventMemberDetail
		rows.Scan(&e.EventId, &e.GroupId, &e.MemberId, &e.Action, &e.Status, &e.MemberName, &e.Email, &e.Mobile)

		result = append(result, &e)
	}

	return result, nil
}

func (e *eventMemberRepoImpl) CountEventMemberByEventIds(ids string) ([]*modelMember.CountMembersByEventId, error) {
	result := make([]*modelMember.CountMembersByEventId, 0)

	query := ` select em.event_id, count(*) from event_member em where em.event_id in (%s)
		group by em.event_id
	`
	rows, err := e.DB.SelectRaw(fmt.Sprintf(query, ids))

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}

	for rows.Next() {
		var cm modelMember.CountMembersByEventId
		rows.Scan(&cm.EventId, &cm.Count)
		result = append(result, &cm)
	}
	return result, nil
}
