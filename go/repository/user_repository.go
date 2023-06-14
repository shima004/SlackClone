//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, userID uint) error
	// UpdateUser(ctx context.Context, user entities.User) error
	FetchUserPassword(ctx context.Context, email string) (string, error)
}
