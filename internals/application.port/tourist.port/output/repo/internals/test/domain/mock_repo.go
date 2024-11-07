// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockDBStore is a mock of DBStore interface.
type MockDBStore struct {
	ctrl     *gomock.Controller
	recorder *MockDBStoreMockRecorder
}

// MockDBStoreMockRecorder is the mock recorder for MockDBStore.
type MockDBStoreMockRecorder struct {
	mock *MockDBStore
}

// NewMockDBStore creates a new mock instance.
func NewMockDBStore(ctrl *gomock.Controller) *MockDBStore {
	mock := &MockDBStore{ctrl: ctrl}
	mock.recorder = &MockDBStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBStore) EXPECT() *MockDBStoreMockRecorder {
	return m.recorder
}

// DeleteTouristByEmail mocks base method.
func (m *MockDBStore) DeleteTouristByEmail(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTouristByEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTouristByEmail indicates an expected call of DeleteTouristByEmail.
func (mr *MockDBStoreMockRecorder) DeleteTouristByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTouristByEmail", reflect.TypeOf((*MockDBStore)(nil).DeleteTouristByEmail), email)
}

// DeleteTouristByID mocks base method.
func (m *MockDBStore) DeleteTouristByID(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTouristByID", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTouristByID indicates an expected call of DeleteTouristByID.
func (mr *MockDBStoreMockRecorder) DeleteTouristByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTouristByID", reflect.TypeOf((*MockDBStore)(nil).DeleteTouristByID), userID)
}

// GetAllTourists mocks base method.
func (m *MockDBStore) GetAllTourists() ([]dto.TouristObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTourists")
	ret0, _ := ret[0].([]dto.TouristObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTourists indicates an expected call of GetAllTourists.
func (mr *MockDBStoreMockRecorder) GetAllTourists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTourists", reflect.TypeOf((*MockDBStore)(nil).GetAllTourists))
}

// GetTouristByID mocks base method.
func (m *MockDBStore) GetTouristByID(userID string) (dto.TouristObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTouristByID", userID)
	ret0, _ := ret[0].(dto.TouristObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTouristByID indicates an expected call of GetTouristByID.
func (mr *MockDBStoreMockRecorder) GetTouristByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTouristByID", reflect.TypeOf((*MockDBStore)(nil).GetTouristByID), userID)
}

// InsertTourist mocks base method.
func (m *MockDBStore) InsertTourist(user dto.TouristObject) (*dto.TouristObject, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTourist", user)
	ret0, _ := ret[0].(*dto.TouristObject)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// InsertTourist indicates an expected call of InsertTourist.
func (mr *MockDBStoreMockRecorder) InsertTourist(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTourist", reflect.TypeOf((*MockDBStore)(nil).InsertTourist), user)
}

// SearchTouristByEmail mocks base method.
func (m *MockDBStore) SearchTouristByEmail(email string) (dto.TouristObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTouristByEmail", email)
	ret0, _ := ret[0].(dto.TouristObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTouristByEmail indicates an expected call of SearchTouristByEmail.
func (mr *MockDBStoreMockRecorder) SearchTouristByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTouristByEmail", reflect.TypeOf((*MockDBStore)(nil).SearchTouristByEmail), email)
}