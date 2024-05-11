package message

import (
	"fmt"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	"github.com/go-web/pkg/model/message"
)

const tableName = "message"

var columns = []string{
	"member_id", "name", "group_id", "event_id", "content",
}

type IMessageRepo interface {
	CreateMessage(m *message.Message) error
	RetrieveMessageForEvent(eventId string, offset int) ([]*message.Message, error)
}

func NewMessageRepo(DB function.DBFunction) IMessageRepo {
	return &repoImpl{
		DB,
	}
}

type repoImpl struct {
	DB function.DBFunction
}

func (r *repoImpl) CreateMessage(m *message.Message) error {
	values := []interface{}{
		m.MemberId,
		m.Name,
		m.GroupId,
		m.EventId,
		m.Content,
	}

	err := r.DB.Insert(tableName, columns, values)

	return err
}

func (r *repoImpl) RetrieveMessageForEvent(eventId string, offset int) ([]*message.Message, error) {
	result := make([]*message.Message, 0)
	query := `
	select 
		m.event_id,
		m.group_id,
		m.member_id,
		m.name,
		m.content,
		m.create_time,
		m.update_time
	from message m
	where
		m.event_id='%s'
		order by m.create_time asc
		OFFSET %d
		LIMIT 20
	`

	rows, err := r.DB.SelectRaw(fmt.Sprintf(query, eventId, offset))

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}

	for rows.Next() {
		var e message.Message
		rows.Scan(&e.EventId, &e.GroupId, &e.MemberId, &e.Name, &e.Content, &e.CreateTime, &e.UpdateTime)

		result = append(result, &e)
	}

	return result, nil
}
