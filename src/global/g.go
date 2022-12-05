package global

import (
	"log"

	"github.com/blog/src/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Settings = conf.ServerConf{}
	DB       *gorm.DB
)

func LoadSettings() {
	initSettings()
	initDB()
}

func initDB() {
	mysqlUri := Settings.Mysql.Uri
	db, err := gorm.Open(mysql.Open(mysqlUri), &gorm.Config{})
	if err != nil {
		log.Fatal("initDB error: can not open mysql")
	}
	DB = db
}

func initSettings() {
	LoadYaml("conf.yaml", &Settings)
}