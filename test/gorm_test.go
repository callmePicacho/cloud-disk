package test

import (
	"cloud-disk/core/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]models.UserBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(data)
}
