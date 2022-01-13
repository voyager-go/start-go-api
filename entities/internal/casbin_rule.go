package internal

type CasbinRule struct {
	ID     uint64 `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT"`
	Ptype  string `gorm:"column:ptype;type:varchar(100);default:'p'"`
	RoleID string `gorm:"column:v0;type:varchar(100)"`
	Path   string `gorm:"column:v1;type:varchar(100)"`
	Method string `gorm:"column:v2;type:varchar(100)"`
}
