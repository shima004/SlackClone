package user

import (
	"context"

	"github.com/shima004/slackclone/entities"
	"github.com/shima004/slackclone/gateways/repository/datasource/dsmysql"
)

type UserRepo struct {
	dsmysqlUser dsmysql.User
}

func NewUserRepo(dsmysqlUser dsmysql.User) *UserRepo {
	return &UserRepo{
		dsmysqlUser: dsmysqlUser,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user entities.User) (uint, error) {
	return ur.dsmysqlUser.Create(ctx, &user)
}

func (ur *UserRepo) DeleteUser(ctx context.Context, userID uint) error {
	return ur.dsmysqlUser.Delete(ctx, userID)
}

func (ur *UserRepo) ReadUserPassword(ctx context.Context, email string) (string, error) {
	user, err := ur.dsmysqlUser.Read(ctx, email)
	if err != nil {
		return "", err
	}

	return user.Password, nil
}
