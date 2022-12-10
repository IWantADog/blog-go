package global

import (
	"log"

	"github.com/blog/src/conf"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Settings = conf.ServerConf{}
	globalDB *gorm.DB
	rdb      *redis.Client
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

func GetRedis() *redis.Client {
	// TODO: 加锁，避免多次实例化
	if rdb == nil {
		rdb = initRedis()
	}
	return rdb
}

func initRedis() *redis.Client {
	uri := Settings.Redis.Uri
	opt, err := redis.ParseURL(uri)
	if err != nil {
		log.Fatalf("initRedis error: %s\n", err.Error())
	}
	return redis.NewClient(opt)
}

func initSettings() {
	LoadYaml("conf.yaml", &Settings)
}
