package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/core/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitDB(c.Mysql.DataSource),
		RDB:    models.InitRDB(c.Redis.Addr),
	}
}
