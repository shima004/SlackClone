//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type MessageRepository interface {
	ReadMessages(ctx context.Context, channelID uint, limit int, offset int) ([]*entities.Message, error)
	CreateMessage(ctx context.Context, message *entities.Message) (uint, error)
	DeleteMessage(ctx context.Context, messageID uint) error
	UpdateMessage(ctx context.Context, message *entities.Message) error
}
