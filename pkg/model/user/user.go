package model

import (
	"github.com/go-web/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type UserBase struct {
	MemberId  uuid.UUID `json:"userId"`
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
}

type User struct {
	UserBase
	model.Audit
	Id     uint64     `json:"id"`
	Status UserStatus `json:"status"`
}

type DeviceType string

const (
	MOBILE DeviceType = "MOBILE"
)

type UserSession struct {
	SessionId    uuid.UUID  `json:"sessionId"`
	UserId       uuid.UUID  `json:"userId"`
	SessionToken string     `json:"sessionToken"`
	DeviceId     string     `json:"deviceId"`
	DeviceType   DeviceType `json:"deviceType"`
	ExpiryTime   int        `json:"expiryTime"`
}

type GetUsersRequest struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type GetUsersResponse struct {
	Users []*User `json:"data"`
}

type UserStatus string

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
)

type UserFilter struct {
	Filter               model.Filter `json:"filter"`
	GroupId              string       `json:"groupId"`
	ExcludeUserByGroupId bool         `json:"excludeUserByGroupId"`
}
