package ports

import "github.com/unawaretub86/top-secret-split/internal/domain/entities"

type RepositoryPort interface {
	GetSatellites(string) (*entities.Satellites, error)
	CreateSatellite(string, *entities.Satellite) (*entities.Satellite, error)
}
