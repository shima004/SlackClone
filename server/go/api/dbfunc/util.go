package dbfunc

import (
	"Slack/model"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func init() {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	panic(err)
	// }

	db := sqlConnect()
	db.AutoMigrate(&model.User{}) // テーブル作成
	defer db.Close()
}

func sqlConnect() (database *gorm.DB) {
	dbMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	db, err := gorm.Open(dbMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(dbMS, CONNECT)
		}
	}
	return db
}
