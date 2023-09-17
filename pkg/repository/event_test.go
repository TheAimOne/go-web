package repository

import (
	"errors"
	"testing"

	mockFunction "github.com/go-web/mocks/database/function"
	model "github.com/go-web/pkg/model/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var dbFunctionMock *mockFunction.DBFunction
var eventRepository EventRepository

func init() {
	dbFunctionMock = new(mockFunction.DBFunction)
	eventRepository = NewEventRepository(dbFunctionMock)
}

func TestCreateEvent(t *testing.T) {
	t.Run("DB Error", func(t *testing.T) {
		dbFunctionMock.
			On("Insert", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("Error from DB")).
			Once()

		err := eventRepository.CreateEvent(&model.Event{})

		assert.NotNil(t, err)
	})

	t.Run("DB Success", func(t *testing.T) {
		dbFunctionMock.
			On("Insert", mock.Anything, mock.Anything, mock.Anything).
			Return(nil).
			Once()

		err := eventRepository.CreateEvent(&model.Event{})

		assert.Nil(t, err)
	})
}

func TestGetEventsByGroupId(t *testing.T) {
	// t.Run("DB Error", func(t *testing.T) {
	// 	dbFunctionMock.
	// 		On("SelectAll", mock.Anything, mock.Anything, mock.Anything).
	// 		Return(errors.New("Error from DB")).
	// 		Once()

	// 	err := eventRepository.CreateEvent(&model.Event{})

	// 	assert.NotNil(t, err)
	// })
}
