package internal

import (
	"github.com/voyager-go/start-go-api/global"
)

// BaseUser 用户信息表
type BaseUser struct {
	global.Model
	Nickname string `gorm:"column:nickname;type:varchar(80);NOT NULL;comment:昵称" json:"nickname"`   // 昵称
	Phone    string `gorm:"column:phone;type:varchar(11);NOT NULL;unique;comment:手机号" json:"phone"` // 手机号
	Password string `gorm:"column:password;type:varchar(200);NOT NULL;comment:密码" json:"password"`  // 密码
	Status   *int8  `gorm:"column:status;type:tinyint(1);NOT NULL;comment:状态" json:"status"`        // 状态
}
