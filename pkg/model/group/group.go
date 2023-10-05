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
	Status   string    `json:"status"`
	model.Audit
}

type AddMembersToGroupRequest struct {
	GroupId uuid.UUID      `json:"groupId"`
	Members []*GroupMember `json:"members"`
}

type CreateGroupModel struct {
	GroupInfo Group          `json:"groupInfo"`
	Members   []*GroupMember `json:"members"`
}

type GroupMemberByIdResponse struct {
	MemberId  uuid.UUID `json:"memberId"`
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"isAdmin"`
	Status    string    `json:"status"`
}

type GroupsByNameRequest struct {
	Name    string `json:"name"`
	Page    int    `json:"page"`
	PerPage int    `json:"perPage"`
}

type GroupsByMemberResponse struct {
	Data []*Group `json:"data"`
}

const (
	ACTIVE = "ACTIVE"
)
