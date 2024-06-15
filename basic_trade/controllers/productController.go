package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models/entity"
	"basic_trade/models/request"
	"fmt"
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

	if helpers.IsValidImageSize(int(productReq.ImageURL.Size)) {
		if helpers.IsValidImageExtension(productReq.ImageURL.Filename) {
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
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "File extension should be in JPG, JPEG, PNG or SVG.",
			})
		}

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "File size should be less than 5MB.",
		})
	}
}

func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	Products := []entity.Product{}

	err := db.Debug().Find(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request.",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})
}

func GetProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()
	Products := []entity.Product{}
	productUUID := ctx.Param("productUUID")
	fmt.Println("here 0", productUUID)

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	fmt.Println("here 1")

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})
}

func UpdateProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Products := []entity.Product{}
	updatedProduct := entity.Product{}
	productUUID := ctx.Param("productUUID")

	updatedProduct.AdminID = uint(adminData["id"].(float64))
	updatedProduct.UUID = productUUID

	if contentType == appJSON {
		ctx.ShouldBindJSON(&updatedProduct)
	} else {
		ctx.ShouldBind(&updatedProduct)
	}

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Model(&Products).Where("uuid = ?", productUUID).Updates(&updatedProduct).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedProduct,
	})
}

func DeleteProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Products := []entity.Product{}
	productToDelete := entity.Product{}
	productUUID := ctx.Param("productUUID")

	productToDelete.AdminID = uint(adminData["id"].(float64))
	productToDelete.UUID = productUUID

	if contentType == appJSON {
		ctx.ShouldBindJSON(&productToDelete)
	} else {
		ctx.ShouldBind(&productToDelete)
	}

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Model(&Products).Where("uuid = ?", productUUID).Delete(&productToDelete).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted record.",
	})
}
