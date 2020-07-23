package main

import (
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/routes"

	"github.com/go-playground/validator/v10"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris/v12"
)

var err error

func main() {
	// Database connection initialization
	config.DBInit()

	// validator instances initialization
	v := validator.New()

	app := iris.Default()

	// register validator
	app.Validator = v

	// close db when interrupt
	iris.RegisterOnInterrupt(func() {
		_ = config.DB.Close()
	})

	// add routes
	routes.SetupRouter(app)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Listen("0.0.0.0:1337")
}
