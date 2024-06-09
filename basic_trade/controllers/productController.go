package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Product := models.Product{}
	adminID := uint(adminData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.AdminID = adminID
	newUUID := uuid.New()
	Product.UUID = newUUID.String()

	err := db.Debug().Create(&Product).Error
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
}

func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	Products := []models.Product{}

	err := db.Debug().Find(&Products).Error // db.Debug().Preload("Admin").Find(&Products).Error
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

func GetProductbyID(ctx *gin.Context) {
	db := database.GetDB()
	Products := []models.Product{}
	productUUID := ctx.Param("productUUID")

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})

}
