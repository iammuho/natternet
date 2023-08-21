// Code generated by MockGen. DO NOT EDIT.
// Source: room_repository.go

// Package mockchatrepository is a generated GoMock package.
package mockchatrepository

import (
	reflect "reflect"

	values "github.com/iammuho/natternet/internal/chat/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockRoomRepository is a mock of RoomRepository interface.
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository.
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance.
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRoomRepository) Create(room *values.RoomDBValue) *errorhandler.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", room)
	ret0, _ := ret[0].(*errorhandler.Response)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRoomRepositoryMockRecorder) Create(room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoomRepository)(nil).Create), room)
}