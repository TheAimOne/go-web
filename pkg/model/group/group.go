package model

import (
	"github.com/go-web/pkg/model"
	uuid "github.com/satori/go.uuid"
)

type Group struct {
	GroupId     uuid.UUID `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Size        int64     `json:"size"`
	model.Audit
}

type GroupMember struct {
	GroupId  uuid.UUID `json:"groupId"`
	MemberId uuid.UUID `json:"memberId"`
	IsAdmin  bool      `json:"isAdmin"`
	model.Audit
}

type CreateGroupModel struct {
	GroupInfo Group          `json:"groupInfo"`
	Members   []*GroupMember `json:"members"`
}
