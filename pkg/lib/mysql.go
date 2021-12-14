package lib

import (
	_ "github.com/go-sql-driver/mysql" // mysql驱动
)

// DataBaseConfig 数据库配置
type DataBaseConfig struct {
	Host string
	User string
	Password string
	DataBase string
}

func NewMysql(config DataBaseConfig)  {

}
