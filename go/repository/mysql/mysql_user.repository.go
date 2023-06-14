package mysql

import (
	"github.com/shima004/slackclone/entities"
	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{Conn: conn}
}

func (mur *MysqlUserRepository) CreateUser(user entities.User) error {
	return mur.Conn.Create(&user).Error
}

func (mur *MysqlUserRepository) DeleteUser(userID uint) error {
	return mur.Conn.Delete(&entities.User{}, userID).Error
}

func (mur *MysqlUserRepository) FetchUserPassword(email string) (string, error) {
	var user entities.User
	if err := mur.Conn.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}
