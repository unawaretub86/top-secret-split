package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/unawaretub86/top-secret-split/internal/adapters/http"
	repositories "github.com/unawaretub86/top-secret-split/internal/adapters/repository"
	"github.com/unawaretub86/top-secret-split/internal/domain/request"
	"github.com/unawaretub86/top-secret-split/internal/domain/services"
)

const httpMethodGet = "GET"

var (
	httpRest   = http.NewRestHttp()
	repoDB     = repositories.NewDynamoDBInstance()
	topService = services.NewTopSecretService(repoDB, httpRest)
)

// HandleRequest maneja la solicitud.
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	requestID := lc.AwsRequestID

	// Si la solicitud es un método "GET", llama a la función handleGETRequest.
	if request.HTTPMethod == httpMethodGet {
		return handleGetRequest(requestID)
	}

	// Si no es un método "GET", llama a la función handlePOSTRequest y
	return handlePostRequest(requestID, request.Body, request.PathParameters)
}

// procesamos el caso en el que el method de la solicitud es GET
func handleGetRequest(requestID string) (*events.APIGatewayProxyResponse, error) {
	// Se llama al service para procesar la informacion y retornar la posicion y el mensaje
	bodyResponse, err := topService.GetPositionMessage(requestID)
	if err != nil {
		return handleError(err)
	}

	// convertimos a bytes para poder responder
	bytesResponse, err := request.ConvertToBytes(bodyResponse, requestID)
	if err != nil {
		return handleError(err)
	}

	// procesamos la respuesta en caso de ser correcta
	return createResponse(200, string(bytesResponse)), nil
}

// procesamos el caso en el que el method de la solicitud es POST
func handlePostRequest(requestID, body string, pathParameters map[string]string) (*events.APIGatewayProxyResponse, error) {
	// Se llama al service para procesar la informacion y guardar el satellite enviado
	bodyResponse, err := topService.TopSecretSplit(pathParameters, body, requestID)
	if err != nil {
		return handleError(err)
	}

	// convertimos a bytes para poder responder
	bytesResponse, err := request.ConvertToBytes(bodyResponse, requestID)
	if err != nil {
		return handleError(err)
	}

	// procesamos la respuesta en caso de ser correcta
	return createResponse(200, string(bytesResponse)), nil
}

// construimos la respuesta de tipo APIGatewayProxyResponse
func createResponse(statusCode int, body string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       body,
	}
}

// construimos la respuesta de tipo APIGatewayProxyResponse de tipo error en caso de que este exista
func handleError(err error) (*events.APIGatewayProxyResponse, error) {
	log.Printf("Error: %s", err.Error())
	return createResponse(404, fmt.Sprintf("Error: %s", err.Error())), nil
}
