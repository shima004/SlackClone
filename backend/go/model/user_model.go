package model

type User struct {
	// gorm.Model
	ID 		   uint   `gorm:"primaryKey;default:auto_random()"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
