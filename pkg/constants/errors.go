package constants

import "errors"

// DB Errors
var (
	ErrorCreatingSql          = errors.New("Error creating sql")
	ErrorReadingFromDB        = errors.New("Error reading from db")
	ErrorNoRecordsInDB        = errors.New("No records in db")
	ErrorParsingRecordsFromDB = errors.New("Error parsing records from db")
)

// Handler Errors
var (
	ErrorReadingBody = errors.New("Error reading body")
)

// Service Errors
var (
	ErrorParsingParams = errors.New("Invalid Params")
	ErrorCreatingEvent = errors.New("Error creating event")

	ErrorGettingEventMembers = errors.New("Error getting event members")
	ErrorCreatingVenue       = errors.New("Error creating Venue")
	ErrorGettingVenues       = errors.New("Error Getting Venues")

	ErrorCreatingGroup        = errors.New("Error Creating Group")
	ErrorFetchingGroup        = errors.New("Error Fetching Group")
	ErrorCreatingGroupMembers = errors.New("Error adding or already joined")
	ErrorFetchingGroupMembers = errors.New("Error Fetching group members")

	ErrorGettingUser  = errors.New("Error Getting User")
	ErrorCreatingUser = errors.New("Error Creating User")

	ErrorJoiningEvent = errors.New("Error joining or already joined")
)
