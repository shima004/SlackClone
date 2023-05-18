package mysql

import (
	"context"

	"github.com/shima004/slackclone/model"
	"gorm.io/gorm"
)

type MysqlMessageRepository struct {
	db *gorm.DB
}

func NewMysqlMessageRepository(db *gorm.DB) *MysqlMessageRepository {
	return &MysqlMessageRepository{db: db}
}

func (r *MysqlMessageRepository) FetchMessages(ctx context.Context, userID uint) ([]model.Message, error) {
	var messages []model.Message
	result := r.db.Where("user_id = ?", userID).Find(&messages)
	return messages, result.Error
}

func (r *MysqlMessageRepository) PostMessage(ctx context.Context, message model.Message) error {
	result := r.db.Create(&message)
	return result.Error
}
