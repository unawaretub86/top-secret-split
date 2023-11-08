package mocks_test

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/unawaretub86/top-secret-split/internal/domain/entities"
)

type MockHttpPort struct {
	ctrl     *gomock.Controller
	recorder *MockHttpPortMockRecorder
}

type MockHttpPortMockRecorder struct {
	mock *MockHttpPort
}

// crea una nueva instancia del mock.
// recibe un controlador de gomock como argumento y devuelve un mock.
func NewMockHttpPort(ctrl *gomock.Controller) *MockHttpPort {
	mock := &MockHttpPort{ctrl: ctrl}
	mock.recorder = &MockHttpPortMockRecorder{mock}
	return mock
}

// EXPECT() es para registrar llamadas a métodos en el mock.
func (m *MockHttpPort) EXPECT() *MockHttpPortMockRecorder {
	return m.recorder
}

// Recibe un argumento GetLocationMessage  y distances
// y devuelve una respuesta simulada.
func (m *MockHttpPort) GetLocationMessage(satellites entities.Satellites, requestId string) (*entities.LocationMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationMessage", satellites, requestId)
	ret0, _ := ret[0].(*entities.LocationMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopSecretSplit  se utiliza para obtener llamadas al método TopSecretSplit en el mock.
func (mr *MockHttpPortMockRecorder) GetLocationMessage(satellites entities.Satellites, requestId string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationMessage", reflect.TypeOf((*MockHttpPort)(nil).GetLocationMessage), satellites, requestId)
}
