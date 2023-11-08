package request

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

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
func ItemsToSatellites(requestID, tableName string, output *dynamodb.BatchGetItemOutput) (*entities.Satellites, error) {
	var satellites entities.Satellites

	if data, found := output.Responses[tableName]; found {
		// satData es un []map[string]*dynamodb.AttributeValue.

		// Se itera sobre satData y se convierten los elementos a Satellite.
		for _, item := range data {
			var satellite entities.Satellite
			if err := dynamodbattribute.UnmarshalMap(item, &satellite); err != nil {
				fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", requestID, err)
				return nil, err
			}

			satellites.Satellite = append(satellites.Satellite, satellite)
		}
	}

	return &satellites, nil
}
