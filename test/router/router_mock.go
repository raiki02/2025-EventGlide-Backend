// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/router/router.go

// package tests is a generated GoMock package.
package tests

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRouterHdl is a mock of RouterHdl interface.
type MockRouterHdl struct {
	ctrl     *gomock.Controller
	recorder *MockRouterHdlMockRecorder
}

// MockRouterHdlMockRecorder is the mock recorder for MockRouterHdl.
type MockRouterHdlMockRecorder struct {
	mock *MockRouterHdl
}

// NewMockRouterHdl creates a new mock instance.
func NewMockRouterHdl(ctrl *gomock.Controller) *MockRouterHdl {
	mock := &MockRouterHdl{ctrl: ctrl}
	mock.recorder = &MockRouterHdlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterHdl) EXPECT() *MockRouterHdlMockRecorder {
	return m.recorder
}

// RegisterRouters mocks base method.
func (m *MockRouterHdl) RegisterRouters() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRouters")
}

// RegisterRouters indicates an expected call of RegisterRouters.
func (mr *MockRouterHdlMockRecorder) RegisterRouters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRouters", reflect.TypeOf((*MockRouterHdl)(nil).RegisterRouters))
}

// Run mocks base method.
func (m *MockRouterHdl) Run() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run")
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockRouterHdlMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockRouterHdl)(nil).Run))
}
