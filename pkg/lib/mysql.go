package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	//_ "github.com/go-sql-driver/mysql" // mysql驱动
)

// DataBaseConfig 数据库配置
type DataBaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

// NewMysql 构造MySQL服务
func NewMysql(config DataBaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 是否设置单数表名，设置为 是
		},
	})
	if err != nil {
		return nil, fmt.Errorf("无法连接数据库，请先检查MySQL配置信息，错误详情为: %s", err.Error())
	}
	// GORM 使用 database/sql 维护连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, err
}
