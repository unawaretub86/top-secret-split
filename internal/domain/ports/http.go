package ports

import "github.com/unawaretub86/top-secret-split/internal/domain/entities"

type RepositoryRest interface {
	GetLocationMessage(entities.Satellites, string) (*entities.LocationMessage, error)
}
