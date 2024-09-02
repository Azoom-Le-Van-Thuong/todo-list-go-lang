package common

type SQLModel struct {
	Id        int    `json:"id,omitempty" gorm:"column:id"`
	CreatedAt string `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt string `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
