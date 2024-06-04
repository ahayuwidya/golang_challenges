package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"not null" json:"customerName"`
	Items        []Item    `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;" json:"items"` // associate to struct Item
	OrderedAt    time.Time `gorm:"default:null json:"orderedAt"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(o)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

// func (o *Order) BeforeDelete(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(o)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }
