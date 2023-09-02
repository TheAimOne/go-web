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
)
