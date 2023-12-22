package model

import (
	"time"

	"gorm.io/gorm"
)

type Entity struct {
	ID        uint           `gorm:"primarykey;column:id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at" json:"deleted_at"`
	Name      string         `gorm:"column:name" json:"name" form:"name" binding:"required"`
	Data      string         `gorm:"column:data" json:"data" form:"data" binding:"required"`
}

func (Entity) TableName() string { return "entities" }
