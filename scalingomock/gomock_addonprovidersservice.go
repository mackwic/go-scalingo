// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-scalingo (interfaces: AddonProvidersService)

// Package scalingomock is a generated GoMock package.
package scalingomock

import (
	go_scalingo "github.com/Scalingo/go-scalingo"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAddonProvidersService is a mock of AddonProvidersService interface
type MockAddonProvidersService struct {
	ctrl     *gomock.Controller
	recorder *MockAddonProvidersServiceMockRecorder
}

// MockAddonProvidersServiceMockRecorder is the mock recorder for MockAddonProvidersService
type MockAddonProvidersServiceMockRecorder struct {
	mock *MockAddonProvidersService
}

// NewMockAddonProvidersService creates a new mock instance
func NewMockAddonProvidersService(ctrl *gomock.Controller) *MockAddonProvidersService {
	mock := &MockAddonProvidersService{ctrl: ctrl}
	mock.recorder = &MockAddonProvidersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAddonProvidersService) EXPECT() *MockAddonProvidersServiceMockRecorder {
	return m.recorder
}

// AddonProviderPlansList mocks base method
func (m *MockAddonProvidersService) AddonProviderPlansList(arg0 string) ([]*go_scalingo.Plan, error) {
	ret := m.ctrl.Call(m, "AddonProviderPlansList", arg0)
	ret0, _ := ret[0].([]*go_scalingo.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProviderPlansList indicates an expected call of AddonProviderPlansList
func (mr *MockAddonProvidersServiceMockRecorder) AddonProviderPlansList(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProviderPlansList", reflect.TypeOf((*MockAddonProvidersService)(nil).AddonProviderPlansList), arg0)
}

// AddonProvidersList mocks base method
func (m *MockAddonProvidersService) AddonProvidersList() ([]*go_scalingo.AddonProvider, error) {
	ret := m.ctrl.Call(m, "AddonProvidersList")
	ret0, _ := ret[0].([]*go_scalingo.AddonProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddonProvidersList indicates an expected call of AddonProvidersList
func (mr *MockAddonProvidersServiceMockRecorder) AddonProvidersList() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddonProvidersList", reflect.TypeOf((*MockAddonProvidersService)(nil).AddonProvidersList))
}