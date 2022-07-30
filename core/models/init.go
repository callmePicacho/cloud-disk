package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// InitDB 初始化 MySQL
func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("gorm New Engine Error:%v", err)
		return nil
	}
	return db
}

// InitRDB 初始化 Redis
func InitRDB(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       1,  // use default DB
	})
}
