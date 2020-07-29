package models

import (
	"github.com/jinzhu/gorm"
)

// Collection of products data model
type Collection struct {
	gorm.Model
	Name         string    `json:"name" validate:"required" gorm:"size:100"`
	Description  string    `json:"description" gorm:"size:255"`
	Products     []Product `json:"-" gorm:"many2many:collection_products;" validate:"dive" `
	OtherDetails string    `json:"otherDetails" gorm:"size:255"`
}

// BrandCollection Collection to specific Brand
type BrandCollection struct {
	gorm.Model
	Name         string    `json:"name" validate:"required" gorm:"size:100"`
	Description  string    `json:"description" gorm:"size:255"`
	Products     []Product `json:"-" gorm:"many2many:brand_collection_products;" validate:"dive" `
	Brand        Brand     `json:"brand" gorm:"foreignkey:BrandID;association_autoupdate:false;association_autocreate:false"`
	BrandID      uint      `json:"brandID"`
	OtherDetails string    `json:"otherDetails" gorm:"size:255"`
}

// CollectionInput collection Input
type CollectionInput struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	Products     []int  `json:"products" validate:"required,min=1,unique"`
	OtherDetails string `json:"otherDetails"`
}
