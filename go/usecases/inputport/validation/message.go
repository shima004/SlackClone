package validation

import (
	"github.com/shima004/chat-server/entities"
	cerror "github.com/shima004/chat-server/entities/error"
)

func textValidator(text string) error {
	e := cerror.NewValidationError()

	if text == "" {
		e.Add("text", "must not be empty")
	}

	if len(text) > 1000 {
		e.Add("text", "must be less than 1000 characters")
	}

	if e.HasErrors() {
		return e
	}

	return nil
}

type FatchMessagesInput struct {
	ChannelID uint
	Limit     int
	Offset    int
}

func (i *FatchMessagesInput) Validate() error {
	e := cerror.NewValidationError()

	if i.Limit < 0 {
		e.Add("limit", "must be greater than 0")
	}
	if i.Offset < 0 {
		e.Add("offset", "must be greater than 0")
	}

	if e.HasErrors() {
		return e
	}

	return nil
}

type PostMessageInput struct {
	Message *entities.Message
}

func (i *PostMessageInput) Validate() error {
	if err := textValidator(i.Message.Text); err != nil {
		return err
	}

	return nil
}

type DeleteMessageInput struct {
	MessageID uint
	UserID    uint
}

func (i *DeleteMessageInput) Validate() error {
	return nil
}

type UpdateMessageInput struct {
	Message *entities.Message
}

func (i *UpdateMessageInput) Validate() error {
	if err := textValidator(i.Message.Text); err != nil {
		return err
	}

	return nil
}
