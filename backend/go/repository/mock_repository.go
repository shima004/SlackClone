package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FetchMessages(ctx context.Context, auther string) ([]model.Message, error) {
	args := m.Called()
	return args.Get(0).([]model.Message), args.Error(1)
}
