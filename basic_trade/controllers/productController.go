package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models/entity"
	"basic_trade/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)
	adminID := uint(adminData["id"].(float64))

	productReq := request.ProductRequest{}
	if contentType == appJSON {
		ctx.ShouldBindJSON(&productReq)
	} else {
		ctx.ShouldBind(&productReq)
	}

	productReq.AdminID = adminID
	newUUID := uuid.New()
	productReq.UUID = newUUID.String()

	// Extract the filename without extension
	fileName := helpers.RemoveExtension(productReq.ImageURL.Filename)
	uploadResult, err := helpers.UploadFile(productReq.ImageURL, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Product := entity.Product{
		UUID:     productReq.UUID,
		Name:     productReq.Name,
		ImageURL: uploadResult,
		AdminID:  productReq.AdminID,
	}

	err = db.Debug().Create(&Product).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Product,
	})

	// Product := models.Product{}

	// if contentType == appJSON {
	// 	ctx.ShouldBindJSON(&Product)
	// } else {
	// 	ctx.ShouldBind(&Product)
	// }

	// Product.AdminID = adminID
	// newUUID := uuid.New()
	// Product.UUID = newUUID.String()

	// err := db.Debug().Create(&Product).Error
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   "Bad request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// ctx.JSON(http.StatusOK, gin.H{
	// 	"data": Product,
	// })
}

// func GetProduct(ctx *gin.Context) {
// 	db := database.GetDB()
// 	Products := []models.Product{}

// 	err := db.Debug().Find(&Products).Error // db.Debug().Preload("Admin").Find(&Products).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request.",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": Products,
// 	})
// }

// func GetProductbyUUID(ctx *gin.Context) {
// 	db := database.GetDB()
// 	Products := []models.Product{}
// 	productUUID := ctx.Param("productUUID")

// 	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data": Products,
// 	})

// }

// func UpdateProductbyUUID(ctx *gin.Context) {
// 	db := database.GetDB()
// 	contentType := helpers.GetContentType(ctx)

// 	Products := []models.Product{}
// 	updatedProduct := models.Product{}
// 	productUUID := ctx.Param("productUUID")

// 	if contentType == appJSON {
// 		ctx.ShouldBindJSON(&updatedProduct)
// 	} else {
// 		ctx.ShouldBind(&updatedProduct)
// 	}

// 	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	err = db.Debug().Model(&Products).Where("uuid = ?", productUUID).Updates(&updatedProduct).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": err.Error(),
// 		})
// 	}
// }
