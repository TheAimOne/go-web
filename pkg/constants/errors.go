package constants

import "errors"

// DB Errors
var (
	ErrorCreatingSql          = errors.New("error creating sql")
	ErrorReadingFromDB        = errors.New("error reading from db")
	ErrorNoRecordsInDB        = errors.New("no records in db")
	ErrorParsingRecordsFromDB = errors.New("error parsing records from db")
)

// Handler Errors
var (
	ErrorReadingBody = errors.New("error reading body")
)

// Service Errors
var (
	ErrorAuthenticating          = errors.New("Error in authentication")
	ErrorCreatingToken           = errors.New("Error creating token")
	ErrorAuthTypeNotProvided     = errors.New("Auth type not provided")
	ErrorRefreshTokenNotProvided = errors.New("Refresh token not provided")
	ErrorRefreshTokenExpired     = errors.New("Refresh token expired")
	ErrorInvalidDetails          = errors.New("Invalid Details")

	ErrorParsingParams = errors.New("invalid Params")
	ErrorCreatingEvent = errors.New("error creating event")

	ErrorGettingEventMembers   = errors.New("error getting event members")
	ErrorCreatingVenue         = errors.New("error creating Venue")
	ErrorGettingVenues         = errors.New("error Getting Venues")
	ErrorGettingCountByEventId = errors.New("error getting count by Event ID")

	ErrorCreatingGroup        = errors.New("error Creating Group")
	ErrorFetchingGroup        = errors.New("error Fetching Group")
	ErrorCreatingGroupMembers = errors.New("error adding or already joined")
	ErrorFetchingGroupMembers = errors.New("error Fetching group members")

	ErrorGettingUser  = errors.New("error Getting User")
	ErrorCreatingUser = errors.New("error Creating User")

	ErrorJoiningEvent = errors.New("error joining or already joined")

	ErrorFilterIsNull  = errors.New("filter object is empty")
	ErrorPagination    = errors.New("page number or page size is invalid")
	ErrorSearchingUser = errors.New("error Searching User")

	ErrorCreatingMessage = errors.New("error creating message")
	ErrorRetrieveMessage = errors.New("error retrieving message")
)
