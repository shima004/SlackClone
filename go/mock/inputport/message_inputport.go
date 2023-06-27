// Code generated by MockGen. DO NOT EDIT.
// Source: message_inputport.go

// Package mock_inputport is a generated GoMock package.
package mock_inputport

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/shima004/chat-server/entities"
)

// MockMessageUsecase is a mock of MessageUsecase interface.
type MockMessageUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMessageUsecaseMockRecorder
}

// MockMessageUsecaseMockRecorder is the mock recorder for MockMessageUsecase.
type MockMessageUsecaseMockRecorder struct {
	mock *MockMessageUsecase
}

// NewMockMessageUsecase creates a new mock instance.
func NewMockMessageUsecase(ctrl *gomock.Controller) *MockMessageUsecase {
	mock := &MockMessageUsecase{ctrl: ctrl}
	mock.recorder = &MockMessageUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageUsecase) EXPECT() *MockMessageUsecaseMockRecorder {
	return m.recorder
}

// DeleteMessage mocks base method.
func (m *MockMessageUsecase) DeleteMessage(ctx context.Context, messageID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", ctx, messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage.
func (mr *MockMessageUsecaseMockRecorder) DeleteMessage(ctx, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockMessageUsecase)(nil).DeleteMessage), ctx, messageID)
}

// FetchMessages mocks base method.
func (m *MockMessageUsecase) FetchMessages(ctx context.Context, channelID uint, limit, offset int) ([]*entities.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMessages", ctx, channelID, limit, offset)
	ret0, _ := ret[0].([]*entities.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessages indicates an expected call of FetchMessages.
func (mr *MockMessageUsecaseMockRecorder) FetchMessages(ctx, channelID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessages", reflect.TypeOf((*MockMessageUsecase)(nil).FetchMessages), ctx, channelID, limit, offset)
}

// PostMessage mocks base method.
func (m *MockMessageUsecase) PostMessage(ctx context.Context, message *entities.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostMessage indicates an expected call of PostMessage.
func (mr *MockMessageUsecaseMockRecorder) PostMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostMessage", reflect.TypeOf((*MockMessageUsecase)(nil).PostMessage), ctx, message)
}

// UpdateMessage mocks base method.
func (m *MockMessageUsecase) UpdateMessage(ctx context.Context, message *entities.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMessage indicates an expected call of UpdateMessage.
func (mr *MockMessageUsecaseMockRecorder) UpdateMessage(ctx, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockMessageUsecase)(nil).UpdateMessage), ctx, message)
}
