package dbfunc

import (
	"Slack/model"
)

func PostLoginInfo(email string, password string) (user model.User, err error) {
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザー情報を取得
	var u model.User
	if err = db.Where("email = ?", email).First(&u).Error; err != nil {
		return u, err
	}

	return u, err
}
