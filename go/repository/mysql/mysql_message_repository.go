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

func (r *MysqlMessageRepository) FetchMessages(ctx context.Context, channelID uint, limit int, offset int) ([]model.Message, error) {
	var messages []model.Message
	result := r.db.Where("channel_id = ?", channelID).Order("created_at desc").Limit(limit).Offset(offset).Find(&messages)
	return messages, result.Error
}

func (r *MysqlMessageRepository) PostMessage(ctx context.Context, message model.Message) error {
	result := r.db.Create(&message)
	return result.Error
}

func (r *MysqlMessageRepository) DeleteMessage(ctx context.Context, messageID uint) error {
	result := r.db.Delete(&model.Message{}, messageID)
	return result.Error
}

func (r *MysqlMessageRepository) UpdateMessage(ctx context.Context, message model.Message) error {
	// データが有るかどうか確認
	var count int64
	r.db.Model(&model.Message{}).Where("id = ?", message.ID).Count(&count)
	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	result := r.db.Model(&model.Message{}).Where("id = ?", message.ID).Updates(&message)
	return result.Error
}
