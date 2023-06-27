package mysqlimpl

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"gorm.io/gorm"
)

type MysqlChannel struct {
	Conn *gorm.DB
}

func NewMysqlChannel(conn *gorm.DB) *MysqlChannel {
	return &MysqlChannel{Conn: conn}
}

func (m *MysqlChannel) Create(ctx context.Context, channel *entities.Channel) (uint, error) {
	result := m.Conn.Create(&channel)
	return channel.ID, result.Error
}

func (m *MysqlChannel) Delete(ctx context.Context, channelID uint) error {
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

func (m *MysqlChannel) Read(ctx context.Context, channelID uint) (*entities.Channel, error) {
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
