package message

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/gateways/repository/datasource/dsmysql"
)

type MessageRepo struct {
	mysqldsMessage dsmysql.Message
}

func NewMessageRepo(mysqldsMessage dsmysql.Message) *MessageRepo {
	return &MessageRepo{
		mysqldsMessage: mysqldsMessage,
	}
}

func (mr *MessageRepo) ReadMessages(ctx context.Context, channelID uint, limit int, offset int) ([]*entities.Message, error) {
	return mr.mysqldsMessage.Read(ctx, channelID, limit, offset)
}

func (mr *MessageRepo) CreateMessage(ctx context.Context, message *entities.Message) (uint, error) {
	return mr.mysqldsMessage.Create(ctx, message)
}

func (mr *MessageRepo) DeleteMessage(ctx context.Context, messageID uint) error {
	return mr.mysqldsMessage.Delete(ctx, messageID)
}

func (mr *MessageRepo) UpdateMessage(ctx context.Context, message *entities.Message) error {
	return mr.mysqldsMessage.Update(ctx, message)
}
