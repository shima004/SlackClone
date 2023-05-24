package mysql

import (
	"context"

	"github.com/shima004/slackclone/model"
	"gorm.io/gorm"
)

type mysqlChannelRepository struct {
	Conn *gorm.DB
}

func NewMysqlChannelRepository(conn *gorm.DB) *mysqlChannelRepository {
	return &mysqlChannelRepository{Conn: conn}
}

func (m *mysqlChannelRepository) CreateChannel(ctx context.Context, name string) (uint, error) {
	channel := model.Channel{Name: name}
	result := m.Conn.Create(&channel)
	return channel.ID, result.Error
}

func (m *mysqlChannelRepository) DeleteChannel(channelID uint) error {
	result := m.Conn.Delete(&model.Channel{}, channelID)
	if result.Error != nil {
		return result.Error
	}

	result = m.Conn.Delete(&model.ChannelUser{}, "channel_id = ?", channelID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
