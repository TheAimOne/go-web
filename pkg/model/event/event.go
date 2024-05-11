package model

import (
	"time"

	model "github.com/go-web/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type EventType string

const (
	Sports EventType = "sports"
)

type Event struct {
	Id                     int64       `json:"id"`
	EventId                uuid.UUID   `json:"eventId"`
	GroupId                uuid.UUID   `json:"groupId"`
	VenueId                uuid.UUID   `json:"venueId"`
	CreatorId              uuid.UUID   `json:"creatorId"`
	Name                   string      `json:"name"`
	Description            string      `json:"description"`
	Type                   EventType   `json:"type"`
	Status                 string      `json:"status"`
	Params                 interface{} `json:"params"`
	TotalCost              float64     `json:"totalCost"`
	Currency               string      `json:"currency"`
	NoOfParticipants       int64       `json:"noOfParticipants"`
	NoOfJoinedParticipants int64       `json:"noOfJoinedParticipants"`
	StartDateAndTime       time.Time   `json:"startDateAndTime"`
	EndDateAndTime         time.Time   `json:"endDateAndTime"`
	CreateTime             time.Time
	UpdateTime             time.Time
	DeletedTine            time.Time
}

type EventDetail struct {
	Event
	CostPerPerson float32 `json:"costPerPerson"`
	VenueName     string  `json:"venueName"`
	VenueAddress  string  `json:"venueAddress"`
	Latitude      int64   `json:"latitude"`
	Longitude     int64   `json:"longitude"`
}

type EventResponse struct {
	EventId uuid.UUID `json:"eventId"`
	Status  string    `json:"status"`
}

type GetEventRequest struct {
	GroupId                string
	GetCountOfParticipants bool
}

type GetEventResponse struct {
	Data []*EventDetail `json:"data"`
}

type EventFilter struct {
	Filter                 model.Filter
	GetCountOfParticipants bool
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
