package mocks_test

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
)

type MockRepositoryPort struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryPortMockRecorder
}

type MockRepositoryPortMockRecorder struct {
	mock *MockRepositoryPort
}

// crea una nueva instancia del mock.
// recibe un controlador de gomock como argumento y devuelve un mock.
func NewMockRepositoryPort(ctrl *gomock.Controller) *MockRepositoryPort {
	mock := &MockRepositoryPort{ctrl: ctrl}
	mock.recorder = &MockRepositoryPortMockRecorder{mock}
	return mock
}

// EXPECT() es para registrar llamadas a métodos en el mock.
func (m *MockRepositoryPort) EXPECT() *MockRepositoryPortMockRecorder {
	return m.recorder
}

// Recibe un argumento CreateSatellite  y distances
// y devuelve una respuesta simulada.
func (m *MockRepositoryPort) CreateSatellite(requestID string, satellite *entities.Satellite) (*entities.Satellite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSatellite", requestID, satellite)
	ret0, _ := ret[0].(*entities.Satellite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopSecretSplit  se utiliza para obtener llamadas al método TopSecretSplit en el mock.
func (mr *MockRepositoryPortMockRecorder) CreateSatellite(requestID string, satellite *entities.Satellite) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSatellite", reflect.TypeOf((*MockRepositoryPort)(nil).CreateSatellite), requestID, satellite)
}

// Recibe un argumento requestID  y distances
// y devuelve una respuesta simulada.
func (m *MockRepositoryPort) GetSatellites(requestId string) (*entities.Satellites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSatellites", requestId)
	ret0, _ := ret[0].(*entities.Satellites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopSecretSplit  se utiliza para obtener llamadas al método GetPositionMessage en el mock.
func (mr *MockRepositoryPortMockRecorder) GetSatellites(requestId string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSatellites", reflect.TypeOf((*MockRepositoryPort)(nil).GetSatellites), requestId)
}
