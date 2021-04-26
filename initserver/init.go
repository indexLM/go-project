package initserver

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go-project/global"
	"go-project/handler"
	"go-project/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func MyGorm() {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.MyServer.Mysql.Username,
		global.MyServer.Mysql.Password,
		global.MyServer.Mysql.Host,
		global.MyServer.Mysql.Db)
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connInfo, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置, &gorm.Config{})
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("初始化Gorm框架失败")
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("初始化Gorm框架失败")
		return
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(global.MyServer.Mysql.Conn.MaxIdle)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(global.MyServer.Mysql.Conn.MaxOpen)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	global.MyLogger.WithFields(logrus.Fields{"info": "sys"}).Error("初始化Gorm框架成功")
	initDb(db)
	global.MyDb = db
}

//初始化数据库（mysql）
func MySqlx() {
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.MyServer.Mysql.Username,
		global.MyServer.Mysql.Password,
		global.MyServer.Mysql.Host,
		global.MyServer.Mysql.Db)
	var err error
	db, err := sqlx.Open("mysql", connInfo)
	if err != nil {
		panic("初始化sqlx框架失败")
	}
	// 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(global.MyServer.Mysql.Conn.MaxIdle)
	// 设置打开数据库连接的最大数量
	db.SetMaxOpenConns(global.MyServer.Mysql.Conn.MaxOpen)
	// 设置了连接可复用的最大时间
	db.SetConnMaxLifetime(time.Hour)
	global.MySqlx = db
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
	//全局异常处理
	Router.Use(middleware.Recover)
	//跨域请求放行中间件
	Router.Use(middleware.Cors())
	RouterGroup := Router.Group("")
	//路由注册
	handler.RouterAuthInit(RouterGroup)
	return Router
}
