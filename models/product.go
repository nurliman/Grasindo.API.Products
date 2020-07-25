package models

import (
	"github.com/jinzhu/gorm"
)

// Product data model
type Product struct {
	gorm.Model
	Name         string  `json:"name" validate:"required" gorm:"size:100"`
	Code         *string `json:"code" gorm:"size:30;unique_index;not null" validate:"required"`
	Description  string  `json:"description" gorm:"size:255"`
	OtherDetails string  `json:"otherDetails" gorm:"size:255"`
	Brand        Brand   `json:"brand" gorm:"association_autoupdate:false;association_autocreate:false"`
	BrandID      uint    `json:"-"`
}

// ProductInput data model
type ProductInput struct {
	Name         string  `json:"name" validate:"required"`
	Code         *string `json:"code" validate:"required"`
	Description  string  `json:"description"`
	OtherDetails string  `json:"otherDetails"`
	BrandID      uint    `json:"brandID"`
}
