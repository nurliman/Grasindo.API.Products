package controllers

import (
	"github.com/fatih/structs"
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

	var brandID uint
	brandIDPath, _ := ctx.Params().GetUint("brandID")
	brandIDBody := productInput.BrandID

	if brandIDPath > 0 {
		brandID = brandIDPath
	} else if brandIDBody > 0 {
		brandID = brandIDBody
	} else {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, "brandID must provided in JSON body"))
		return
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
	_, _ = ctx.JSON(APIResponse(true, product, "Product Added!"))
}

// GetProducts get products
func GetProducts(ctx iris.Context) {
	var products []*models.Product
	offset := ctx.URLParamIntDefault("offset", 0)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParamDefault("orderBy", "id")
	sort := ctx.URLParamDefault("sort", "")
	name := ctx.URLParam("name")

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

	_, _ = ctx.JSON(APIResponse(true, &products, "Here your Products!"))
}

// GetProduct get a product by id
func GetProduct(ctx iris.Context) {
	var product models.Product
	productID, _ := ctx.Params().GetUint("productID")
	brandIDPath, _ := ctx.Params().GetUint("brandID")

	db := config.DB

	if brandIDPath > 0 {
		var brand models.Brand
		if err := config.DB.
			Select("id").
			Where("id = ?", brandIDPath).
			First(&brand).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		db = db.Where("brand_id = ?", brandIDPath)
	}

	if err := db.
		Preload("Brand").
		Where("id = ?", productID).
		First(&product).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &product, "Here your Product!"))
}

// EditProduct edit brand by id
func EditProduct(ctx iris.Context) {

	brandIDPath, _ := ctx.Params().GetUint("brandID")
	db := config.DB

	if brandIDPath > 0 {
		var brand models.Brand
		if err := config.DB.
			Select("id").
			Where("id = ?", brandIDPath).
			First(&brand).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		db = db.Where("brand_id = ?", brandIDPath)
	}

	var product models.Product
	productID, _ := ctx.Params().GetUint("productID")

	if err := db.
		Preload("Brand").
		Where("id = ?", productID).
		First(&product).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	productInput := new(models.ProductInput)
	if err := ctx.ReadJSON(productInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	brandIDBody := productInput.BrandID
	if brandIDBody > 0 && brandIDBody != product.BrandID {
		var brand models.Brand
		if err := config.DB.
			Where("id = ?", brandIDBody).
			First(&brand).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		product.Brand = brand
	}

	if err := config.DB.
		Model(&product).
		Updates(structs.Map(productInput)).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &product, "Product Updated!"))
}

// DeleteProduct delete a product
func DeleteProduct(ctx iris.Context) {
	var product models.Product
	productID, _ := ctx.Params().GetUint("productID")
	brandIDPath, _ := ctx.Params().GetUint("brandID")

	db := config.DB

	if brandIDPath > 0 {
		var brand models.Brand
		if err := config.DB.
			Select("id").
			Where("id = ?", brandIDPath).
			First(&brand).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		db = db.Where("brand_id = ?", brandIDPath)
	}

	if err := config.DB.
		Select("id").
		Where("id = ?", productID).
		First(&product).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
	}

	_, _ = ctx.JSON(APIResponse(true, iris.Map{"deleted": true}, "Product Deleted!"))
}
