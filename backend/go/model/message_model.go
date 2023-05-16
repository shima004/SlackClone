package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID     int    `json:"id"`
	Auther string `json:"auther"`
	Text   string `json:"text"`
}
