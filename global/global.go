package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go-project/config"
	"gorm.io/gorm"
)

var (
	MyServer *config.Server
	MyLogger *logrus.Logger
	MyDb     *gorm.DB
	MySqlx   *sqlx.DB
	MyRedis  *redis.Client
)
