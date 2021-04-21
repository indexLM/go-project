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

//func main() {
//	file, _ := os.OpenFile("D://sql.txt", os.O_RDWR, 0666)
//	defer file.Close()
//	fread := bufio.NewReader(file)
//	filee := xlsx.NewFile()
//	sheet, _ := filee.AddSheet("Sheet1")
//	for {
//		line, _, err := fread.ReadLine()
//		if err == io.EOF {
//			break
//		}
//		s := string(line)
//		split := strings.Split(s, "|")
//		row := sheet.AddRow()
//		row.SetHeightCM(1) //设置每行的高度
//		for _, value := range split {
//			cell := row.AddCell()
//			space := strings.Trim(value," ")
//			if strings.EqualFold(space,"10101") {
//				cell.Value = "男"
//			} else if strings.EqualFold(space,"10102") {
//				cell.Value = "女"
//			} else if strings.EqualFold(space,"10301") {
//				cell.Value = "身份证"
//			}else if strings.EqualFold(space,"10302") {
//				cell.Value = "护照"
//			}else if strings.EqualFold(space,"10303") {
//				cell.Value = "港澳通行证"
//			}else if strings.EqualFold(space,"2") {
//				cell.Value = "高管"
//			}else if strings.EqualFold(space,"3") {
//				cell.Value = "高管配偶"
//			} else if strings.EqualFold(space,"0") {
//				cell.Value = "普通员工"
//			} else {
//				cell.Value = value
//			}
//		}
//	}
//	err1 := filee.Save("D://file.xlsx")
//	if err1 != nil {
//		panic(err1)
//	}
//}
