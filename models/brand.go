package models

import (
	"github.com/jinzhu/gorm"
)

// Brand data model
type Brand struct {
	gorm.Model
	Name         string    `json:"name" validate:"required"`
	Description  string    `json:"description"`
	Products     []Product `json:"products" validate:"dive"`
	OtherDetails string    `json:"otherDetails"`
}
