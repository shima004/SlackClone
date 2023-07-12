//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package inputport

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"github.com/shima004/chat-server/usecases/inputport/validation"
)

type ChannelUsecase interface {
	CreateChannel(ctx context.Context, in *validation.CreateChannelInput) (uint, error)
	DeleteChannel(ctx context.Context, in *validation.DeleteChannelInput) error
	FetchChannel(ctx context.Context, in *validation.FetchChannelInput) (*entities.Channel, error)
}
