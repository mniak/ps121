// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/duplicomp (interfaces: Stream)

// Package duplicomp is a generated GoMock package.
package duplicomp

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Receive mocks base method.
func (m *MockStream) Receive() (protoreflect.ProtoMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(protoreflect.ProtoMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockStreamMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockStream)(nil).Receive))
}

// Send mocks base method.
func (m *MockStream) Send(arg0 protoreflect.ProtoMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockStreamMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockStream)(nil).Send), arg0)
}
