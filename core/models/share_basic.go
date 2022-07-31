package models

import "gorm.io/gorm"

type ShareBasic struct {
	gorm.Model
	Identity               string
	UserIdentity           string // 用户标识
	UserRepositoryIdentity string // 用户资源池资源唯一标识
	ExpiredTime            int
	ClickNum               int
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
