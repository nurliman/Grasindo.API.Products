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

	var brandID uint = 0
	brandIDPath, _ := ctx.Params().GetUint("brandID")
	brandIDBody := productInput.BrandID

	if brandIDPath > 0 {
		brandID = brandIDPath
	} else if brandIDBody > 0 {
		brandID = brandIDBody
	}

	var brand models.Brand
	if err := config.DB.
		Where("id = ?", brandID).
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

	brandIDPath, _ := ctx.Params().GetUint("brandID")
	brandIDQuery := ctx.URLParamIntDefault("brandID", 0)

	if brandIDPath > 0 {
		query = query.Where("brand_id = ?", brandIDPath)
	} else if brandIDQuery > 0 {
		query = query.Where("brand_id = ?", brandIDQuery)
	}

	if err := query.Preload("Brand").Find(&products).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &products, "Here your Brands!"))
}
