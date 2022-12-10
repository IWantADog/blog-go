package main

import (
	"github.com/blog/src/global"
	"github.com/blog/src/model"
)

func migrateDB() {
	global.LoadSettings()
	db := global.GetDB()
	modelList := []interface{}{
		&model.Author{},
		&model.Blog{},
		&model.Tag{},
	}
	db.AutoMigrate(modelList...)
}

func main() {
	migrateDB()
}
