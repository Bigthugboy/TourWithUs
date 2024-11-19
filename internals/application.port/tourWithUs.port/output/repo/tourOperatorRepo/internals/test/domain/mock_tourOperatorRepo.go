// Code generated by MockGen. DO NOT EDIT.
// Source: tourOperatorRepo.go

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	operatorDto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
	gomock "github.com/golang/mock/gomock"
)

// MockTourOperatorRepo is a mock of TourOperatorRepo interface.
type MockTourOperatorRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTourOperatorRepoMockRecorder
}

// MockTourOperatorRepoMockRecorder is the mock recorder for MockTourOperatorRepo.
type MockTourOperatorRepoMockRecorder struct {
	mock *MockTourOperatorRepo
}

// NewMockTourOperatorRepo creates a new mock instance.
func NewMockTourOperatorRepo(ctrl *gomock.Controller) *MockTourOperatorRepo {
	mock := &MockTourOperatorRepo{ctrl: ctrl}
	mock.recorder = &MockTourOperatorRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTourOperatorRepo) EXPECT() *MockTourOperatorRepoMockRecorder {
	return m.recorder
}

// DeleteTourOperator mocks base method.
func (m *MockTourOperatorRepo) DeleteTourOperator(id string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTourOperator", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTourOperator indicates an expected call of DeleteTourOperator.
func (mr *MockTourOperatorRepoMockRecorder) DeleteTourOperator(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTourOperator", reflect.TypeOf((*MockTourOperatorRepo)(nil).DeleteTourOperator), id)
}

// GetAllTourOperator mocks base method.
func (m *MockTourOperatorRepo) GetAllTourOperator() ([]operatorDto.OperatorDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTourOperator")
	ret0, _ := ret[0].([]operatorDto.OperatorDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTourOperator indicates an expected call of GetAllTourOperator.
func (mr *MockTourOperatorRepoMockRecorder) GetAllTourOperator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTourOperator", reflect.TypeOf((*MockTourOperatorRepo)(nil).GetAllTourOperator))
}

// GetTourOperatorByEmail mocks base method.
func (m *MockTourOperatorRepo) GetTourOperatorByEmail(email string) (operatorDto.OperatorDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTourOperatorByEmail", email)
	ret0, _ := ret[0].(operatorDto.OperatorDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTourOperatorByEmail indicates an expected call of GetTourOperatorByEmail.
func (mr *MockTourOperatorRepoMockRecorder) GetTourOperatorByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTourOperatorByEmail", reflect.TypeOf((*MockTourOperatorRepo)(nil).GetTourOperatorByEmail), email)
}

// GetTourOperatorById mocks base method.
func (m *MockTourOperatorRepo) GetTourOperatorById(id string) (operatorDto.OperatorDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTourOperatorById", id)
	ret0, _ := ret[0].(operatorDto.OperatorDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTourOperatorById indicates an expected call of GetTourOperatorById.
func (mr *MockTourOperatorRepoMockRecorder) GetTourOperatorById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTourOperatorById", reflect.TypeOf((*MockTourOperatorRepo)(nil).GetTourOperatorById), id)
}

// GetTourOperatorByRating mocks base method.
func (m *MockTourOperatorRepo) GetTourOperatorByRating(limit int) ([]operatorDto.OperatorDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTourOperatorByRating", limit)
	ret0, _ := ret[0].([]operatorDto.OperatorDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTourOperatorByRating indicates an expected call of GetTourOperatorByRating.
func (mr *MockTourOperatorRepoMockRecorder) GetTourOperatorByRating(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTourOperatorByRating", reflect.TypeOf((*MockTourOperatorRepo)(nil).GetTourOperatorByRating), limit)
}

// SaveTourOperator mocks base method.
func (m *MockTourOperatorRepo) SaveTourOperator(object operatorDto.OperatorDto) (operatorDto.SavedOperatorRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTourOperator", object)
	ret0, _ := ret[0].(operatorDto.SavedOperatorRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTourOperator indicates an expected call of SaveTourOperator.
func (mr *MockTourOperatorRepoMockRecorder) SaveTourOperator(object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTourOperator", reflect.TypeOf((*MockTourOperatorRepo)(nil).SaveTourOperator), object)
}