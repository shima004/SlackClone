package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Message struct {
	gorm.Model
	UserId string    `json:"user_id"`
	Time   time.Time `json:"time"`
	Body   string    `json:"body"`
}
