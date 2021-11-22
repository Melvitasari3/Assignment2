package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Item struct {
	ItemId      uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"order_id"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(item)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
func (item *Item) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(item)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
