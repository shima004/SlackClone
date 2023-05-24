package mysql

import (
	"database/sql"

	"github.com/shima004/slackclone/config"
	"github.com/shima004/slackclone/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, *sql.DB, error) {
	appInfo := config.LoadConfig()
	dsn := appInfo.MysqlInfo.User + ":" + appInfo.MysqlInfo.Password + "@tcp(" + appInfo.MysqlInfo.Host + ":" + appInfo.MysqlInfo.Port + ")/" + appInfo.MysqlInfo.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	InitTable(db, &model.Message{})
	return db, sqlDB, nil
}

func InitTable(db *gorm.DB, models ...interface{}) {
	// db.Migrator().DropTable(models...)
	db.Migrator().AutoMigrate(models...)
}
