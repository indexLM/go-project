package main

import (
	"fmt"
	"go-project/global"
	"go-project/initserver"
	"go-project/utils"
	"os"
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取项目目录失败")
	}
	err = utils.CofParse(dir+"/config.yaml", &global.MyServer)
	if err != nil {
		fmt.Println("初始化配置文件失败")
	}
	initserver.Logger()
	//initserver.MyGorm()
	initserver.MySqlx()
	utils.LogInfo("sys", "初始化sqlx框架成功")
	initserver.Redis()
	utils.LogInfo("sys", "初始化go-redis框架成功")
	initserver.Jwt()
	utils.LogInfo("sys", "初始化JWT成功")

}

func main() {
	address := fmt.Sprintf(":%d", global.MyServer.System.Port)
	router := initserver.Router()
	utils.LogInfo("sys", "初始化gin框架成功")
	err := router.Run(address)
	if err != nil {
		utils.LogError("sys", "web服务启动失败")
	}
}
