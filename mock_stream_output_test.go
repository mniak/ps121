// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/ps121 (interfaces: OutputStream)

// Package ps121 is a generated GoMock package.
package ps121

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// MockOutputStream is a mock of OutputStream interface.
type MockOutputStream struct {
	ctrl     *gomock.Controller
	recorder *MockOutputStreamMockRecorder
}

// MockOutputStreamMockRecorder is the mock recorder for MockOutputStream.
type MockOutputStreamMockRecorder struct {
	mock *MockOutputStream
}

// NewMockOutputStream creates a new mock instance.
func NewMockOutputStream(ctrl *gomock.Controller) *MockOutputStream {
	mock := &MockOutputStream{ctrl: ctrl}
	mock.recorder = &MockOutputStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOutputStream) EXPECT() *MockOutputStreamMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockOutputStream) Send(arg0 protoreflect.ProtoMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockOutputStreamMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOutputStream)(nil).Send), arg0)
}
