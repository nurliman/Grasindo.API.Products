package main

import (
	//"net"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris/v12"
)

var err error

func main() {
	config.DB, err = gorm.Open("postgres", config.DBConfigBuilder())
	if err != nil {
		panic(err)
	}

	app := iris.Default()

	routes.SetupRouter(app)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Listen("0.0.0.0:1337")
}
