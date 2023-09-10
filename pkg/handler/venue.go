package handler

import (
	"log"

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

	newVenue, err := VenueServiceImpl.CreateVenue(requestVenue)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(newVenue), nil
}

func GetVenueHandler(request interface{}) (*model.Response, error) {
	requestFilter, err := util.ReadJson[model.Filter](request, model.Filter{})
	if err != nil {
		return nil, err
	}

	venueList, err := VenueServiceImpl.GetVenues(requestFilter)
	if err != nil {
		log.Println(err)
		return nil, model.NewError(500, err.Error())
	}

	return util.GetResponse(venueList), nil
}
