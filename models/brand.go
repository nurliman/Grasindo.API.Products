package models

import (
	"github.com/jinzhu/gorm"
)

// Brand data model
type Brand struct {
	gorm.Model
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description"`
	Products     []Product `json:"products" binding:"dive"`
	OtherDetails string    `json:"otherDetails"`
}
