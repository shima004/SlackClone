package model

import "gorm.io/gorm"

type Message struct {
	// ID        uint `gorm:"primaryKey;default:auto_random()"`
	// CreatedAt string
	// UpdatedAt string
	// DeletedAt string `gorm:"index"`
	gorm.Model
	UserID    uint   `json:"user_id" validate:"required"`
	ChannelID uint   `json:"channel_id" validate:"required"`
	Text      string `json:"text" validate:"required"`
}
