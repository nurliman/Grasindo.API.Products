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

// GetBrands get brands
func GetBrands(ctx iris.Context) {
	var brands []*models.Brand
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	name := ctx.URLParam("name")
	orderBy := ctx.URLParam("orderBy")
	sort := ctx.URLParam("sort")

	query := GetAll(name, orderBy, offset, limit, sort)
	if err := query.Find(&brands).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brands, "Here your Brands!"))
}

// GetBrandByID give brand id return brand info
func GetBrandByID(ctx iris.Context) {
	var brand models.Brand
	id, _ := ctx.Params().GetUint("id")

	if err := config.DB.
		Where("id = ?", id).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brand, "Here your Brand!"))
}
