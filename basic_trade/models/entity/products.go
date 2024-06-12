package entity

import "time"

type Product struct {
	ID        uint `gorm:"primaryKey"`
	UUID      string
	Name      string `json:"name"`
	ImageURL  string `json:"image_url"`
	AdminID   uint   `json:"admin_id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Variants  []Variant `gorm:"foreignKey:ProductID"`
}