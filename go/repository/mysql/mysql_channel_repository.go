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

func (m *mysqlChannelRepository) CreateChannel(ctx context.Context, channel *model.Channel) (uint, error) {
	result := m.Conn.Create(&channel)
	return channel.ID, result.Error
}

func (m *mysqlChannelRepository) DeleteChannel(ctx context.Context, channelID uint) error {
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

func (m *mysqlChannelRepository) FetchChannel(ctx context.Context, channelID uint) (*model.Channel, error) {
	channel := &model.Channel{}
	// チャンネルが存在するか確認
	var count int64
	m.Conn.Model(&model.Channel{}).Where("id = ?", channelID).Count(&count)
	if count == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	result := m.Conn.First(&channel, channelID)
	return channel, result.Error
}
