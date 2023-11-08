package mocks_test

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
)

type MockTopSecretPort struct {
	ctrl     *gomock.Controller
	recorder *MockMockTopSecretPortMockRecorder
}

type MockMockTopSecretPortMockRecorder struct {
	mock *MockTopSecretPort
}

// crea una nueva instancia del mock.
// recibe un controlador de gomock como argumento y devuelve un mock.
func NewMockTopSecretPort(ctrl *gomock.Controller) *MockTopSecretPort {
	mock := &MockTopSecretPort{ctrl: ctrl}
	mock.recorder = &MockMockTopSecretPortMockRecorder{mock}
	return mock
}

// EXPECT() es para registrar llamadas a métodos en el mock.
func (m *MockTopSecretPort) EXPECT() *MockMockTopSecretPortMockRecorder {
	return m.recorder
}

// Recibe un argumento TopSecretSplit  y distances
// y devuelve una respuesta simulada.
func (m *MockTopSecretPort) TopSecretSplit(pathParameters map[string]string, body, requestId string) (*entities.Satellite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopSecretSplit", pathParameters, body, requestId)
	ret0, _ := ret[0].(*entities.Satellite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopSecretSplit  se utiliza para obtener llamadas al método TopSecretSplit en el mock.
func (mr *MockMockTopSecretPortMockRecorder) TopSecretSplit(pathParameters map[string]string, body, requestId string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopSecretSplit", reflect.TypeOf((*MockTopSecretPort)(nil).TopSecretSplit), pathParameters, body, requestId)
}

// Recibe un argumento requestID  y distances
// y devuelve una respuesta simulada.
func (m *MockTopSecretPort) GetPositionMessage(requestId string) (*entities.LocationMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionMessage", requestId)
	ret0, _ := ret[0].(*entities.LocationMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopSecretSplit  se utiliza para obtener llamadas al método GetPositionMessage en el mock.
func (mr *MockMockTopSecretPortMockRecorder) GetPositionMessage(requestId string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionMessage", reflect.TypeOf((*MockTopSecretPort)(nil).GetPositionMessage), requestId)
}
