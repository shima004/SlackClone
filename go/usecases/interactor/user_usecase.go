package interactor

import (
	"context"
	"time"

	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/gateways/repository/user"
)

type DefaultUserUsecase struct {
	UserRepository user.UserRepo
	ContextTimeout time.Duration
}

func (u *DefaultUserUsecase) CreateUser(ctx context.Context, user entities.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	_, err := u.UserRepository.CreateUser(ctx, user)

	return err
}

func (u *DefaultUserUsecase) DeleteUser(ctx context.Context, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.UserRepository.DeleteUser(ctx, userID)
}

func (u *DefaultUserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	p, err := u.UserRepository.ReadUserPassword(ctx, email)
	if err != nil {
		return "", err
	}

	if p != password {
		return "", entities.ErrInvalidPassword
	}

	return "token", nil
}
