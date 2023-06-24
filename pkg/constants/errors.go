package constants

import "errors"

// DB Errors
var (
	ErrorCreatingSql = errors.New("Error creating sql")
)

// Handler Errors
var (
	ErrorReadingBody = errors.New("Error reading body")
)

// Service Errors
var (
	ErrorParsingParams = errors.New("Invalid Params")
	ErrorCreatingEvent = errors.New("Error creating event")
)
