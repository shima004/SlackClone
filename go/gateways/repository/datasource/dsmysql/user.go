package dsmysql

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type User interface {
	Create(ctx context.Context, user *entities.User) (uint, error)
	Delete(ctx context.Context, userID uint) error
	Read(ctx context.Context, email string) (*entities.User, error)
}
