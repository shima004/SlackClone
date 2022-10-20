package model

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	GroupId     string    `json:"group_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     string    `json:"owner_id"`
	Users       []User    `json:"users"`
	Channels    []Channel `json:"channels"`
}
