// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gorkaio/gboy/pkg/gameboy (interfaces: CPU)

// Package gameboy_mock is a generated GoMock package.
package gameboy_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCPU is a mock of CPU interface
type MockCPU struct {
	ctrl     *gomock.Controller
	recorder *MockCPUMockRecorder
}

// MockCPUMockRecorder is the mock recorder for MockCPU
type MockCPUMockRecorder struct {
	mock *MockCPU
}

// NewMockCPU creates a new mock instance
func NewMockCPU(ctrl *gomock.Controller) *MockCPU {
	mock := &MockCPU{ctrl: ctrl}
	mock.recorder = &MockCPUMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCPU) EXPECT() *MockCPUMockRecorder {
	return m.recorder
}

// Step mocks base method
func (m *MockCPU) Step() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Step")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Step indicates an expected call of Step
func (mr *MockCPUMockRecorder) Step() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Step", reflect.TypeOf((*MockCPU)(nil).Step))
}