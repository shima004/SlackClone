package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
