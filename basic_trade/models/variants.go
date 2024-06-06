package models

import "time"

type Variant struct {
	ID          uint `gorm:"primaryKey"`
	UUID        string
	VariantName string `json:"variant_name"`
	Quantity    uint   `json:"quantity"`
	ProductID   uint   `json:"product_id"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
