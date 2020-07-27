package models

import (
	"github.com/jinzhu/gorm"
)

// Collection of products data model
type Collection struct {
	gorm.Model
	Name         string    `json:"name" validate:"required" gorm:"size:100"`
	Description  string    `json:"description" gorm:"size:255"`
	Products     []Product `json:"products" gorm:"many2many:collection_products;" validate:"dive" `
	OtherDetails string    `json:"otherDetails" gorm:"size:255"`
}

// BrandCollection Collection to specific Brand
type BrandCollection struct {
	Collection
	Products []Product `json:"products" gorm:"many2many:brand_collection_products;" validate:"dive" `
	BrandID  uint      `json:"brandID"`
}
