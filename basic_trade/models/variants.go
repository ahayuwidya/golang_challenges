package models

import "time"

type Variant struct {
	ID          uint `gorm:"primaryKey"`
	UUID        string
	VariantName string
	Quantity    uint
	ProductID   uint
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
