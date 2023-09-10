package model

import uuid "github.com/satori/go.uuid"

type Venue struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	OpeningTime string    `json:"openingTime"`
	ClosingTime string    `json:"closingTime"`
	Rating      int32     `json:"rating"`
}

type GetVenueResponse struct {
	Venues []*Venue `json:"venues"`
}
