package dbfunc

import (
	"Slack/model"

	"github.com/labstack/echo/v4"
)

func PostCoin(c echo.Context, total int) (err error) { //引数と返り値
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザーが存在するかを確認
	user, err := GetUserFromToken(c)
	if err != nil {
		return err
	}

	coin := user.Coin + total //ユーザーのコインを更新
	// データベースを更新
	err = db.Model(model.User{}).Where("uuid = ?", user.UUID).Select("Coin").Updates(model.User{Coin: coin}).Error

	return err
}

func SetCoinFromUUID(uuid string, coin int) (err error) { //引数と返り値
	db := sqlConnect()
	defer db.Close()

	// データベースからユーザーが存在するかを確認
	user, err := GetUserFromUUID(uuid)
	if err != nil {
		return err
	}

	coin = user.Coin + coin //ユーザーのコインを更新
	// データベースを更新
	err = db.Model(model.User{}).Where("uuid = ?", uuid).Select("Coin").Updates(model.User{Coin: coin}).Error

	return err
}
