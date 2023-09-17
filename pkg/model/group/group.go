package model

import uuid "github.com/satori/go.uuid"

type Group struct {
	GroupId     uuid.UUID `json:"groupId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Size        int64     `json:"size"`
}

type GroupMember struct {
	GroupId  uuid.UUID `json:"groupId"`
	MemberId uuid.UUID `json:"memberId"`
	IsAdmin  bool      `json:"isAdmin"`
}

type CreateGroupModel struct {
	GroupInfo Group          `json:"groupInfo"`
	Members   []*GroupMember `json:"members"`
}
