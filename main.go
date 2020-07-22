package main

import (
	//"net"
	"github.com/nurliman/Grasindo.API.Products/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	routes.SetupRouter(app)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Listen("0.0.0.0:1337")
}
