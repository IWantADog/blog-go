package src

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Settings = make(map[string]interface{})
var DB *gorm.DB
var App = gin.Default()

func init() {
	// initSettings()
	// initDB()
}

func initDB() {
	mysqlUri, ok := Settings["mysql"]
	if ok == false {
		log.Fatal("initDB error: can not get mysql uir")
	}
	db, err := gorm.Open(mysql.Open(mysqlUri.(string)), &gorm.Config{})
	if err != nil {
		log.Fatal("initDB error: can not open mysql")
	}
	DB = db
}

func initSettings() {
	// currentFile, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal("load conf faild")
	// }
	// confFilePath := filepath.Dir(currentFile) + "conf.yaml"
}
