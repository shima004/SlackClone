package model

import "gorm.io/gorm"

type ChannelUser struct {
	gorm.Model
	ChannelID uint `json:"channel_id"`
	UserID    uint `json:"user_id"`
}
