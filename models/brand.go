package models

import (
	"github.com/jinzhu/gorm"
)

// Brand data model
type Brand struct {
	gorm.Model
	Name         string            `json:"name" gorm:"size:100;unique_index;not null" validate:"required"`
	Description  string            `json:"description" gorm:"size:255"`
	Products     []Product         `json:"-" gorm:"foreignkey:BrandID" validate:"dive"`
	Collections  []BrandCollection `json:"-" gorm:"foreignkey:BrandID"`
	OtherDetails string            `json:"otherDetails" gorm:"size:255"`
}

// BrandInput used when create or update data
type BrandInput struct {
	Name         string `json:"name" validate:"gte=2,required"`
	Description  string `json:"description"`
	OtherDetails string `json:"otherDetails"`
}
