package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string
	Items        []Item    // associate to struct Item
	OrderedAt    time.Time `gorm:"default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
