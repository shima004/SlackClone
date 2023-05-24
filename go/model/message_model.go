package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID    uint   `json:"user_id" validate:"required"`
	ChannelID uint   `json:"channel_id" validate:"required"`
	Text      string `json:"text" validate:"required"`
}
