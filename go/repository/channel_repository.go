//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type ChannelRepository interface {
	CreateChannel(ctx context.Context, channel *entities.Channel) (uint, error)
	DeleteChannel(ctx context.Context, channelID uint) error
	FetchChannel(ctx context.Context, channelID uint) (*entities.Channel, error)
}
