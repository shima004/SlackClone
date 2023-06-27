package dsmysql

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type Channel interface {
	Create(ctx context.Context, channel *entities.Channel) (uint, error)
	Delete(ctx context.Context, channelID uint) error
	Read(ctx context.Context, channelID uint) (*entities.Channel, error)
}
