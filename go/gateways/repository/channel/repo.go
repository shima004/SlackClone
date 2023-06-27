package channel

import (
	"context"

	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/gateways/repository/datasource/dsmysql"
)

type ChannelRepo struct {
	dsmysqlChannel dsmysql.Channel
}

func NewChannelRepo(dsmysqlChannel dsmysql.Channel) *ChannelRepo {
	return &ChannelRepo{
		dsmysqlChannel: dsmysqlChannel,
	}
}

func (cr *ChannelRepo) CreateChannel(ctx context.Context, channel *entities.Channel) (uint, error) {
	return cr.dsmysqlChannel.Create(ctx, channel)
}

func (cr *ChannelRepo) DeleteChannel(ctx context.Context, channelID uint) error {
	return cr.dsmysqlChannel.Delete(ctx, channelID)
}

func (cr *ChannelRepo) ReadChannel(ctx context.Context, channelID uint) (*entities.Channel, error) {
	return cr.dsmysqlChannel.Read(ctx, channelID)
}
