//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type MessageRepository interface {
	FetchMessages(ctx context.Context, channelID uint, limit int, offset int) ([]entities.Message, error)
	PostMessage(ctx context.Context, message entities.Message) error
	DeleteMessage(ctx context.Context, messageID uint) error
	UpdateMessage(ctx context.Context, message entities.Message) error
}
