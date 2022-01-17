package internal

import "github.com/voyager-go/start-go-api/global"

type SysApi struct {
	global.Model
	Path        string `gorm:"column:path;type:varchar(255);NOT NULL;comment:api访问路径" json:"path"`                                  // api访问路径
	Description string `gorm:"column:description;type:varchar(255);comment:api中文描述" json:"description"`                             // api中文描述
	Group       string `gorm:"column:group;type:varchar(30);NOT NULL;comment:api所在分组" json:"group"`                                 // api所在分组
	Method      int8   `gorm:"column:method;type:tinyint(1);NOT NULL;comment:方法 1:创建POST 2:查看GET 3:更新PUT 4:删除DELETE" json:"method"` // 方法 1:创建POST 2:查看GET 3:更新PUT 4:删除DELETE
	IsUse       *int8  `gorm:"column:is_use;type:tinyint(1);default:1;NOT NULL;comment:是否可用 0禁用 1启用" json:"is_use"`                 // 是否可用 0禁用 1启用
}
