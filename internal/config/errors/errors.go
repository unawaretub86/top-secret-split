package errors

import "errors"

var (
	ErrNotEnoughSatellites = errors.New("Satellites are not enough")
	ErrInvalidSatellite    = errors.New("invalid satellite")
)
