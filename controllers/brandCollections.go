package controllers

import (
	"github.com/fatih/structs"
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/models"
)

// AddBrandCollection add collection to brand
func AddBrandCollection(ctx iris.Context) {

	brandID, _ := ctx.Params().GetUint("brandID")
	var brand models.Brand

	if err := config.DB.
		Where("id = ?", brandID).
		First(&brand).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	collectionInput := new(models.CollectionInput)
	if err := ctx.ReadJSON(collectionInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	var products []models.Product
	for i := 0; i < len(collectionInput.Products); i++ {
		var product models.Product
		if err := config.DB.
			Select("id, code").
			Where("id = ?", collectionInput.Products[i]).
			First(&product).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		products = append(products, product)
	}

	brandCollection := &models.BrandCollection{
		Name:         collectionInput.Name,
		Description:  collectionInput.Description,
		OtherDetails: collectionInput.OtherDetails,
		Brand:        brand,
	}

	if err := config.DB.Create(&brandCollection).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.
		Model(&brandCollection).
		Association("Products").
		Append(products).
		Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(APIResponse(true, brandCollection, "Collection Added!"))
}

// GetBrandCollections get all collections of brand
func GetBrandCollections(ctx iris.Context) {
	var brandCollections []*models.BrandCollection
	offset := ctx.URLParamIntDefault("offset", 0)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParamDefault("orderBy", "id")
	sort := ctx.URLParamDefault("sort", "")
	name := ctx.URLParam("name")

	brandID, _ := ctx.Params().GetUint("brandID")

	query := GetAll(name, orderBy, offset, limit, sort)
	if err := query.
		Preload("Brand").
		Where("brand_id = ?", brandID).
		Find(&brandCollections).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brandCollections, "Here your Collections!"))
}

// GetBrandCollection Get a brand's collection
func GetBrandCollection(ctx iris.Context) {
	var brandCollection models.BrandCollection
	brandCollectionID, _ := ctx.Params().GetUint("brandCollectionID")
	brandID, _ := ctx.Params().GetUint("brandID")

	if err := config.DB.
		Preload("Brand").
		Where("brand_id = ?", brandID).
		Where("id = ?", brandCollectionID).
		First(&brandCollection).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brandCollection, "Here your Collection!"))
}

// EditBrandCollection edit brand's collection
func EditBrandCollection(ctx iris.Context) {

	var brandCollection models.BrandCollection
	brandCollectionID, _ := ctx.Params().GetUint("brandCollectionID")
	brandID, _ := ctx.Params().GetUint("brandID")

	if err := config.DB.
		Preload("Brand").
		Where("brand_id = ?", brandID).
		Where("id = ?", brandCollectionID).
		First(&brandCollection).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	collectionInput := new(models.CollectionInput)
	if err := ctx.ReadJSON(collectionInput); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	newBrandID := collectionInput.BrandID
	if newBrandID > 0 && newBrandID != brandCollection.BrandID {
		var brand models.Brand
		if err := config.DB.
			Where("id = ?", newBrandID).
			First(&brand).
			Error; err != nil {
			ctx.StatusCode(GetErrorStatus(err))
			_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
			return
		}
		brandCollection.Brand = brand
	}

	if err := config.DB.
		Model(&brandCollection).
		Updates(structs.Map(collectionInput)).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &brandCollection, "Collection Updated!"))
}

// DeleteBrandCollection delete a Brand's Collection
func DeleteBrandCollection(ctx iris.Context) {
	var brandCollection models.BrandCollection
	brandCollectionID, _ := ctx.Params().GetUint("brandCollectionID")
	brandID, _ := ctx.Params().GetUint("brandID")

	if err := config.DB.
		Select("id, brand_id").
		Where("brand_id = ?", brandID).
		Where("id = ?", brandCollectionID).
		First(&brandCollection).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	if err := config.DB.Delete(&brandCollection).Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
	}

	_, _ = ctx.JSON(APIResponse(true, iris.Map{"deleted": true}, "Product Deleted!"))
}

// GetBrandCollectionProducts get products of brandCollection
func GetBrandCollectionProducts(ctx iris.Context) {
	offset := ctx.URLParamIntDefault("offset", 0)
	limit := ctx.URLParamIntDefault("limit", 15)
	orderBy := ctx.URLParamDefault("orderBy", "id")
	sort := ctx.URLParamDefault("sort", "")
	name := ctx.URLParam("name")

	var brandCollection models.BrandCollection
	var products []models.Product
	brandCollectionID, _ := ctx.Params().GetUint("brandCollectionID")
	brandID, _ := ctx.Params().GetUint("brandID")

	if err := config.DB.
		Select("id, brand_id").
		Where("brand_id = ?", brandID).
		Where("id = ?", brandCollectionID).
		First(&brandCollection).
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	query := GetAll(name, orderBy, offset, limit, sort)
	if err := query.
		Model(&brandCollection).
		Preload("Brand").
		Related(&products, "Products").
		Error; err != nil {
		ctx.StatusCode(GetErrorStatus(err))
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	_, _ = ctx.JSON(APIResponse(true, &products, "Here your Products!"))
}
