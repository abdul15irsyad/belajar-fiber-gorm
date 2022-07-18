package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required,min=3" json:"name" form:"name"`
	Email     string `json:"email" validate:"required,email" json:"email" form:"email"`
	Age       int    `json:"age" validate:"required,number" json:"age" form:"age"`
	Address   string `json:"address" validate:"required" json:"address" form:"address"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
