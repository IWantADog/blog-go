package global

import (
	"log"

	"github.com/blog/src/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Settings = conf.ServerConf{}
	globalDB *gorm.DB
)

func LoadSettings() {
	initSettings()
	initDB()
}

func GetDB() *gorm.DB {
	if globalDB == nil {
		log.Fatal("error: db is nil")
	}
	return globalDB
}

func initDB() {
	mysqlUri := Settings.Mysql.Uri
	db, err := gorm.Open(mysql.Open(mysqlUri), &gorm.Config{})
	if err != nil {
		log.Fatal("initDB error: can not open mysql")
	}
	globalDB = db
}

func initSettings() {
	LoadYaml("conf.yaml", &Settings)
}
