package model

import (
	"errors"
)

var (
	// ErrNotFound is returned when the store failed to find an expected resource.
	ErrNotFound = errors.New("not found in store")
	// ErrNoRowsAffected is returned when no rows were updated during update query.
	ErrNoRowsAffected = errors.New("no rows affected")
	// LatitudeAndLogitudeNotSent error to handle latitude and longitude errors.
	LatitudeAndLogitudeNotSent = errors.New(" Latitude and Longitude were not recieved for changing data")
)
