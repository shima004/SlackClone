//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entities.User) (uint, error)
	DeleteUser(ctx context.Context, userID uint) error
	ReadUserPassword(ctx context.Context, email string) (string, error)
}
