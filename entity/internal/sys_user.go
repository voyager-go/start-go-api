package internal

// SysUser 用户信息表
type SysUser struct {
	Model
	Nickname string `gorm:"length:80,column:nickname" json:"nickname"`
	Phone    string `gorm:"unique,length:11,column:phone" json:"phone"`
	Password string `gorm:"length:80,column:password" json:"password"`
	Status   int8   `gorm:"index:,column:status" json:"status"`
}
