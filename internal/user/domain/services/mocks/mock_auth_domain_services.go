// Code generated by MockGen. DO NOT EDIT.
// Source: auth_services.go

// Package mockuserdomainservices is a generated GoMock package.
package mockuserdomainservices

import (
	reflect "reflect"

	dto "github.com/iammuho/natternet/internal/user/application/auth/dto"
	values "github.com/iammuho/natternet/internal/user/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthDomainServices is a mock of AuthDomainServices interface.
type MockAuthDomainServices struct {
	ctrl     *gomock.Controller
	recorder *MockAuthDomainServicesMockRecorder
}

// MockAuthDomainServicesMockRecorder is the mock recorder for MockAuthDomainServices.
type MockAuthDomainServicesMockRecorder struct {
	mock *MockAuthDomainServices
}

// NewMockAuthDomainServices creates a new mock instance.
func NewMockAuthDomainServices(ctrl *gomock.Controller) *MockAuthDomainServices {
	mock := &MockAuthDomainServices{ctrl: ctrl}
	mock.recorder = &MockAuthDomainServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthDomainServices) EXPECT() *MockAuthDomainServicesMockRecorder {
	return m.recorder
}

// SignIn mocks base method.
func (m *MockAuthDomainServices) SignIn(req *dto.SignInReqDTO) (*values.UserValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", req)
	ret0, _ := ret[0].(*values.UserValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthDomainServicesMockRecorder) SignIn(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthDomainServices)(nil).SignIn), req)
}

// SignUp mocks base method.
func (m *MockAuthDomainServices) SignUp(req *dto.SignupReqDTO) (*values.UserValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", req)
	ret0, _ := ret[0].(*values.UserValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthDomainServicesMockRecorder) SignUp(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthDomainServices)(nil).SignUp), req)
}
