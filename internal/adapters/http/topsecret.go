package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
	"github.com/unawaretub86/top-secret-split/internal/domain/request"
)

var contentType = "application/json"

type RestHttp struct {
	topSecretUrl string
	client       *http.Client
}

func NewRestHttp() *RestHttp {
	topSecretUrl := os.Getenv("TOP_SECRET_URL")

	return &RestHttp{
		client:       &http.Client{},
		topSecretUrl: topSecretUrl,
	}
}

// Se encarga de hacer una peticion POST a la lambda top-secret con los satellites previamente almacenados
func (r *RestHttp) GetLocationMessage(satellites entities.Satellites, requestId string) (*entities.LocationMessage, error) {
	// convertimos a bytes para poder realizar la peticion
	payload, err := request.ConvertToBytes(satellites, requestId)
	if err != nil {
		fmt.Printf("[RequestId: %s][Error marshaling API Gateway request: %v]", requestId, err)
		return nil, err
	}

	// realizamos la peticion
	resp, err := http.Post(r.topSecretUrl, contentType, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("[RequestId: %s], [Error: %v]", requestId, err)
		return nil, err
	}

	defer resp.Body.Close()

	// leemos la respuesta para poder retornarla de tipo []byte
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[RequestId: %s], [Error: %v]", requestId, err)
		return nil, err
	}

	locationMessage := &entities.LocationMessage{}
	err = json.Unmarshal(bytes, &locationMessage)
	if err != nil {
		return nil, fmt.Errorf("[RequestId: %s][Error Unmarshaling API Gateway request: %v]", requestId, err)
	}

	return locationMessage, nil
}
