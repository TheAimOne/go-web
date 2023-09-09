package repository

import (
	"log"

	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	"github.com/go-web/pkg/model"
	venueModel "github.com/go-web/pkg/model/venue"
	uuid "github.com/satori/go.uuid"
)

const venueTableName = "venue"

var venueColumns = []string{"venue_id", "name", "address", "latitude", "longitude", "opening_time", "closing_time", "rating"}
var filterMap = map[string]string{
	"NAME":         "name",
	"ADDRESS":      "address",
	"RATING":       "rating",
	"OPENING_TIME": "opening_time",
	"CLOSING_TIME": "closing_time",
}

type VenueRepository interface {
	CreateVenue(*venueModel.Venue) error
	GetVenues(*model.Filter) ([]*venueModel.Venue, error)
}

func NewVenueRepository(DB function.DBFunction) VenueRepository {
	return &VenueRepoImpl{
		DB: DB,
	}
}

type VenueRepoImpl struct {
	DB function.DBFunction
}

func (v *VenueRepoImpl) CreateVenue(venue *venueModel.Venue) error {
	if venue.Id == uuid.Nil {
		venue.Id = uuid.NewV4()
	}
	values := []interface{}{
		venue.Id,
		venue.Name,
		venue.Address,
		venue.Latitude,
		venue.Longitude,
		venue.OpeningTime,
		venue.ClosingTime,
		venue.Rating,
	}

	err := v.DB.Insert(venueTableName, venueColumns, values)
	return err
}

func (v *VenueRepoImpl) GetVenues(filter *model.Filter) ([]*venueModel.Venue, error) {
	result := make([]*venueModel.Venue, 0)

	rows, err := v.DB.SelectPaginateAndFilter(venueTableName, *filter, venueColumns, filterMap)

	if err != nil {
		log.Println(err)
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}

	for rows.Next() {
		var venue venueModel.Venue
		rows.Scan(&venue.Id, &venue.Name, &venue.Address, &venue.Latitude, &venue.Longitude, &venue.OpeningTime, &venue.ClosingTime, &venue.Rating)
		result = append(result, &venue)
	}

	return result, nil
}
