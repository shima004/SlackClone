package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) FetchMessages(ctx context.Context, auther string) ([]model.Message, error) {
	args := m.Called()
	return args.Get(0).([]model.Message), args.Error(1)
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
