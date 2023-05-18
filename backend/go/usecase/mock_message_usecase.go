package usecase

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/mock"
)

type MockMessageUsercase struct {
	mock.Mock
}

func (m *MockMessageUsercase) FetchMessages(ctx context.Context, userID uint) (res []model.Message, err error) {
	ret := m.Called(ctx, userID)

	var r0 []model.Message
	if rf, ok := ret.Get(0).(func(context.Context, uint) []model.Message); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockMessageUsercase) PostMessage(ctx context.Context, message model.Message) (err error) {
	ret := m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Message) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
