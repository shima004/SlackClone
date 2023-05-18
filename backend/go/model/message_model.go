package model

type Message struct {
	// gorm.Model
	ID 		   uint   `gorm:"primaryKey;default:auto_random()"`
	UserID    uint   `json:"user_id" validate:"required"`
	ChannelID uint   `json:"channel_id" validate:"required"`
	Text      string `json:"text" validate:"required"`
}
