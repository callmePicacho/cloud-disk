package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Engine = InitDB()

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
