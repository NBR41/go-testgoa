// Code generated by MockGen. DO NOT EDIT.
// Source: model.go

// Package local is a generated GoMock package.
package sql

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// Mockpassworder is a mock of passworder interface
type Mockpassworder struct {
	ctrl     *gomock.Controller
	recorder *MockpassworderMockRecorder
}

// MockpassworderMockRecorder is the mock recorder for Mockpassworder
type MockpassworderMockRecorder struct {
	mock *Mockpassworder
}

// NewMockpassworder creates a new mock instance
func NewMockpassworder(ctrl *gomock.Controller) *Mockpassworder {
	mock := &Mockpassworder{ctrl: ctrl}
	mock.recorder = &MockpassworderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockpassworder) EXPECT() *MockpassworderMockRecorder {
	return m.recorder
}

// CryptPassword mocks base method
func (m *Mockpassworder) CryptPassword(password string) ([]byte, []byte, error) {
	ret := m.ctrl.Call(m, "CryptPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CryptPassword indicates an expected call of CryptPassword
func (mr *MockpassworderMockRecorder) CryptPassword(password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CryptPassword", reflect.TypeOf((*Mockpassworder)(nil).CryptPassword), password)
}

// ComparePassword mocks base method
func (m *Mockpassworder) ComparePassword(password string, salt, hash []byte) (bool, error) {
	ret := m.ctrl.Call(m, "ComparePassword", password, salt, hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComparePassword indicates an expected call of ComparePassword
func (mr *MockpassworderMockRecorder) ComparePassword(password, salt, hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*Mockpassworder)(nil).ComparePassword), password, salt, hash)
}