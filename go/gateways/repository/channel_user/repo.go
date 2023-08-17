package channeluser

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/gateways/repository/datasource/dsmysql"
)

type ChannelUserRepo struct {
	mysqldsChannelUser dsmysql.ChannelUser
}

func NewChannelUserRepo(mysqldsChannelUser dsmysql.ChannelUser) *ChannelUserRepo {
	return &ChannelUserRepo{
		mysqldsChannelUser: mysqldsChannelUser,
	}
}

func (cur *ChannelUserRepo) CreateChannelUser(ctx context.Context, channelUser *entities.ChannelUser) (uint, error) {
	return cur.mysqldsChannelUser.Create(ctx, channelUser)
}

func (cur *ChannelUserRepo) DeleteChannelUser(ctx context.Context, userID uint, channelID uint) error {
	return cur.mysqldsChannelUser.Delete(ctx, userID, channelID)
}

func (cur *ChannelUserRepo) ReadChannelUsersByChannelID(ctx context.Context, channelID uint) ([]*entities.ChannelUser, error) {
	return cur.mysqldsChannelUser.ReadByChannelID(ctx, channelID)
}

func (cur *ChannelUserRepo) ReadChannelUsersByUserID(ctx context.Context, userID uint) ([]*entities.ChannelUser, error) {
	return cur.mysqldsChannelUser.ReadByUserID(ctx, userID)
}
