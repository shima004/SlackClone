package usecase

import (
	"context"
	"time"

	"github.com/shima004/slackclone/model"
	"github.com/shima004/slackclone/repository"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, userID uint) error
	// UpdateUser(ctx context.Context, user model.User) error
	Login(ctx context.Context, email string, password string) (string, error)
}

type DefaultUserUsecase struct {
	UserRepository repository.UserRepository
	ContextTimeout time.Duration
}

func (u *DefaultUserUsecase) CreateUser(ctx context.Context, user model.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.UserRepository.CreateUser(ctx, user)
}

func (u *DefaultUserUsecase) DeleteUser(ctx context.Context, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.UserRepository.DeleteUser(ctx, userID)
}

func (u *DefaultUserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.UserRepository.FetchUserPassword(ctx, email)
}
