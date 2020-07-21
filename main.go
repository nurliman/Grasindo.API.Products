package main

import (
	"net/http"
	"github.com/nurliman/Grasindo.API.Products/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	routes.SetupRouter(app)

	app.Handle("GET", "/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "pong"})
	})


	server := &http.Server{Addr:"0.0.0.0:55551"}
	app.Run(iris.Server(server))
}
