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

//type People struct {
//	Name string `res:"name"`
//	oppo *Phone `res:"phone"`
//}
//type Phone struct {
//	number uint64 `rr:"ceshi"`
//}
//
//func main() {
//	p := new(People)
//	p2 := new(Phone)
//	p2.number=18001
//	p.Name="小明"
//	p.oppo=p2
//	typeOf := reflect.TypeOf(*p)
//	valueOf := reflect.ValueOf(*p)
//	fmt.Println(typeOf.Kind())
//	fmt.Println(typeOf)
//	fmt.Println(valueOf)
//	i1 := typeOf.NumField()
//	for i := 0; i < i1; i++ {
//		field := typeOf.Field(i)
//		v := valueOf.Field(i)
//		if field.Type.Kind() == reflect.Ptr {
//			elem := field.Type.Elem()
//			numField := elem.NumField()
//			value := v.Elem()
//			for i2 := 0; i2 < numField; i2++ {
//				valueE := value.Field(i2)
//				structField := elem.Field(i2)
//				fmt.Println("指针",structField.Name)
//				fmt.Println("指针",structField.Tag.Get("rr"))
//				fmt.Println("指针",valueE)
//			}
//		}else  {
//			tag := field.Tag
//			name := field.Name
//			get := tag.Get("res")
//			value := v
//			fmt.Println(name)
//			fmt.Println(get)
//			fmt.Println(value)
//		}
//	}
//}
