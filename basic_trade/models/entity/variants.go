package entity

import "time"

type Variant struct {
	ID          uint `gorm:"primaryKey"`
	UUID        string
	VariantName string `json:"variant_name" form:"variant_name" valid:"required,alpha"`
	Quantity    uint   `json:"quantity" form:"quantity" valid:"numeric"`
	ProductID   uint   `json:"product_id" form:"product_id"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
