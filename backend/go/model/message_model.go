package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	ChannelID uint   `json:"channel_id"`
	Text      string `json:"text"`
}
