package models

import "time"

type Product struct {
	ID        uint `gorm:"primaryKey"`
	UUID      string
	Name      string
	ImageURL  string
	AdminID   uint
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Variants  []Variant `gorm:"foreignKey:ProductID"`
}
