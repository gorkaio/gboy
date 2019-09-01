// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gorkaio/gboy/pkg/gameboy (interfaces: Memory)

// Package gameboy_mock is a generated GoMock package.
package gameboy_mock

import (
	gomock "github.com/golang/mock/gomock"
	memory "github.com/gorkaio/gboy/pkg/memory"
	reflect "reflect"
)

// MockMemory is a mock of Memory interface
type MockMemory struct {
	ctrl     *gomock.Controller
	recorder *MockMemoryMockRecorder
}

// MockMemoryMockRecorder is the mock recorder for MockMemory
type MockMemoryMockRecorder struct {
	mock *MockMemory
}

// NewMockMemory creates a new mock instance
func NewMockMemory(ctrl *gomock.Controller) *MockMemory {
	mock := &MockMemory{ctrl: ctrl}
	mock.recorder = &MockMemoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMemory) EXPECT() *MockMemoryMockRecorder {
	return m.recorder
}

// Eject mocks base method
func (m *MockMemory) Eject() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Eject")
}

// Eject indicates an expected call of Eject
func (mr *MockMemoryMockRecorder) Eject() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Eject", reflect.TypeOf((*MockMemory)(nil).Eject))
}

// Load mocks base method
func (m *MockMemory) Load(arg0 memory.Cart) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Load", arg0)
}

// Load indicates an expected call of Load
func (mr *MockMemoryMockRecorder) Load(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockMemory)(nil).Load), arg0)
}

// Read mocks base method
func (m *MockMemory) Read(arg0 uint16) byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(byte)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockMemoryMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockMemory)(nil).Read), arg0)
}

// Write mocks base method
func (m *MockMemory) Write(arg0 uint16, arg1 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0, arg1)
}

// Write indicates an expected call of Write
func (mr *MockMemoryMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockMemory)(nil).Write), arg0, arg1)
}
