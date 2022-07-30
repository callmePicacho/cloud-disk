package models

import "gorm.io/gorm"

type RepositoryPool struct {
	gorm.Model
	Identity string
	Hash     string
	Name     string
	Ext      string
	Size     int64
	Path     string
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
