package models

import "time"

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Quantity    int
	OrderID     uint // ID from struct Order `gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
