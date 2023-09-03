package repository

import (
	"github.com/go-web/database/function"
	"github.com/go-web/pkg/constants"
	"github.com/go-web/pkg/model"
	venueModel "github.com/go-web/pkg/model/venue"
	uuid "github.com/satori/go.uuid"
)

const venueTableName = "venue"

const venueColumns = []string{"venue_id", "name", "address", "latitude", "longitude", "opening_time", "closing_time", "rating"}

type VenueRepository interface {
	CreateVenue(*venueModel.Venue) error
	GetVenues(*model.Filter) (*venueModel.GetVenueResponse, error)
}

func NewVenueRepository(DB function.DBFunction) VenueRepository {

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

func (v *VenueRepoImpl) GetVenues(filter *model.Filter) (*venueModel.GetVenueResponse, error) {
	result := make([]*venueModel.Venue, 0)

	rows, err := v.DB.SelectPaginateAndFilter(venueTableName, *filter, columns)

	if err != nil {
		return nil, constants.ErrorReadingFromDB
	}
	if rows == nil {
		return nil, constants.ErrorNoRecordsInDB
	}

}
