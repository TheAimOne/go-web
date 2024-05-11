package service

import (
	"errors"
	"testing"

	mockMessageRepo "github.com/go-web/mocks/pkg/repository/message"
	model "github.com/go-web/pkg/model/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var messageRepository *mockMessageRepo.IMessageRepo
var messageServiceImpl MessageService

func init() {
	messageRepository = new(mockMessageRepo.IMessageRepo)
	messageServiceImpl = NewMessageService(messageRepository)
}

func TestCreateMessage(t *testing.T) {
	t.Run("Create message error", func(t *testing.T) {
		m := &model.SendMessageReq{}

		messageRepository.On("CreateMessage", mock.Anything).
			Return(errors.New("error creating")).
			Once()

		res, err := messageServiceImpl.CreateMessage(m)

		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.Equal(t, "error creating message", err.Error())
	})

	t.Run("Create message success", func(t *testing.T) {
		m := &model.SendMessageReq{}

		messageRepository.On("CreateMessage", mock.Anything).
			Return(nil).
			Once()

		res, err := messageServiceImpl.CreateMessage(m)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, "Ok!", res.Status)
	})
}
