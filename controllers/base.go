package controllers

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/nurliman/Grasindo.API.Products/config"
)

// Response Structure
type Response struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// Lists Response list structure
type Lists struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

// APIResponse is helper for giving response
func APIResponse(status bool, objects interface{}, msg string) (r *Response) {
	r = &Response{Status: status, Data: objects, Msg: msg}
	return
}

// GetErrorStatus give error and return http status code
func GetErrorStatus(err error) int {
	if gorm.IsRecordNotFoundError(err) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}

// GetAll reusable func
func GetAll(string, orderBy string, offset, limit int, sort string) *gorm.DB {
	db := config.DB

	if len(orderBy) > 0 {
		db = db.Order(orderBy + " " + sort)
	} else {
		db = db.Order("created_at desc")
	}

	if len(string) > 0 {
		db = db.Where("name LIKE  ?", "%"+string+"%")
	}

	if limit > 0 {
		db = db.Limit(limit)
	}

	if offset > 0 {
		db = db.Offset(offset)
	}

	return db
}
