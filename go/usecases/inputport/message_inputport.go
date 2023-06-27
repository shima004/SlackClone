//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package inputport

import (
	"context"

	"github.com/shima004/slackclone/entities"
)

type MessageUsecase interface {
	FetchMessages(ctx context.Context, channelID uint, limit int, offset int) (res []*entities.Message, err error)
	PostMessage(ctx context.Context, message *entities.Message) (err error)
	DeleteMessage(ctx context.Context, messageID uint) (err error)
	UpdateMessage(ctx context.Context, message *entities.Message) (err error)
}
