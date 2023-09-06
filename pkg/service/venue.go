package service

import (
	"fmt"

	"github.com/go-web/pkg/constants"
	"github.com/go-web/pkg/model"
	venueModel "github.com/go-web/pkg/model/venue"
	"github.com/go-web/pkg/repository"
)

type VenueImpl struct {
	venuRepository repository.VenueRepository
}

func (v *VenueImpl) CreateVenue(venueRequest *venueModel.Venue) (*venueModel.Venue, error) {
	err := v.venuRepository.CreateVenue(venueRequest)
	if err != nil {
		return nil, constants.ErrorCreatingVenue
	}
	return venueRequest, err
}

func (v *VenueImpl) GetVenues(request *model.Filter) (*venueModel.GetVenueResponse, error) {
	venues, err := v.venuRepository.GetVenues(request)
	fmt.Println("IN SERVICE")
	if err != nil {
		return nil, constants.ErrorGettingVenues
	}
	venueResponse := venueModel.GetVenueResponse{
		Venues: venues,
	}
	return &venueResponse, nil
}
