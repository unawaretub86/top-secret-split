package request

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mitchellh/mapstructure"

	"github.com/unawaretub86/top-secret-split/internal/config/errors"
	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
)

// convierte a bytes para poder procesarla peticion
func ConvertToBytes(v any, requestId string) ([]byte, error) {
	payload, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", requestId, err)
		return nil, err
	}

	return payload, nil
}

// convierte de typo dynamo al struct entitites.Satellites
func ItemsToSatellites(requestID string, output *dynamodb.BatchGetItemOutput) (*entities.Satellites, error) {
	var satellites entities.Satellites

	err := mapstructure.Decode(output.Responses, &satellites)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, errors.ErrNotEnoughSatellites)
		return nil, err
	}

	return &satellites, nil
}
