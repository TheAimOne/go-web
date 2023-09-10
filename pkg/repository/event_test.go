package repository

import (
	"testing"

	mock "github.com/go-web/mocks/database/function"
)

var dbFunctionMock *mock.DBFunction
var eventRepository EventRepository

func init() {
	dbFunctionMock = new(mock.DBFunction)
	eventRepository = NewEventRepository(dbFunctionMock)
}

func TestCreateEvent(t *testing.T) {
	t.Run("DB Error", func(t *testing.T) {

	})

	t.Run("DB Success", func(t *testing.T) {

	})

}

func TestGetEventsByGroupId(t *testing.T) {

}
