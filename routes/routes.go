package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/controllers"
)

// SetupRouter contains endpoint list and setting
func SetupRouter(app *iris.Application) {
	// endpoint /api/v1
	v1 := app.Party("/api/v1")

	v1.Handle("GET", "/products", func(ctx iris.Context) { ctx.JSON(iris.Map{"message": "yeay"}) })

	v1.PartyFunc("/brands", func(brands iris.Party) {
		brands.Post("/", controllers.AddBrand)
		brands.Get("/", controllers.GetBrands)
		brands.Get("/{id:uint}", controllers.GetBrandByID)
	})
}
