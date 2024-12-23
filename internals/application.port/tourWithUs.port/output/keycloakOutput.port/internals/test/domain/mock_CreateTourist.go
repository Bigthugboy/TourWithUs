// Code generated by MockGen. DO NOT EDIT.
// Source: CreateTourist.go

// Package domain is a generated GoMock package.
package domain

import (
	dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKeycloakOutPutPort is a mock of KeycloakOutPutPort interface.
type MockKeycloakOutPutPort struct {
	ctrl     *gomock.Controller
	recorder *MockKeycloakOutPutPortMockRecorder
}

// MockKeycloakOutPutPortMockRecorder is the mock recorder for MockKeycloakOutPutPort.
type MockKeycloakOutPutPortMockRecorder struct {
	mock *MockKeycloakOutPutPort
}

// NewMockKeycloakOutPutPort creates a new mock instance.
func NewMockKeycloakOutPutPort(ctrl *gomock.Controller) *MockKeycloakOutPutPort {
	mock := &MockKeycloakOutPutPort{ctrl: ctrl}
	mock.recorder = &MockKeycloakOutPutPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeycloakOutPutPort) EXPECT() *MockKeycloakOutPutPortMockRecorder {
	return m.recorder
}

// RetrieveTourist mocks base method.
func (m *MockKeycloakOutPutPort) RetrieveTourist(details dto.RetrieveTourist) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrieveTourist", details)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrieveTourist indicates an expected call of RetrieveTourist.
func (mr *MockKeycloakOutPutPortMockRecorder) RetrieveTourist(details interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveTourist", reflect.TypeOf((*MockKeycloakOutPutPort)(nil).RetrieveTourist), details)
}

// SaveTourist mocks base method.
func (m *MockKeycloakOutPutPort) SaveTourist(tourist *dto.TouristDetails) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTourist", tourist)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTourist indicates an expected call of SaveTourist.
func (mr *MockKeycloakOutPutPortMockRecorder) SaveTourist(tourist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTourist", reflect.TypeOf((*MockKeycloakOutPutPort)(nil).SaveTourist), tourist)
}
