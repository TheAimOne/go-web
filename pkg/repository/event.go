package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model"
	eventModel "github.com/go-web/pkg/model/event"
)

const tableName = "event"

var columns = []string{
	"event_id", "group_id", "created_by", "venue_id", "name", "type", "status", "params",
	"total_cost", "currency", "no_of_participants", "description", "start_date_time", "end_date_time",
}

var eventDetailQuery = `
	select e.id, e.event_id, e.group_id, e.venue_id, e.created_by, e.name, e.type, e.status,  
	e.total_cost, e.currency, e.no_of_participants, e.description, e.start_date_time, 
	e.end_date_time, v.name "venueName", v.address "venueAddress", v.latitude, v.longitude
	from "event" e left outer join venue v 
	on e.venue_id = v.id `

var eventFilterMap = map[string]string{
	"eventId":          "e.id",
	"groupId":          "e.group_id",
	"venueId":          "e.venue_id",
	"status":           "e.status",
	"startDateAndTime": "e.start_date_time",
	"venueName":        "v.name",
}

type EventRepository interface {
	CreateEvent(*eventModel.Event) error
	GetEventsByGroupId(groupId string) ([]*eventModel.EventDetail, error)
	GetEventsByFilter(filter *model.Filter) ([]*eventModel.EventDetail, error)
}

func NewEventRepository(DB function.DBFunction) EventRepository {
	return &eventRepoImpl{
		DB,
	}
}

type eventRepoImpl struct {
	DB function.DBFunction
}

func (e *eventRepoImpl) CreateEvent(event *eventModel.Event) error {
	// TODO since this is buggy way of insert - order of colums should match values
	// need to come up with alternate design
	values := []interface{}{
		event.EventId,
		event.GroupId,
		event.CreatorId,
		event.VenueId,
		event.Name,
		event.Type,
		event.Status,
		event.Params,
		event.TotalCost,
		event.Currency,
		event.NoOfParticipants,
		event.Description,
		event.StartDateAndTime,
		event.EndDateAndTime,
	}

	err := e.DB.Insert(tableName, columns, values)
	return err
}

func (e *eventRepoImpl) GetEventsByGroupId(groupId string) ([]*eventModel.EventDetail, error) {

	finalQuery := eventDetailQuery + " where e.group_id = '%s' "

	query := fmt.Sprintf(finalQuery, groupId)
	rows, err := e.DB.SelectRaw(query)

	return getEventValuesFromRows(rows, err)
}

func getEventValuesFromRows(rows *sql.Rows, err error) ([]*eventModel.EventDetail, error) {
	result := make([]*eventModel.EventDetail, 0)
	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}

	for rows.Next() {
		var evt eventModel.EventDetail
		rows.Scan(&evt.Id, &evt.EventId, &evt.GroupId, &evt.VenueId, &evt.CreatorId, &evt.Name,
			&evt.Type, &evt.Status, &evt.TotalCost, &evt.Currency, &evt.NoOfParticipants, &evt.Description,
			&evt.StartDateAndTime, &evt.EndDateAndTime,
			&evt.VenueName, &evt.VenueAddress, &evt.Latitude, &evt.Longitude)
		result = append(result, &evt)
	}
	return result, nil
}

func (e *eventRepoImpl) GetEventsByFilter(filter *model.Filter) ([]*eventModel.EventDetail, error) {
	rows, err := e.DB.SelectPaginateAndFilterByQuery(eventDetailQuery, *filter, eventFilterMap)
	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}
	return getEventValuesFromRows(rows, err)
}
