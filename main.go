package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-project/global"
	"go-project/initserver"
	"go-project/utils"
	"os"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("get dir error")
	}
	err = utils.CofParse(dir+"/config.yaml", &global.MyServer)
	if err != nil {
		global.MyLogger.WithFields(logrus.Fields{"err": err}).Error("initserver config file error")
	}
	initserver.Logger()
	global.MyLogger.WithFields(logrus.Fields{"err": err}).Info("initserver config file success")
	initserver.Mysql()
	global.MyLogger.WithFields(logrus.Fields{"err": err}).Info("initserver mysql success")
	initserver.Redis()
}

func main() {
	address := fmt.Sprintf(":%d", global.MyServer.System.Port)
	router := initserver.Router()
	err := router.Run(address)
	global.MyLogger.WithFields(logrus.Fields{"err": err}).Info("gin start error")
}
