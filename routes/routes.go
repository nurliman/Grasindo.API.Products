package routes

import (
	"github.com/kataras/iris/v12"
)

// SetupRouter contains endpoint list and setting
func SetupRouter(app *iris.Application) {
	// endpoint /api/v1
	v1 := app.Party("/api/v1")

	v1.Handle("GET", "/products", func(ctx iris.Context) { ctx.JSON(iris.Map{"message": "yeay"}) })

}
