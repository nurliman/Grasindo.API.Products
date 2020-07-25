package models

import (
	"github.com/jinzhu/gorm"
)

// Brand data model
type Brand struct {
	gorm.Model
	Name         string       `json:"name" gorm:"size:100;unique_index;not null" validate:"required"`
	Description  string       `json:"description" gorm:"size:255"`
	Products     []Product    `json:"-" validate:"dive"`
	Collections  []Collection `json:"collections" gorm:"many2many:brand_collections;" `
	OtherDetails string       `json:"otherDetails" gorm:"size:255"`
}

// BrandInput used when create or update data
type BrandInput struct {
	Name         string `json:"name" validate:"gte=2,required"`
	Description  string `json:"description"`
	OtherDetails string `json:"otherDetails"`
}
