package dsmysql

import (
	"context"

	"github.com/shima004/chat-server/entities"
)

type Message interface {
	Create(ctx context.Context, message *entities.Message) (uint, error)
	Delete(ctx context.Context, messageID uint) error
	ReadByChannelID(ctx context.Context, channelID uint, limit int, offset int) ([]*entities.Message, error)
	ReadByMessageID(ctx context.Context, messageID uint) (*entities.Message, error)
	Update(ctx context.Context, message *entities.Message) error
}
