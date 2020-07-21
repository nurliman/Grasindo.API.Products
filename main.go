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


	// create any custom tcp listener, unix sock file or tls tcp listener.
	// l, err := net.Listen("tcp4", "0.0.0.0:8080")
	// if err != nil {
	// 	panic(err)
	// }

	// use of the custom listener.
	app.Listen("0.0.0.0:8080")
}
