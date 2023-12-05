// Code generated by MockGen. DO NOT EDIT.
// Source: user_inputport.go

// Package mock_inputport is a generated GoMock package.
package mock_inputport

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/shima004/chat-server/entities"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserUsecase) CreateUser(ctx context.Context, user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserUsecaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserUsecase)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserUsecase) DeleteUser(ctx context.Context, userID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserUsecaseMockRecorder) DeleteUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserUsecase)(nil).DeleteUser), ctx, userID)
}

// Login mocks base method.
func (m *MockUserUsecase) Login(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserUsecaseMockRecorder) Login(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserUsecase)(nil).Login), ctx, email, password)
}

// ReadUserByUserID mocks base method.
func (m *MockUserUsecase) ReadUserByUserID(ctx context.Context, userID uint) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserByUserID", ctx, userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserByUserID indicates an expected call of ReadUserByUserID.
func (mr *MockUserUsecaseMockRecorder) ReadUserByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserByUserID", reflect.TypeOf((*MockUserUsecase)(nil).ReadUserByUserID), ctx, userID)
}

// SubscribeChannel mocks base method.
func (m *MockUserUsecase) SubscribeChannel(ctx context.Context, userID, channelID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChannel", ctx, userID, channelID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeChannel indicates an expected call of SubscribeChannel.
func (mr *MockUserUsecaseMockRecorder) SubscribeChannel(ctx, userID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChannel", reflect.TypeOf((*MockUserUsecase)(nil).SubscribeChannel), ctx, userID, channelID)
}
