package services

import (
	"encoding/json"
	"fmt"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
	"github.com/unawaretub86/top-secret-split/internal/domain/ports"
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

	err := json.Unmarshal([]byte(body), &satellite)
	if err != nil {
		return nil, fmt.Errorf("[RequestId: %s][Error Unmarshaling API Gateway request: %v]", requestId, err)
	}

	// obtenemos path param para guardar el satellite por nombre
	satellite.Name = pathParameters[pathParameterName]

	err = satellite.ValidateSatellite()
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

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
	err = satellites.ValidateSatellites()
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

	// obtenemos locationMessage desde top-secret
	result, err := service.rest.GetLocationMessage(*satellites, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestId, err)
		return nil, err
	}

	return result, nil
}
