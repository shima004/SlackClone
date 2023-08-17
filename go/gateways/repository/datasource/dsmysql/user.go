package dsmysql

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type User interface {
	Create(ctx context.Context, user *entities.User) (uint, error)
	Delete(ctx context.Context, userID uint) error
	// Read(ctx context.Context, email string) (*entities.User, error)
	ReadByUserID(ctx context.Context, userID uint) (*entities.User, error)
	ReadByUserEmail(ctx context.Context, userEmail string) (*entities.User, error)
}
