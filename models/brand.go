package models

import (
	"github.com/jinzhu/gorm"
)

// Brand data model
type Brand struct {
	gorm.Model
	Name         string    `json:"name" gorm:"unique_index;not null" validate:"required"`
	Description  string    `json:"description"`
	Products     []Product `json:"-" validate:"dive"`
	OtherDetails string    `json:"otherDetails"`
}

// BrandInput used when create or update data
type BrandInput struct {
	Name         string `json:"name" validate:"gte=2,required"`
	Description  string `json:"description"`
	OtherDetails string `json:"otherDetails"`
}
