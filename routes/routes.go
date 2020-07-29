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
		brands.Get("/{brandID:uint}", controllers.GetBrand)
		brands.Put("/{brandID:uint}", controllers.EditBrand)
		brands.Delete("/{brandID:uint}", controllers.DeleteBrand)

		brands.PartyFunc("/{brandID:uint}/products", func(products iris.Party) {
			products.Post("/", controllers.AddProduct)
			products.Get("/", controllers.GetProducts)
			products.Get("/{productID:uint}", controllers.GetProduct)
			products.Put("/{productID:uint}", controllers.EditProduct)
			products.Delete("/{productID:uint}", controllers.DeleteProduct)
		})

		brands.PartyFunc("/{brandID:uint}/collections", func(brandCollections iris.Party) {
			brandCollections.Post("/", controllers.AddBrandCollection)
			brandCollections.Get("/", controllers.GetBrandCollections)
			brandCollections.Get("/{brandCollectionID:uint}", controllers.GetBrandCollection)
			brandCollections.Put("/{brandCollectionID:uint}", controllers.EditBrandCollection)
			brandCollections.Delete("/{brandCollectionID:uint}", controllers.DeleteBrandCollection)

			brandCollections.Get("/{brandCollectionID:uint}/products", controllers.GetBrandCollectionProducts)
		})
	})

	v1.PartyFunc("/products", func(products iris.Party) {
		products.Post("/", controllers.AddProduct)
		products.Get("/", controllers.GetProducts)
		products.Get("/{productID:uint}", controllers.GetProduct)
		products.Put("/{productID:uint}", controllers.EditProduct)
		products.Delete("/{productID:uint}", controllers.DeleteProduct)
	})

}
