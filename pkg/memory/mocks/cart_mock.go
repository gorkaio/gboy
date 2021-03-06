// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gorkaio/gboy/pkg/memory (interfaces: Cart)

// Package memory_mock is a generated GoMock package.
package memory_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCart is a mock of Cart interface
type MockCart struct {
	ctrl     *gomock.Controller
	recorder *MockCartMockRecorder
}

// MockCartMockRecorder is the mock recorder for MockCart
type MockCartMockRecorder struct {
	mock *MockCart
}

// NewMockCart creates a new mock instance
func NewMockCart(ctrl *gomock.Controller) *MockCart {
	mock := &MockCart{ctrl: ctrl}
	mock.recorder = &MockCartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCart) EXPECT() *MockCartMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockCart) Read(arg0 uint16) byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(byte)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockCartMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockCart)(nil).Read), arg0)
}

// Write mocks base method
func (m *MockCart) Write(arg0 uint16, arg1 byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Write", arg0, arg1)
}

// Write indicates an expected call of Write
func (mr *MockCartMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCart)(nil).Write), arg0, arg1)
}
