// Code generated by MockGen. DO NOT EDIT.
// Source: handler/handler.go

// Package mock_handler is a generated GoMock package.
package mock

import (
	reflect "reflect"
	model "se-school-case/pkg/model"

	gomock "github.com/golang/mock/gomock"
)

// MockRateInterface is a mock of RateInterface interface.
type MockRateInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRateInterfaceMockRecorder
}

// MockRateInterfaceMockRecorder is the mock recorder for MockRateInterface.
type MockRateInterfaceMockRecorder struct {
	mock *MockRateInterface
}

// NewMockRateInterface creates a new mock instance.
func NewMockRateInterface(ctrl *gomock.Controller) *MockRateInterface {
	mock := &MockRateInterface{ctrl: ctrl}
	mock.recorder = &MockRateInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateInterface) EXPECT() *MockRateInterfaceMockRecorder {
	return m.recorder
}

// GetRate mocks base method.
func (m *MockRateInterface) GetRate() (model.Rate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate")
	ret0, _ := ret[0].(model.Rate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockRateInterfaceMockRecorder) GetRate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockRateInterface)(nil).GetRate))
}

// SaveRate mocks base method.
func (m *MockRateInterface) SaveRate(currencyFrom, currencyTo string, exchangeRate float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveRate", currencyFrom, currencyTo, exchangeRate)
}

// SaveRate indicates an expected call of SaveRate.
func (mr *MockRateInterfaceMockRecorder) SaveRate(currencyFrom, currencyTo, exchangeRate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRate", reflect.TypeOf((*MockRateInterface)(nil).SaveRate), currencyFrom, currencyTo, exchangeRate)
}

// MockRateFetchInterface is a mock of RateFetchInterface interface.
type MockRateFetchInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRateFetchInterfaceMockRecorder
}

// MockRateFetchInterfaceMockRecorder is the mock recorder for MockRateFetchInterface.
type MockRateFetchInterfaceMockRecorder struct {
	mock *MockRateFetchInterface
}

// NewMockRateFetchInterface creates a new mock instance.
func NewMockRateFetchInterface(ctrl *gomock.Controller) *MockRateFetchInterface {
	mock := &MockRateFetchInterface{ctrl: ctrl}
	mock.recorder = &MockRateFetchInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRateFetchInterface) EXPECT() *MockRateFetchInterfaceMockRecorder {
	return m.recorder
}

// FetchExchangeRate mocks base method.
func (m *MockRateFetchInterface) FetchExchangeRate() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchExchangeRate")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchExchangeRate indicates an expected call of FetchExchangeRate.
func (mr *MockRateFetchInterfaceMockRecorder) FetchExchangeRate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchExchangeRate", reflect.TypeOf((*MockRateFetchInterface)(nil).FetchExchangeRate))
}
