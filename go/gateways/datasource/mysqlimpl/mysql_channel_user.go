package mysqlimpl

import (
	"context"

	"github.com/shima004/chat-server/entities"
	"gorm.io/gorm"
)

type MysqlChannelUser struct {
	Conn *gorm.DB
}

func NewMysqlChannelUser(conn *gorm.DB) *MysqlChannelUser {
	return &MysqlChannelUser{Conn: conn}
}

func (m *MysqlChannelUser) Create(ctx context.Context, channelUser *entities.ChannelUser) (uint, error) {
	result := m.Conn.Create(&channelUser)
	return channelUser.ID, result.Error
}

func (m *MysqlChannelUser) Delete(ctx context.Context, userID uint, channelID uint) error {
	result := m.Conn.Delete(&entities.ChannelUser{}, "user_id = ? AND channel_id = ?", userID, channelID)
	return result.Error
}

func (m *MysqlChannelUser) ReadByChannelID(ctx context.Context, channelID uint) ([]*entities.ChannelUser, error) {
	channelUsers := []*entities.ChannelUser{}
	result := m.Conn.Where("channel_id = ?", channelID).Find(&channelUsers)
	return channelUsers, result.Error
}

func (m *MysqlChannelUser) ReadByUserID(ctx context.Context, userID uint) ([]*entities.ChannelUser, error) {
	channelUsers := []*entities.ChannelUser{}
	result := m.Conn.Where("user_id = ?", userID).Find(&channelUsers)
	return channelUsers, result.Error
}
