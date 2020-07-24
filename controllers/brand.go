package controllers

import (
	"github.com/fatih/structs"
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/models"
)

// AddBrand add new brand
func AddBrand(ctx iris.Context) {
	brandInput := new(models.BrandInput)

	if err := ctx.ReadJSON(brandInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	brand := &models.Brand{
		Name:         brandInput.Name,
		Description:  brandInput.Description,
		OtherDetails: brandInput.OtherDetails,
	}

	if err := config.DB.Create(brand).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(APIResponse(true, brand, "Brand Added!"))
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

// GetBrand return brand info by giving brand id
func GetBrand(ctx iris.Context) {
	var brand models.Brand
	brandId, _ := ctx.Params().GetUint("brandId")

	if err := config.DB.
		Where("id = ?", brandId).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brand, "Here your Brand!"))
}

// EditBrand edit brand by id
func EditBrand(ctx iris.Context) {
	var brand models.Brand
	brandId, _ := ctx.Params().GetUint("brandId")

	if err := config.DB.
		Select("id").
		Where("id = ?", brandId).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	brandInput := new(models.BrandInput)

	if err := ctx.ReadJSON(brandInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.
		Model(&brand).
		Updates(structs.Map(brandInput)).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brand, "Brand Updated!"))
}

// DeleteBrand delete a brand
func DeleteBrand(ctx iris.Context) {
	var brand models.Brand
	brandId, _ := ctx.Params().GetUint("brandId")

	if err := config.DB.
		Select("id").
		Where("id = ?", brandId).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.Delete(&brand).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
	}

	_, _ = ctx.JSON(APIResponse(true, iris.Map{"deleted": true}, "Brand Deleted!"))
}
