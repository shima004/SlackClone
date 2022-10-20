package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `json:"user_id"`
	Email    string `json:"email";gorm:"type:varchar(100);gorm:unique_index"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
