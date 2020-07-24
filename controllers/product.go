package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/models"
)

// AddProduct add new product
func AddProduct(ctx iris.Context) {

	productInput := new(models.ProductInput)
	if err := ctx.ReadJSON(productInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	var brandId uint = 0
	brandIdPath, _ := ctx.Params().GetUint("brandId")
	brandIdBody := productInput.BrandID
	
	if  brandIdPath > 0 {
		brandId = brandIdPath
	} else if brandIdBody >0 {
		brandId = brandIdBody
	}
		
	var brand models.Brand
	if err := config.DB.
		Where("id = ?", brandId).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	product := &models.Product{
		Name:         productInput.Name,
		Code:         productInput.Code,
		Description:  productInput.Description,
		OtherDetails: productInput.OtherDetails,
		Brand:        brand,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(APIResponse(true, product, "Brand Added!"))
}

// GetProducts get products
func GetProducts(ctx iris.Context) {
	var products []*models.Product
	offset := ctx.URLParamIntDefault("offset", 1)
	limit := ctx.URLParamIntDefault("limit", 15)
	name := ctx.URLParam("name")
	orderBy := ctx.URLParam("orderBy")
	sort := ctx.URLParam("sort")

	query := GetAll(name, orderBy, offset, limit, sort)

	brandIdPath, _ := ctx.Params().GetUint("brandId")
	brandIdQuery := ctx.URLParamIntDefault("brandId",0)
	if  brandIdPath > 0 {
		query = query.Where("brand_id = ?", brandIdPath)
	} else if brandIdQuery > 0 {
		query = query.Where("brand_id = ?", brandIdQuery)
	}

	if err := query.Preload("Brand").Find(&products).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &products, "Here your Brands!"))
}
