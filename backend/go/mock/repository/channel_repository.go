// Code generated by MockGen. DO NOT EDIT.
// Source: channel_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/shima004/slackclone/model"
)

// MockChannelRepository is a mock of ChannelRepository interface.
type MockChannelRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChannelRepositoryMockRecorder
}

// MockChannelRepositoryMockRecorder is the mock recorder for MockChannelRepository.
type MockChannelRepositoryMockRecorder struct {
	mock *MockChannelRepository
}

// NewMockChannelRepository creates a new mock instance.
func NewMockChannelRepository(ctrl *gomock.Controller) *MockChannelRepository {
	mock := &MockChannelRepository{ctrl: ctrl}
	mock.recorder = &MockChannelRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelRepository) EXPECT() *MockChannelRepositoryMockRecorder {
	return m.recorder
}

// CreateChannel mocks base method.
func (m *MockChannelRepository) CreateChannel(ctx context.Context, channel *model.Channel) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChannel", ctx, channel)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChannel indicates an expected call of CreateChannel.
func (mr *MockChannelRepositoryMockRecorder) CreateChannel(ctx, channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannel", reflect.TypeOf((*MockChannelRepository)(nil).CreateChannel), ctx, channel)
}

// DeleteChannel mocks base method.
func (m *MockChannelRepository) DeleteChannel(ctx context.Context, channelID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteChannel", ctx, channelID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteChannel indicates an expected call of DeleteChannel.
func (mr *MockChannelRepositoryMockRecorder) DeleteChannel(ctx, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteChannel", reflect.TypeOf((*MockChannelRepository)(nil).DeleteChannel), ctx, channelID)
}
