//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, userID uint) error
	// UpdateUser(ctx context.Context, user model.User) error
	FetchUserPassword(ctx context.Context, email string) (string, error)
}
