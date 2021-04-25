package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go-project/config"
	"gorm.io/gorm"
)

var (
	MyServer *config.Server
	MyLogger *logrus.Logger
	MyDb     *gorm.DB
	MyRedis  *redis.Client
)
