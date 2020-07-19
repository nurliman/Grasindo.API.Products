package main

import (
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

func main() {

	config.DB, err = gorm.Open("postgres", config.DBConfigBuilder())

	if err != nil {
		panic("Failed to connect to database!")
	}

	defer config.DB.Close()

	config.DB.AutoMigrate()

	router := routes.SetupRouter()

	router.Run()
}
