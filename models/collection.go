package models

import (
	"github.com/jinzhu/gorm"
)

// Collection of products data model
type Collection struct {
	gorm.Model
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description"`
	Products     []Product `json:"products" gorm:"many2many:collection_products;" validate:"dive" `
	OtherDetails string    `json:"otherDetails"`
}
