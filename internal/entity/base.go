package entity

import (
	"time"
)

type Base struct {
	Id        uint64     `gorm:"primaryKey;column:id"`
	CreateAt  *time.Time `gorm:"column:create_at"`
	UpdateAt  *time.Time `gorm:"column:update_at"`
	CreateBy  uint64     `gorm:"column:create_by"`
	UpdatedBy uint64     `gorm:"column:update_by"`
}
