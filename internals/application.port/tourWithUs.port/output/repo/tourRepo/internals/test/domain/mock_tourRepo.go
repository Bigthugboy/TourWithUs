// Code generated by MockGen. DO NOT EDIT.
// Source: tourRepo.go

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"
	time "time"

	tourDto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/tourDto"
	gomock "github.com/golang/mock/gomock"
)

// MockTourRepository is a mock of TourRepository interface.
type MockTourRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTourRepositoryMockRecorder
}

// MockTourRepositoryMockRecorder is the mock recorder for MockTourRepository.
type MockTourRepositoryMockRecorder struct {
	mock *MockTourRepository
}

// NewMockTourRepository creates a new mock instance.
func NewMockTourRepository(ctrl *gomock.Controller) *MockTourRepository {
	mock := &MockTourRepository{ctrl: ctrl}
	mock.recorder = &MockTourRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTourRepository) EXPECT() *MockTourRepositoryMockRecorder {
	return m.recorder
}

// CreateTour mocks base method.
func (m *MockTourRepository) CreateTour(object tourDto.TourObject) (tourDto.CreateTourResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTour", object)
	ret0, _ := ret[0].(tourDto.CreateTourResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTour indicates an expected call of CreateTour.
func (mr *MockTourRepositoryMockRecorder) CreateTour(object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTour", reflect.TypeOf((*MockTourRepository)(nil).CreateTour), object)
}

// DeleteTour mocks base method.
func (m *MockTourRepository) DeleteTour(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTour", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTour indicates an expected call of DeleteTour.
func (mr *MockTourRepositoryMockRecorder) DeleteTour(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTour", reflect.TypeOf((*MockTourRepository)(nil).DeleteTour), id)
}

// GetAllTours mocks base method.
func (m *MockTourRepository) GetAllTours() ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTours")
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTours indicates an expected call of GetAllTours.
func (mr *MockTourRepositoryMockRecorder) GetAllTours() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTours", reflect.TypeOf((*MockTourRepository)(nil).GetAllTours))
}

// GetAvailableTours mocks base method.
func (m *MockTourRepository) GetAvailableTours() ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableTours")
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableTours indicates an expected call of GetAvailableTours.
func (mr *MockTourRepositoryMockRecorder) GetAvailableTours() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableTours", reflect.TypeOf((*MockTourRepository)(nil).GetAvailableTours))
}

// GetTourById mocks base method.
func (m *MockTourRepository) GetTourById(id string) (tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTourById", id)
	ret0, _ := ret[0].(tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTourById indicates an expected call of GetTourById.
func (mr *MockTourRepositoryMockRecorder) GetTourById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTourById", reflect.TypeOf((*MockTourRepository)(nil).GetTourById), id)
}

// GetToursByDateRange mocks base method.
func (m *MockTourRepository) GetToursByDateRange(startDate, endDate time.Time) ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToursByDateRange", startDate, endDate)
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToursByDateRange indicates an expected call of GetToursByDateRange.
func (mr *MockTourRepositoryMockRecorder) GetToursByDateRange(startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToursByDateRange", reflect.TypeOf((*MockTourRepository)(nil).GetToursByDateRange), startDate, endDate)
}

// GetToursByLocation mocks base method.
func (m *MockTourRepository) GetToursByLocation(location string) ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToursByLocation", location)
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToursByLocation indicates an expected call of GetToursByLocation.
func (mr *MockTourRepositoryMockRecorder) GetToursByLocation(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToursByLocation", reflect.TypeOf((*MockTourRepository)(nil).GetToursByLocation), location)
}

// GetToursByPriceRange mocks base method.
func (m *MockTourRepository) GetToursByPriceRange(minPrice, maxPrice float64) ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToursByPriceRange", minPrice, maxPrice)
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToursByPriceRange indicates an expected call of GetToursByPriceRange.
func (mr *MockTourRepositoryMockRecorder) GetToursByPriceRange(minPrice, maxPrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToursByPriceRange", reflect.TypeOf((*MockTourRepository)(nil).GetToursByPriceRange), minPrice, maxPrice)
}

// GetToursByType mocks base method.
func (m *MockTourRepository) GetToursByType(tourType tourDto.TourType) ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToursByType", tourType)
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToursByType indicates an expected call of GetToursByType.
func (mr *MockTourRepositoryMockRecorder) GetToursByType(tourType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToursByType", reflect.TypeOf((*MockTourRepository)(nil).GetToursByType), tourType)
}

// SearchTours mocks base method.
func (m *MockTourRepository) SearchTours(query string) ([]tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchTours", query)
	ret0, _ := ret[0].([]tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchTours indicates an expected call of SearchTours.
func (mr *MockTourRepositoryMockRecorder) SearchTours(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchTours", reflect.TypeOf((*MockTourRepository)(nil).SearchTours), query)
}

// UpdateTour mocks base method.
func (m *MockTourRepository) UpdateTour(id string, updatedFields tourDto.UpdateTourDto) (tourDto.TourObject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTour", id, updatedFields)
	ret0, _ := ret[0].(tourDto.TourObject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTour indicates an expected call of UpdateTour.
func (mr *MockTourRepositoryMockRecorder) UpdateTour(id, updatedFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTour", reflect.TypeOf((*MockTourRepository)(nil).UpdateTour), id, updatedFields)
}