package model

import (
	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model
	ChannelId   string    `json:"channel_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     string    `json:"owner_id"`
	Users       []User    `json:"users"`
	Messages    []Message `json:"messages"`
}
