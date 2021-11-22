package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Orders struct {
	OrderId      uint   `gorm:"primaryKey"  json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item `gorm:"ForeignKey:OrderId" json:"items"`
}

func (orders *Orders) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(orders)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (orders *Orders) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(orders)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
