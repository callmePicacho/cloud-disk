package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/core/internal/middleware"
	"cloud-disk/core/models"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.InitDB(c.Mysql.DataSource),
		RDB:    models.InitRDB(c.Redis.Addr),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
