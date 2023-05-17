package usecase

import (
	"context"

	"github.com/shima004/slackclone/model"
	"github.com/stretchr/testify/mock"
)

type MockMessageUsercase struct {
	mock.Mock
}

func (m *MockMessageUsercase) FetchMessages(ctx context.Context, auther string) (res []model.Message, err error) {
	ret := m.Called(ctx, auther)

	var r0 []model.Message
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Message); ok {
		r0 = rf(ctx, auther)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, auther)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
