package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Auther string `json:"auther"`
	Text   string `json:"text"`
}
