package mysqlimpl

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"gorm.io/gorm"
)

type MysqlMessage struct {
	db *gorm.DB
}

func NewMysqlMessage(db *gorm.DB) *MysqlMessage {
	return &MysqlMessage{db: db}
}

func (r *MysqlMessage) Read(ctx context.Context, channelID uint, limit int, offset int) ([]*entities.Message, error) {
	var messages []*entities.Message
	result := r.db.Where("channel_id = ?", channelID).Order("created_at desc").Limit(limit).Offset(offset).Find(&messages)
	return messages, result.Error
}

func (r *MysqlMessage) Create(ctx context.Context, message *entities.Message) (uint, error) {
	result := r.db.Create(&message)
	return message.ID, result.Error
}

func (r *MysqlMessage) Delete(ctx context.Context, messageID uint) error {
	result := r.db.Delete(&entities.Message{}, messageID)
	return result.Error
}

func (r *MysqlMessage) Update(ctx context.Context, message *entities.Message) error {
	// データが有るかどうか確認
	var count int64
	r.db.Model(&entities.Message{}).Where("id = ?", message.ID).Count(&count)
	if count == 0 {
		return gorm.ErrRecordNotFound
	}

	result := r.db.Model(&entities.Message{}).Where("id = ?", message.ID).Updates(&message)
	return result.Error
}
