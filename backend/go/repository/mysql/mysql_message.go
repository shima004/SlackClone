package mysql

import (
	"github.com/shima004/slackclone/model"
	"gorm.io/gorm"
)

type MysqlMessageRepository struct {
	db *gorm.DB
}

func NewMysqlMessageRepository(db *gorm.DB) *MysqlMessageRepository {
	return &MysqlMessageRepository{db: db}
}

func (r *MysqlMessageRepository) FetchMessages(auther string) ([]model.Message, error) {
	var messages []model.Message
	result := r.db.Where("auther = ?", auther).Find(&messages)
	return messages, result.Error
}

func (r *MysqlMessageRepository) PostMessage(message model.Message) error {
	result := r.db.Create(&message)
	return result.Error
}
