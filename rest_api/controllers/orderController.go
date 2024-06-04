package controllers

import (
	"net/http"
	"rest_api/database"
	"rest_api/models"

	"github.com/gin-gonic/gin"
)

// var OrderData = []models.Order{}

func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	Order := models.Order{}

	if err := ctx.ShouldBindJSON(&Order); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Preload("ID").Create(&Order).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}
}

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}

	err := db.Debug().Find(&Orders).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Orders,
	})
}

func GetOrderbyID(ctx *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}
	orderID := ctx.Param("orderID")

	result := db.First(&Orders, orderID)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Orders,
	})
}

// func UpdateOrder() {

// }

// func DeleteOrder() {
// 	db := database.GetDB()

// }
