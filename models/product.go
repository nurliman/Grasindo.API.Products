package models

import (
	"github.com/jinzhu/gorm"
)

// Product data model
type Product struct {
	gorm.Model
	Name         string  `json:"name" binding:"required"`
	Code         *string `json:"code" gorm:"unique_index;not null" binding:"required"`
	Description  string  `json:"description"`
	OtherDetails string  `json:"otherDetails"`
	Brand        Brand   `json:"brand"`
	BrandID      uint
}
