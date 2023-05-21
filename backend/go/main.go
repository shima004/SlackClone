package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/shima004/slackclone/delivery"
	"github.com/shima004/slackclone/model"
	repo "github.com/shima004/slackclone/repository/mysql"
	"github.com/shima004/slackclone/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	g := e.Group("/api")

	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(&model.Message{})
	db.Migrator().AutoMigrate(&model.Message{})
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	repo := repo.NewMysqlMessageRepository(db)
	mu := usecase.DefaultMessageUsecase{MessageRepository: repo}
	delivery.NewMessageHandler(g, &mu)

	e.Logger.Fatal(e.Start(":8080"))
}
