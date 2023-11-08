package entities

import "github.com/unawaretub86/top-secret-split/internal/config/errors"

type (
	Satellite struct {
		Name     string   `json:"name"`
		Distance *float32 `json:"distance"`
		Message  []string `json:"message"`
	}

	LocationMessage struct {
		X       float32 `json:"X"`
		Y       float32 `json:"Y"`
		Message string  `json:"message"`
	}

	Satellites struct {
		Satellite []Satellite `json:"satellites"`
	}
)

// Validate se encarga de validar que los parametros sean correctos
func (satellite *Satellite) ValidateSatellite() error {
	if satellite.Distance == nil || len(satellite.Message) == 0 {
		return errors.ErrInvalidSatellite
	}

	return nil
}
