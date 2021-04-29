# go web
`author: indexLM`
`time: 2021-03-02`

## 版本说明
- golang: 1.15
- mysql: 5.7
- redis: 4.0.14
## 项目介绍

这个项目主要是为了方便学习gin框架、jwt，因此将其结合，完成用户登陆、修改操作,使用jwt进行鉴权。

## 使用说明
- 下载项目
```shell
$ git clone https://github.com/indexLM/go-project.git
```
将项目放到$GOPATH/src/go-project/目录下。
- 下载相关依赖：
```shell
$ export GO111MODULE=on
$ go mod download
```
- 运行代码，进入到go-project目录
```shell
$ go run main.go
```

## 项目架构

### 目录结构

```lua
gin_jwt_swagger
├── config -- 配置
├── dao -- db、cache操作方法
├── global -- 全局变量声明
├── handler -- API处理
├── initserver  -- 初始化相关服务
├── middleware -- 中间件
├── model -- 模型文件
├── service -- 逻辑服务
├── utils -- 工具文件
├── config.yaml -- 配置文件
└── main.go -- 主函数
```
