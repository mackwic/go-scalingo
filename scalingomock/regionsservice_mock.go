// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: RegionsService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	reflect "reflect"

	scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
)

// MockRegionsService is a mock of RegionsService interface
type MockRegionsService struct {
	ctrl     *gomock.Controller
	recorder *MockRegionsServiceMockRecorder
}

// MockRegionsServiceMockRecorder is the mock recorder for MockRegionsService
type MockRegionsServiceMockRecorder struct {
	mock *MockRegionsService
}

// NewMockRegionsService creates a new mock instance
func NewMockRegionsService(ctrl *gomock.Controller) *MockRegionsService {
	mock := &MockRegionsService{ctrl: ctrl}
	mock.recorder = &MockRegionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegionsService) EXPECT() *MockRegionsServiceMockRecorder {
	return m.recorder
}

// RegionsList mocks base method
func (m *MockRegionsService) RegionsList() ([]scalingo.Region, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegionsList")
	ret0, _ := ret[0].([]scalingo.Region)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegionsList indicates an expected call of RegionsList
func (mr *MockRegionsServiceMockRecorder) RegionsList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegionsList", reflect.TypeOf((*MockRegionsService)(nil).RegionsList))
}
