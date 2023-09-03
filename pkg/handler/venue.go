package handler

import (
	"github.com/go-web/pkg/model"
	venue_model "github.com/go-web/pkg/model/venue"
	"github.com/go-web/pkg/util"
)

// Endpoint for admin to add Venues
func CreateVenueHandler(request interface{}) (*model.Response, error) {
	requestVenue, err := util.ReadJson[venue_model.Venue](request, venue_model.Venue{})
	if err != nil {
		return nil, err
	}

}
