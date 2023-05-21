//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"context"

	"github.com/shima004/slackclone/model"
)

type MessageRepository interface {
	FetchMessages(context.Context, uint) ([]model.Message, error)
	PostMessage(context.Context, model.Message) error
	DeleteMessage(context.Context, uint) error
	UpdateMessage(context.Context, model.Message) error
}
