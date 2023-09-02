package repository

import (
	"encoding/json"
	"fmt"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/event"
)

const tableName = "event"

var columns = []string{
	"event_id", "group_id", "creator_id", "name", "type", "status", "params",
}

type EventRepository interface {
	CreateEvent(*model.Event) error
	GetEventsByGroupId(groupId string) ([]*model.Event, error)
}

func NewEventRepository(DB function.DBFunction) EventRepository {
	return &eventRepoImpl{
		DB,
	}
}

type eventRepoImpl struct {
	DB function.DBFunction
}

func (e *eventRepoImpl) CreateEvent(event *model.Event) error {
	// TODO since this is buggy way of insert - order of colums should match values
	// need to come up with alternate design
	values := []interface{}{
		event.EventId,
		event.GroupId,
		event.CreatorId,
		event.Name,
		event.Type,
		event.Status,
		event.Params,
	}

	err := e.DB.Insert(tableName, columns, values)
	return err
}

func (e *eventRepoImpl) GetEventsByGroupId(groupId string) ([]*model.Event, error) {
	result := make([]*model.Event, 0)

	condition := fmt.Sprintf("where group_id='%s'", groupId)
	rows, err := e.DB.SelectAll(tableName, condition, columns)

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}

	for rows.Next() {
		var e model.Event
		rows.Scan(&e.EventId, &e.GroupId, &e.CreatorId, &e.Name, &e.Type, &e.Status, &e.Params)

		err = json.Unmarshal(e.Params.([]byte), &e.Params)

		if err != nil {
			return nil, constants.ErrorParsingRecordsFromDB
		}

		result = append(result, &e)
	}

	return result, nil
}
