package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type appInfo struct {
	MysqlInfo *mysqlInfo
}

type mysqlInfo struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

func LoadConfig() *appInfo {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	appInfo := &appInfo{
		MysqlInfo: &mysqlInfo{
			User:     dbUser,
			Password: dbPass,
			Database: dbName,
			Host:     dbHost,
			Port:     dbPort,
		},
	}
	return appInfo
}
