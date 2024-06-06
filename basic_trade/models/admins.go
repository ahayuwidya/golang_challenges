package models

import (
	"time"
)

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	UUID      string
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Products  []Product `gorm:"foreignKey:AdminID"`
}
