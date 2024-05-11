package service

import (
	"github.com/go-web/pkg/constants"
	model "github.com/go-web/pkg/model/message"
	"github.com/go-web/pkg/repository/message"
)

const StatusOk = "Ok!"

type messageImpl struct {
	messageRepo message.IMessageRepo
}

func NewMessageService(messageRepo message.IMessageRepo) MessageService {
	return &messageImpl{
		messageRepo,
	}
}

func (s *messageImpl) CreateMessage(m *model.SendMessageReq) (*model.SendMessageResp, error) {
	// TODO 1. Check whether user belongs to group and event
	message := &model.Message{
		MemberId: m.MemberId,
		GroupId:  m.GroupId,
		EventId:  m.EventId,
		Name:     m.Name,
		Content:  m.Content,
	}

	err := s.messageRepo.CreateMessage(message)

	if err != nil {
		return nil, constants.ErrorCreatingMessage
	}

	return &model.SendMessageResp{
		Status: StatusOk,
	}, nil
}

func (s *messageImpl) RetrieveMessageForEvent(m *model.RetrieveMessageReq) (*model.RetrieveMessageResp, error) {
	// TODO 1. Check whether user belongs to group and event

	res, err := s.messageRepo.RetrieveMessageForEvent(m.EventId.String(), m.Offset)

	if err != nil {
		return nil, constants.ErrorRetrieveMessage
	}

	return &model.RetrieveMessageResp{
		Data: res,
	}, nil
}
