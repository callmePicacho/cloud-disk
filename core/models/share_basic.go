package models

import "gorm.io/gorm"

type ShareBasic struct {
	gorm.Model
	Identity           string
	UserIdentity       string // 用户标识
	RepositoryIdentity string // 中心资源池资源标识
	ExpiredTime        int
	ClickNum           int
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
