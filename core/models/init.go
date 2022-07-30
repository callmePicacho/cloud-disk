package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Engine = InitDB()
var RDB = InitRDB()

// InitDB 初始化 MySQL
func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("gorm New Engine Error:%v", err)
		return nil
	}
	return db
}

// InitRDB 初始化 Redis
func InitRDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
}
