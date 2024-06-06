package models

import (
	"time"
)

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Products  []Product `gorm:"foreignKey:AdminID"`
}
