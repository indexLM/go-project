package config

type Server struct {
	System System `json:"system" yaml:"system"`
	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Log    Log    `json:"log" yaml:"log"`
	Jwt    Jwt    `json:"jwt" yaml:"jwt"`
}

//系统配置
type System struct {
	Port int `json:"port" yaml:"port"`
}

//数据库配置
type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Db       string `json:"db" yaml:"db"`
	Conn     struct {
		MaxIdle int `json:"maxIdle" yaml:"maxIdle"`
		MaxOpen int `json:"maxOpen" yaml:"maxOpen"`
	}
}

//缓存配置
type Redis struct {
	Addr     string `json:"addr" yaml:"addr"`
	Db       int    `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
	PoolSize int    `json:"poolSize" yaml:"poolSize"`
	Cache    struct {
		TokenExpired int `json:"tokenExpired" yaml:"tokenExpired"`
	}
}

//日志配置
type Log struct {
	Prefix  string `json:"prefix" yaml:"prefix"`
	LogFile bool   `json:"log_file" yaml:"log-file"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	File    string `json:"file" yaml:"file"`
}

// JWT 签名结构
type Jwt struct {
	SignKey string `json:"signKey" yaml:"signKey"`
	Expires int64  `json:"expires" yaml:"expires"`
}
