//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package inputport

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user entities.User) error
	DeleteUser(ctx context.Context, userID uint) error
	ReadUserByUserID(ctx context.Context, userID uint) (*entities.User, error)
	SubscribeChannel(ctx context.Context, userID uint, channelID uint) error
	Login(ctx context.Context, email string, password string) (string, error)
}
