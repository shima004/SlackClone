package mysql

import (
	"context"

	"github.com/shima004/slackclone/entities"
	"gorm.io/gorm"
)

type MysqlChannelRepository struct {
	Conn *gorm.DB
}

func NewMysqlChannelRepository(conn *gorm.DB) *MysqlChannelRepository {
	return &MysqlChannelRepository{Conn: conn}
}

func (m *MysqlChannelRepository) CreateChannel(ctx context.Context, channel *entities.Channel) (uint, error) {
	result := m.Conn.Create(&channel)
	return channel.ID, result.Error
}

func (m *MysqlChannelRepository) DeleteChannel(ctx context.Context, channelID uint) error {
	result := m.Conn.Delete(&entities.Channel{}, channelID)
	if result.Error != nil {
		return result.Error
	}

	result = m.Conn.Delete(&entities.ChannelUser{}, "channel_id = ?", channelID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (m *MysqlChannelRepository) FetchChannel(ctx context.Context, channelID uint) (*entities.Channel, error) {
	channel := &entities.Channel{}
	// チャンネルが存在するか確認
	var count int64
	m.Conn.Model(&entities.Channel{}).Where("id = ?", channelID).Count(&count)
	if count == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	result := m.Conn.First(&channel, channelID)
	return channel, result.Error
}
