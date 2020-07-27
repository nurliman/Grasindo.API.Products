package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/nurliman/Grasindo.API.Products/config"
	"github.com/nurliman/Grasindo.API.Products/models"
)

// AddBrandCollection add collection to brand
func AddBrandCollection(ctx iris.Context) {

	brandIDPath, _ := ctx.Params().GetUint("brandID")
	var brand models.Brand

	if err := config.DB.
		Where("id = ?", brandIDPath).
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

	brandCollection := &models.BrandCollection{
		Collection: models.Collection{
			Name:         collectionInput.Name,
			Description:  collectionInput.Description,
			OtherDetails: collectionInput.OtherDetails,
		},
		Products: collectionInput.Products,
		Brand:    brand,
	}

	if err := config.DB.Create(&brandCollection).Error; err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(APIResponse(false, nil, err.Error()))
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(APIResponse(true, brandCollection, "Collection Added!"))
}
