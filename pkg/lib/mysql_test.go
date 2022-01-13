package lib

import (
	"github.com/stretchr/testify/assert"
	"github.com/voyager-go/start-go-api/entities"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewMysql(t *testing.T) {
	/*
		CREATE TABLE `sys_user` (
		  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
		  `nickname` varchar(80) COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
		  `phone` varchar(11) COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
		  `password` varchar(200) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
		  `status` tinyint(1) NOT NULL COMMENT '是否启用 0:禁用 1:启用',
		  `created_at` int unsigned DEFAULT NULL COMMENT '创建时间',,
		  `updated_at` int DEFAULT NULL COMMENT '更新时间',
		  PRIMARY KEY (`id`),
		  UNIQUE KEY `phone` (`phone`) USING BTREE COMMENT '手机号唯一',
		  KEY `status` (`status`) USING BTREE COMMENT '是否可用的筛选条件'
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户信息表';
	*/
	dbCfg := DataBaseConfig{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Password: "root",
		DbName:   "startgoapi",
	}
	db, err := NewMysql(dbCfg)
	assert.Nil(t, err)
	passwd, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	assert.Nil(t, err)
	user := entities.User{
		Nickname: "张三",
		Phone:    "15106191191",
		Password: string(passwd),
	}
	err = db.Table("sys_user").Create(&user).Error
	assert.Nil(t, err)
}
