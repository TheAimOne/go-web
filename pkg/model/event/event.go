package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event struct {
	Id          int64       `json:"id"`
	EventId     uuid.UUID   `json:"eventId"`
	GroupId     uuid.UUID   `json:"groupId"`
	CreatorId   uuid.UUID   `json:"creatorId"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Status      string      `json:"status"`
	Params      interface{} `json:"params"`
	CreateTime  time.Time
	UpdateTime  time.Time
	DeletedTine time.Time
}

type EventResponse struct {
	EventId uuid.UUID `json:"eventId"`
	Status  string    `json:"status"`
}

type GetEventRequest struct {
	GroupId string
}

type GetEventResponse struct {
	Data []*Event `json:"eventId"`
}

type AddMemberToEventRequest struct {
	EventId  uuid.UUID `json:"eventId"`
	GroupId  uuid.UUID `json:"groupId"`
	MemberID uuid.UUID `json:"memberId"`
	Action   string    `json:"action"`
	Status   string    `json:"status"`
}

type AddMemberToEventResponse struct {
	Status string `json:"status"`
}
