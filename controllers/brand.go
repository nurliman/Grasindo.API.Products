package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/models"
)

// AddBrand add new brand
func AddBrand(ctx iris.Context) {
	brandJSON := new(models.Brand)

	if err := ctx.ReadJSON(brandJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.Create(&brandJSON).Error; err != nil {
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, brandJSON, "Brand Added!"))
}
