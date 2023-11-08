package services

import (
	"fmt"

	"github.com/unawaretub86/top-secret-split/internal/config/errors"
	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
	"github.com/unawaretub86/top-secret-split/internal/domain/ports"
	"github.com/unawaretub86/top-secret-split/internal/domain/request"
)

var pathParameterName = "satellite-name"

type topSecretService struct {
	repository ports.RepositoryPort
	rest       ports.RepositoryRest
}

func NewTopSecretService(repository ports.RepositoryPort, rest ports.RepositoryRest) *topSecretService {
	return &topSecretService{
		repository: repository,
		rest:       rest,
	}
}

// servicio encargado de procesar informacion para guardar satellites
func (service *topSecretService) TopSecretSplit(pathParameters map[string]string, body, requestId string) (*entities.Satellite, error) {
	satellite := &entities.Satellite{}

	err := request.ConvertToStruct(satellite, body, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, errors.ErrInvalidSatellite)
		return nil, err
	}

	// obtenemos path param para guardar el satellite por nombre
	satellite.Name = pathParameters[pathParameterName]

	satellite.ValidateSatellite()

	return service.repository.CreateSatellite(requestId, satellite)
}

// servicio encargado de procesar informacion para obtener entities.locationMessage
func (service *topSecretService) GetPositionMessage(requestId string) (*entities.LocationMessage, error) {
	// obtenemos satellites desde DB
	satellites, err := service.repository.GetSatellites(requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

	// validamos que sean 3 los satellites para poder obtener locationMessage
	satellites.ValidateSatellites()

	// obtenemos locationMessage desde top-secret
	result, err := service.rest.GetLocationMessage(*satellites, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

	locationMessage := &entities.LocationMessage{}
	err = request.ConvertToStruct(locationMessage, string(result), requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

	return locationMessage, nil
}
