package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"index"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
