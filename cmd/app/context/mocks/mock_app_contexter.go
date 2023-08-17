// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package mockcontext is a generated GoMock package.
package mockcontext

import (
	context "context"
	reflect "reflect"

	jwt "github.com/iammuho/natternet/pkg/jwt"
	logger "github.com/iammuho/natternet/pkg/logger"
	mongodb "github.com/iammuho/natternet/pkg/mongodb"
	gomock "go.uber.org/mock/gomock"
)

// MockAppContext is a mock of AppContext interface.
type MockAppContext struct {
	ctrl     *gomock.Controller
	recorder *MockAppContextMockRecorder
}

// MockAppContextMockRecorder is the mock recorder for MockAppContext.
type MockAppContextMockRecorder struct {
	mock *MockAppContext
}

// NewMockAppContext creates a new mock instance.
func NewMockAppContext(ctrl *gomock.Controller) *MockAppContext {
	mock := &MockAppContext{ctrl: ctrl}
	mock.recorder = &MockAppContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppContext) EXPECT() *MockAppContextMockRecorder {
	return m.recorder
}

// GetContext mocks base method.
func (m *MockAppContext) GetContext() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockAppContextMockRecorder) GetContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockAppContext)(nil).GetContext))
}

// GetJwtContext mocks base method.
func (m *MockAppContext) GetJwtContext() jwt.JwtContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJwtContext")
	ret0, _ := ret[0].(jwt.JwtContext)
	return ret0
}

// GetJwtContext indicates an expected call of GetJwtContext.
func (mr *MockAppContextMockRecorder) GetJwtContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJwtContext", reflect.TypeOf((*MockAppContext)(nil).GetJwtContext))
}

// GetLogger mocks base method.
func (m *MockAppContext) GetLogger() *logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*logger.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockAppContextMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockAppContext)(nil).GetLogger))
}

// GetMongoContext mocks base method.
func (m *MockAppContext) GetMongoContext() mongodb.MongoDBContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMongoContext")
	ret0, _ := ret[0].(mongodb.MongoDBContext)
	return ret0
}

// GetMongoContext indicates an expected call of GetMongoContext.
func (mr *MockAppContextMockRecorder) GetMongoContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMongoContext", reflect.TypeOf((*MockAppContext)(nil).GetMongoContext))
}
