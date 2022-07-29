package models

import "gorm.io/gorm"

type UserBasic struct {
	Identity string
	Name     string
	Password string
	Email    string
	gorm.Model
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
