package initserver

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"go-project/global"
	"go-project/handler"
	"go-project/middleware"
	"io"
	"log"
	"os"
	"time"
)

//初始化日志打印
func Logger() {
	global.MyLogger = logrus.New()
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	fmt.Println(global.MyServer.Log.Prefix)
	writer3, err := os.OpenFile(global.MyServer.Log.Prefix, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("创建日志文件失败,失败原因: %v", err)
	}
	global.MyLogger.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	global.MyLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	global.MyLogger.SetLevel(logrus.InfoLevel)
}

//初始化数据库（mysql）
func Mysql() {
	connInfo := fmt.Sprintf("%s:%s@(%s)/%s?useUnicode=true&allowMultiQueries=true&characterEncoding=UTF-8&zeroDateTimeBehavior=CONVERT_TO_NULL",
		global.MyServer.Mysql.Username,
		global.MyServer.Mysql.Password,
		global.MyServer.Mysql.Host,
		global.MyServer.Mysql.Db)
	var err error
	global.MyDb, err = sql.Open("mysql", connInfo)
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("初始化连接数据库失败")
	}
	err = global.MyDb.Ping()
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("ping mysql")
	}
	global.MyDb.SetMaxIdleConns(global.MyServer.Mysql.Conn.MaxIdle)
	global.MyDb.SetMaxOpenConns(global.MyServer.Mysql.Conn.MaxOpen)
	global.MyDb.SetConnMaxLifetime(5 * time.Minute)
}

//初始化缓存（redis）
func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.MyServer.Redis.Addr,
		Password: global.MyServer.Redis.Password,
		DB:       global.MyServer.Redis.Db,
	})
	global.MyRedis = client
}

//初始化路由
func Router() *gin.Engine {
	// 默认已经连接了 Logger and Recovery 中间件
	var Router = gin.Default()
	//全局中间件
	//跨域请求放行中间件
	Router.Use(middleware.Cors())
	RouterGroup := Router.Group("")
	//路由注册
	handler.RouterAuthInit(RouterGroup)

	global.MyLogger.Info("路由注册成功")
	return Router
}
