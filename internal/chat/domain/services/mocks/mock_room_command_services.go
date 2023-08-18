// Code generated by MockGen. DO NOT EDIT.
// Source: room_command_services.go

// Package mockchatdomainservices is a generated GoMock package.
package mockchatdomainservices

import (
	reflect "reflect"

	dto "github.com/iammuho/natternet/internal/chat/application/dto"
	values "github.com/iammuho/natternet/internal/chat/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomCommandDomainServices is a mock of RoomCommandDomainServices interface.
type MockRoomCommandDomainServices struct {
	ctrl     *gomock.Controller
	recorder *MockRoomCommandDomainServicesMockRecorder
}

// MockRoomCommandDomainServicesMockRecorder is the mock recorder for MockRoomCommandDomainServices.
type MockRoomCommandDomainServicesMockRecorder struct {
	mock *MockRoomCommandDomainServices
}

// NewMockRoomCommandDomainServices creates a new mock instance.
func NewMockRoomCommandDomainServices(ctrl *gomock.Controller) *MockRoomCommandDomainServices {
	mock := &MockRoomCommandDomainServices{ctrl: ctrl}
	mock.recorder = &MockRoomCommandDomainServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomCommandDomainServices) EXPECT() *MockRoomCommandDomainServicesMockRecorder {
	return m.recorder
}

// CreateRoom mocks base method.
func (m *MockRoomCommandDomainServices) CreateRoom(arg0 *dto.CreateRoomReqDTO) (*values.RoomValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", arg0)
	ret0, _ := ret[0].(*values.RoomValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockRoomCommandDomainServicesMockRecorder) CreateRoom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockRoomCommandDomainServices)(nil).CreateRoom), arg0)
}
