// Code generated by MockGen. DO NOT EDIT.
// Source: finder.go
//
// Generated by this command:
//
//	mockgen -source=finder.go -destination=../mock/finder.mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockiFinder is a mock of iFinder interface.
type MockiFinder[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockiFinderMockRecorder[T]
}

// MockiFinderMockRecorder is the mock recorder for MockiFinder.
type MockiFinderMockRecorder[T any] struct {
	mock *MockiFinder[T]
}

// NewMockiFinder creates a new mock instance.
func NewMockiFinder[T any](ctrl *gomock.Controller) *MockiFinder[T] {
	mock := &MockiFinder[T]{ctrl: ctrl}
	mock.recorder = &MockiFinderMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockiFinder[T]) EXPECT() *MockiFinderMockRecorder[T] {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockiFinder[T]) FindAll(ctx context.Context) ([]*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockiFinderMockRecorder[T]) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockiFinder[T])(nil).FindAll), ctx)
}

// FindOne mocks base method.
func (m *MockiFinder[T]) FindOne(ctx context.Context) (*T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx)
	ret0, _ := ret[0].(*T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockiFinderMockRecorder[T]) FindOne(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockiFinder[T])(nil).FindOne), ctx)
}
