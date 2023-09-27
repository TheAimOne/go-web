package model

import "time"

type Audit struct {
	CreatedBy   string    `json:"createdBy"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedTime time.Time `json:"updatedTime"`
	DeletedTime time.Time `json:"deletedTime"`
}
