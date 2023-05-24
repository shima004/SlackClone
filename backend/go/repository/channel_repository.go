//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
)

type ChannelRepository interface {
	CreateChannel(ctx context.Context, channel *model.Channel) (uint, error)
	DeleteChannel(ctx context.Context, channelID uint) error
}
