package model

import (
	"github.com/go-web/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type UserBase struct {
	MemberId  uuid.UUID `json:"UserId"`
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

type UserStatus string

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
)
