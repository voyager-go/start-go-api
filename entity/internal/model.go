package internal

// Model 公用字段
type Model struct {
	ID        int64 `gorm:"primary_key column:id" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"-"`
}
