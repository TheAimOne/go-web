package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type EventMember struct {
	Id          int64     `json:"id"`
	EventId     uuid.UUID `json:"eventId"`
	GroupId     uuid.UUID `json:"groupId"`
	MemberId    uuid.UUID `json:"memberId"`
	Action      string    `json:"action"`
	MemberName  string    `json:"name"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	CreateTime  time.Time
	UpdateTime  time.Time
	DeletedTine time.Time
}

type GetEventMembersRequest struct {
	EventId string `json:"eventId"`
}

type GetEventMembersResponse struct {
	Members []*EventMember `json:"eventId"`
}
