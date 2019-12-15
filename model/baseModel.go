package model

import (
	"time"
)

type BaseModel struct {
	ID        int        `gorm:"column:id;primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"index" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index" json:"deleted_at"`
}
