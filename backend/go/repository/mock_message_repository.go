package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) FetchMessages(ctx context.Context, userID uint) ([]model.Message, error) {
	args := m.Called(ctx, userID)

	var r0 []model.Message
	if rf, ok := args.Get(0).(func(context.Context, uint) []model.Message); ok {
		r0 = rf(ctx, userID)
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]model.Message)
		}
	}

	var r1 error
	if rf, ok := args.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (m *MockMessageRepository) PostMessage(ctx context.Context, message model.Message) error {
	args := m.Called(ctx, message)

	var r0 error
	if rf, ok := args.Get(0).(func(context.Context, model.Message) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
