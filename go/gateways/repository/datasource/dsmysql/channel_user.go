package dsmysql

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type ChannelUser interface {
	Create(ctx context.Context, channelUser *entities.ChannelUser) (uint, error)
	Delete(ctx context.Context, userID uint, channelID uint) error
	ReadByChannelID(ctx context.Context, channelID uint) ([]*entities.ChannelUser, error)
	ReadByUserID(ctx context.Context, userID uint) ([]*entities.ChannelUser, error)
}
