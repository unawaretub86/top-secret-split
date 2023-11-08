package test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
	"github.com/unawaretub86/top-secret-split/internal/domain/services"
	mocks_test "github.com/unawaretub86/top-secret-split/test/mocks"
)

type mocks struct {
	topSecret *mocks_test.MockTopSecretPort
	repoDB    *mocks_test.MockRepositoryPort
	repoRest  *mocks_test.MockHttpPort
}

func TestTopSecretOk(t *testing.T) {
	m := mocks{
		topSecret: mocks_test.NewMockTopSecretPort(gomock.NewController(t)),
		repoDB:    mocks_test.NewMockRepositoryPort(gomock.NewController(t)),
		repoRest:  mocks_test.NewMockHttpPort(gomock.NewController(t)),
	}

	topService := services.NewTopSecretService(m.repoDB, m.repoRest)

	requestID := "123abd"
	var d1 float32 = 600.0
	var d2 float32 = 500.0
	var d3 float32 = 716.80

	var r1 float32 = -185.12361
	var r2 float32 = 310.74167

	satellites := entities.Satellites{
		Satellite: []entities.Satellite{
			{
				Name:     "kenobi",
				Distance: &d1,
				Message:  []string{"este", "", "", "mensaje", ""},
			},
			{
				Name:     "skywalker",
				Distance: &d2,
				Message:  []string{"", "es", "", "", "secreto"},
			},
			{
				Name:     "sato",
				Distance: &d3,
				Message:  []string{"este", "", "un", "", ""},
			},
		},
	}

	satellite := entities.Satellite{
		Distance: &d1,
		Message:  []string{"message1"},
	}

	locationMessage := &entities.LocationMessage{
		X:       r1,
		Y:       r2,
		Message: "este es un mensaje secreto",
	}

	body := `{"distance": 600.0, "message": ["message1"]}`

	pathParameters := map[string]string{"satellite-name": "kenobi"}

	m.repoDB.EXPECT().CreateSatellite(requestID, &satellite).Return(&satellite, nil)
	m.repoDB.EXPECT().GetSatellites(requestID).Return(&satellites, nil)

	m.repoRest.EXPECT().GetLocationMessage(satellites, requestID).Return(locationMessage, nil)

	responsePositionMessage, err1 := topService.GetPositionMessage(requestID)
	responseTopSecretSplit, err2 := topService.TopSecretSplit(pathParameters, body, requestID)

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, responsePositionMessage, &entities.LocationMessage{
		X:       r1,
		Y:       r2,
		Message: "este es un mensaje secreto",
	})

	assert.Equal(t, responseTopSecretSplit, &entities.Satellite{
		Distance: &d1,
		Message:  []string{"message1"},
	})
}
