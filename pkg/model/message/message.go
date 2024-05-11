package message

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SendMessageReq struct {
	GroupId  uuid.UUID `json:"groupId"`
	EventId  uuid.UUID `json:"eventId"`
	MemberId uuid.UUID `json:"memberId"`
	Name     string    `json:"name"`
	Content  string    `json:"content"`
}

type SendMessageResp struct {
	Status string `json:"status"`
}

type RetrieveMessageReq struct {
	GroupId uuid.UUID `json:"groupId"`
	EventId uuid.UUID `json:"eventId"`
	Offset  int       `json:"offset"`
}

type RetrieveMessageResp struct {
	Data []*Message `json:"message"`
}

type Message struct {
	MemberId   uuid.UUID `json:"memberId"`
	Name       string    `json:"name"`
	GroupId    uuid.UUID `json:"groupId"`
	EventId    uuid.UUID `json:"eventId"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	DeleteTime time.Time `json:"deleteTime"`
}
