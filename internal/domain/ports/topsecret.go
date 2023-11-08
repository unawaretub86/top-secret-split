package ports

import "github.com/unawaretub86/top-secret-split/internal/domain/entities"

type TopSecretPort interface {
	TopSecretSplit(string, *entities.Satellite) (*entities.Satellite, error)
}
