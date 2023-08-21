// Code generated by MockGen. DO NOT EDIT.
// Source: room_query_services.go

// Package mockchatdomainservices is a generated GoMock package.
package mockchatdomainservices

import (
	reflect "reflect"

	dto "github.com/iammuho/natternet/internal/chat/application/dto"
	values "github.com/iammuho/natternet/internal/chat/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomQueryDomainServices is a mock of RoomQueryDomainServices interface.
type MockRoomQueryDomainServices struct {
	ctrl     *gomock.Controller
	recorder *MockRoomQueryDomainServicesMockRecorder
}

// MockRoomQueryDomainServicesMockRecorder is the mock recorder for MockRoomQueryDomainServices.
type MockRoomQueryDomainServicesMockRecorder struct {
	mock *MockRoomQueryDomainServices
}

// NewMockRoomQueryDomainServices creates a new mock instance.
func NewMockRoomQueryDomainServices(ctrl *gomock.Controller) *MockRoomQueryDomainServices {
	mock := &MockRoomQueryDomainServices{ctrl: ctrl}
	mock.recorder = &MockRoomQueryDomainServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomQueryDomainServices) EXPECT() *MockRoomQueryDomainServicesMockRecorder {
	return m.recorder
}

// QueryRooms mocks base method.
func (m *MockRoomQueryDomainServices) QueryRooms(arg0 *dto.QueryRoomsReqDTO) ([]*values.RoomValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRooms", arg0)
	ret0, _ := ret[0].([]*values.RoomValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// QueryRooms indicates an expected call of QueryRooms.
func (mr *MockRoomQueryDomainServicesMockRecorder) QueryRooms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRooms", reflect.TypeOf((*MockRoomQueryDomainServices)(nil).QueryRooms), arg0)
}