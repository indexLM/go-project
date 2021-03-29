package global

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go-project/config"
)

var (
	MyServer *config.Server
	MyLogger *logrus.Logger
	MyDb     *sql.DB
	MyRedis  *redis.Client
)
