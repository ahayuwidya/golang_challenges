package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Item struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Quantity    int
	OrderID     uint // ID from struct Order `gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(i)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
