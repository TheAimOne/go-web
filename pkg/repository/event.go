package repository

import (
	"github.com/go-web/database/function"
	model "github.com/go-web/pkg/model/event"
)

const tableName = "event"

var columns = []string{
	"event_id", "group_id", "creator_id", "name", "type", "status", "params",
}

type EventRepository interface {
	CreateEvent(*model.Event) error
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
