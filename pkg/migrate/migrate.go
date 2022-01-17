package migrate

import (
	"github.com/voyager-go/start-go-api/entities"
	"github.com/voyager-go/start-go-api/global"
)

// AuthMigrate
//@author: [张文杰](https://github.com/voyager-go)
//@slogan    岁岁平，岁岁安，岁岁平安
//@description: 数据表自动迁移
//@create_date: 2022/1/17
//@create_time: 3:55 下午
//@param:
//@return:
func AuthMigrate() {
	global.DB.AutoMigrate(
		&entities.SysRole{},
		&entities.User{},
		&entities.SysRoleUser{},
		&entities.SysRoleMenu{},
		&entities.SysApi{},
		&entities.SysMenu{},
	)
}
