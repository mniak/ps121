// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/duplicomp/internal/samples (interfaces: ServerHandler)

// Package tests is a generated GoMock package.
package tests

import (
	context "context"
	reflect "reflect"

	grpc "github.com/mniak/duplicomp/internal/samples/grpc"
	gomock "go.uber.org/mock/gomock"
)

// MockServerHandler is a mock of ServerHandler interface.
type MockServerHandler struct {
	ctrl     *gomock.Controller
	recorder *MockServerHandlerMockRecorder
}

// MockServerHandlerMockRecorder is the mock recorder for MockServerHandler.
type MockServerHandlerMockRecorder struct {
	mock *MockServerHandler
}

// NewMockServerHandler creates a new mock instance.
func NewMockServerHandler(ctrl *gomock.Controller) *MockServerHandler {
	mock := &MockServerHandler{ctrl: ctrl}
	mock.recorder = &MockServerHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerHandler) EXPECT() *MockServerHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockServerHandler) Handle(arg0 context.Context, arg1 *grpc.Ping) (*grpc.Pong, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", arg0, arg1)
	ret0, _ := ret[0].(*grpc.Pong)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockServerHandlerMockRecorder) Handle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockServerHandler)(nil).Handle), arg0, arg1)
}
