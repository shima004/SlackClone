//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package inputport

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/inputport/validation"
)

type MessageUsecase interface {
	FetchMessages(ctx context.Context, in *validation.FatchMessagesInput) (res []*entities.Message, err error)
	PostMessage(ctx context.Context, in *validation.PostMessageInput) (err error)
	DeleteMessage(ctx context.Context, in *validation.DeleteMessageInput) (err error)
	UpdateMessage(ctx context.Context, in *validation.UpdateMessageInput) (err error)
}
