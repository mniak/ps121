// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/ps121 (interfaces: InputStream)

// Package ps121 is a generated GoMock package.
package ps121

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// MockInputStream is a mock of InputStream interface.
type MockInputStream struct {
	ctrl     *gomock.Controller
	recorder *MockInputStreamMockRecorder
}

// MockInputStreamMockRecorder is the mock recorder for MockInputStream.
type MockInputStreamMockRecorder struct {
	mock *MockInputStream
}

// NewMockInputStream creates a new mock instance.
func NewMockInputStream(ctrl *gomock.Controller) *MockInputStream {
	mock := &MockInputStream{ctrl: ctrl}
	mock.recorder = &MockInputStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInputStream) EXPECT() *MockInputStreamMockRecorder {
	return m.recorder
}

// Receive mocks base method.
func (m *MockInputStream) Receive() (protoreflect.ProtoMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(protoreflect.ProtoMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockInputStreamMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockInputStream)(nil).Receive))
}
