package entity

import "time"

type Variant struct {
	ID          uint `gorm:"primaryKey"`
	UUID        string
	VariantName string `json:"variant_name" form:"variant_name"`
	Quantity    uint   `json:"quantity" form:"quantity"`
	ProductID   uint   `json:"product_id" form:"product_id"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
