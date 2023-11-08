package repositories

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/unawaretub86/top-secret-split/internal/config/errors"
	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
	"github.com/unawaretub86/top-secret-split/internal/domain/request"
)

const tableName = "satellites"

type dynamoDBSatelliteRepository struct {
	dynamodb dynamodbiface.DynamoDBAPI
}

func NewDynamoDBInstance() *dynamoDBSatelliteRepository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	dynamodbClient := dynamodb.New(sess)

	return &dynamoDBSatelliteRepository{
		dynamodb: dynamodbClient,
	}
}

// Se encarga de guardar en dynamo db el satellite que es enviado mediante la solicitud
func (d *dynamoDBSatelliteRepository) CreateSatellite(requestID string, satellite *entities.Satellite) (*entities.Satellite, error) {
	// convertirmos a tipo de dato map[string]*dynamodb.AttributeValue, para poder almacenar en dynamoDB
	av, err := dynamodbattribute.MarshalMap(satellite)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, err)
		return nil, err
	}

	// construimos el item a guardar
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	// guardamos item
	output, err := d.dynamodb.PutItem(input)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, err)
		return nil, err
	}

	satelliteData := &entities.Satellite{}
	err = dynamodbattribute.UnmarshalMap(output.Attributes, satelliteData)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, err)
		return nil, err
	}

	return satellite, nil
}

// obtenemos los satellites previamente guardados en DB
func (d *dynamoDBSatelliteRepository) GetSatellites(requestID string) (*entities.Satellites, error) {
	sato := "Sato"
	skywalker := "Skywalker"
	kenobi := "Kenobi"

	// construimos input para realizar un batch y traer los 3 satellites necesarios
	input := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			tableName: {
				Keys: []map[string]*dynamodb.AttributeValue{
					{
						"name": &dynamodb.AttributeValue{
							S: aws.String(sato),
						},
					}, {
						"name": &dynamodb.AttributeValue{
							S: aws.String(skywalker),
						},
					}, {
						"name": &dynamodb.AttributeValue{
							S: aws.String(kenobi),
						},
					},
				},
			},
		},
	}

	// realizamos batch para obtener la informacion de los satellites
	output, err := d.dynamodb.BatchGetItem(input)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, err)
		return nil, err
	}

	if len(output.Responses[tableName]) < 3 {
		fmt.Printf("[RequestId: %s][Error: %v]", requestID, errors.ErrNotEnoughSatellites)
		return nil, err
	}

	return request.ItemsToSatellites(requestID, output)
}
