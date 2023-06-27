package mysqlimpl

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

func (mur *MysqlUserRepository) Create(user entities.User) error {
	return mur.Conn.Create(&user).Error
}

func (mur *MysqlUserRepository) Delete(userID uint) error {
	return mur.Conn.Delete(&entities.User{}, userID).Error
}

func (mur *MysqlUserRepository) Read(email string) (*entities.User, error) {
	user := &entities.User{}
	result := mur.Conn.First(&user, "email = ?", email)
	return user, result.Error
}
