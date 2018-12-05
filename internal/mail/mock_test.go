// Code generated by MockGen. DO NOT EDIT.
// Source: mail.go

// Package mail is a generated GoMock package.
package mail

import (
	content "github.com/NBR41/go-testgoa/internal/mail/content"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockActionner is a mock of Actionner interface
type MockActionner struct {
	ctrl     *gomock.Controller
	recorder *MockActionnerMockRecorder
}

// MockActionnerMockRecorder is the mock recorder for MockActionner
type MockActionnerMockRecorder struct {
	mock *MockActionner
}

// NewMockActionner creates a new mock instance
func NewMockActionner(ctrl *gomock.Controller) *MockActionner {
	mock := &MockActionner{ctrl: ctrl}
	mock.recorder = &MockActionnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionner) EXPECT() *MockActionnerMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockActionner) Do(email string, cnt *content.Mail) error {
	ret := m.ctrl.Call(m, "Do", email, cnt)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockActionnerMockRecorder) Do(email, cnt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockActionner)(nil).Do), email, cnt)
}

// MockContenter is a mock of Contenter interface
type MockContenter struct {
	ctrl     *gomock.Controller
	recorder *MockContenterMockRecorder
}

// MockContenterMockRecorder is the mock recorder for MockContenter
type MockContenterMockRecorder struct {
	mock *MockContenter
}

// NewMockContenter creates a new mock instance
func NewMockContenter(ctrl *gomock.Controller) *MockContenter {
	mock := &MockContenter{ctrl: ctrl}
	mock.recorder = &MockContenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContenter) EXPECT() *MockContenterMockRecorder {
	return m.recorder
}

// GetResetPasswordMail mocks base method
func (m *MockContenter) GetResetPasswordMail(token string) *content.Mail {
	ret := m.ctrl.Call(m, "GetResetPasswordMail", token)
	ret0, _ := ret[0].(*content.Mail)
	return ret0
}

// GetResetPasswordMail indicates an expected call of GetResetPasswordMail
func (mr *MockContenterMockRecorder) GetResetPasswordMail(token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResetPasswordMail", reflect.TypeOf((*MockContenter)(nil).GetResetPasswordMail), token)
}

// GetPasswordUpdatedMail mocks base method
func (m *MockContenter) GetPasswordUpdatedMail() *content.Mail {
	ret := m.ctrl.Call(m, "GetPasswordUpdatedMail")
	ret0, _ := ret[0].(*content.Mail)
	return ret0
}

// GetPasswordUpdatedMail indicates an expected call of GetPasswordUpdatedMail
func (mr *MockContenterMockRecorder) GetPasswordUpdatedMail() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordUpdatedMail", reflect.TypeOf((*MockContenter)(nil).GetPasswordUpdatedMail))
}

// GetNewUserMail mocks base method
func (m *MockContenter) GetNewUserMail(nickname, token string) *content.Mail {
	ret := m.ctrl.Call(m, "GetNewUserMail", nickname, token)
	ret0, _ := ret[0].(*content.Mail)
	return ret0
}

// GetNewUserMail indicates an expected call of GetNewUserMail
func (mr *MockContenterMockRecorder) GetNewUserMail(nickname, token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewUserMail", reflect.TypeOf((*MockContenter)(nil).GetNewUserMail), nickname, token)
}

// GetActivationMail mocks base method
func (m *MockContenter) GetActivationMail(nickname, token string) *content.Mail {
	ret := m.ctrl.Call(m, "GetActivationMail", nickname, token)
	ret0, _ := ret[0].(*content.Mail)
	return ret0
}

// GetActivationMail indicates an expected call of GetActivationMail
func (mr *MockContenterMockRecorder) GetActivationMail(nickname, token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationMail", reflect.TypeOf((*MockContenter)(nil).GetActivationMail), nickname, token)
}

// GetUserActivatedMail mocks base method
func (m *MockContenter) GetUserActivatedMail() *content.Mail {
	ret := m.ctrl.Call(m, "GetUserActivatedMail")
	ret0, _ := ret[0].(*content.Mail)
	return ret0
}

// GetUserActivatedMail indicates an expected call of GetUserActivatedMail
func (mr *MockContenterMockRecorder) GetUserActivatedMail() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserActivatedMail", reflect.TypeOf((*MockContenter)(nil).GetUserActivatedMail))
}