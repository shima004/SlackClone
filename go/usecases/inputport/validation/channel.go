package validation

import (
	"github.com/shima004/chat-server/entities"
	cerror "github.com/shima004/chat-server/entities/error"
)

// type ChannelUsecase interface {
// 	CreateChannel(ctx context.Context, channel *entities.Channel) (uint, error)
// 	DeleteChannel(ctx context.Context, channelID uint) error
// 	FetchChannel(ctx context.Context, channelID uint) (*entities.Channel, error)
// }

type CreateChannelInput struct {
	Channel *entities.Channel
}

func (in *CreateChannelInput) Validate() error {
	e := cerror.NewValidationError()
	if in.Channel.Name == "" {
		e.Add("name", "must not be empty")
	} else if len(in.Channel.Name) > 100 {
		e.Add("name", "must be less than 100 characters")
	}

	if e.HasErrors() {
		return e
	}

	return nil
}

type DeleteChannelInput struct {
	ChannelID uint
}

func (in *DeleteChannelInput) Validate() error {
	return nil
}

type FetchChannelInput struct {
	ChannelID uint
}

func (in *FetchChannelInput) Validate() error {
	return nil
}
