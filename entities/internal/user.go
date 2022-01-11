package internal

import (
	"github.com/voyager-go/start-go-api/global"
)

// User 用户信息表
type User struct {
	global.Model
	Nickname string `gorm:"length:80,column:nickname" json:"nickname"`
	Phone    string `gorm:"unique,length:11,column:phone" json:"phone"`
	Password string `gorm:"length:80,column:password" json:"password"`
	Status   *int8  `gorm:"index:,column:status" json:"status"`
}
