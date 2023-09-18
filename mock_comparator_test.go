// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mniak/duplicomp (interfaces: Comparator)

// Package duplicomp is a generated GoMock package.
package duplicomp

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

// MockComparator is a mock of Comparator interface.
type MockComparator struct {
	ctrl     *gomock.Controller
	recorder *MockComparatorMockRecorder
}

// MockComparatorMockRecorder is the mock recorder for MockComparator.
type MockComparatorMockRecorder struct {
	mock *MockComparator
}

// NewMockComparator creates a new mock instance.
func NewMockComparator(ctrl *gomock.Controller) *MockComparator {
	mock := &MockComparator{ctrl: ctrl}
	mock.recorder = &MockComparatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComparator) EXPECT() *MockComparatorMockRecorder {
	return m.recorder
}

// Compare mocks base method.
func (m *MockComparator) Compare(arg0 protoreflect.ProtoMessage, arg1 error, arg2 protoreflect.ProtoMessage, arg3 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Compare", arg0, arg1, arg2, arg3)
}

// Compare indicates an expected call of Compare.
func (mr *MockComparatorMockRecorder) Compare(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockComparator)(nil).Compare), arg0, arg1, arg2, arg3)
}
