package repository

import "github.com/shima004/slackclone/model"

type Repository interface {
	FindAllMessages([]model.Message, error)
}
