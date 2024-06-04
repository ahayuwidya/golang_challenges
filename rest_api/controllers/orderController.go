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

	err := db.Debug().First(&Orders, orderID).Error
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

func UpdateOrderbyID(ctx *gin.Context) { // fix items;
	db := database.GetDB()
	updatedOrder := models.Order{}
	Orders := []models.Order{}
	orderID := ctx.Param("orderID")

	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Debug().Model(&Orders).Where("id = ?", orderID).Updates(&updatedOrder).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated record.",
		"data":    Orders,
	})

}

func DeleteOrder(ctx *gin.Context) { // fix hooks di orders.go; fix index after deletion di return data
	db := database.GetDB()
	Orders := []models.Order{}
	orderID := ctx.Param("orderID")

	err := db.Debug().Delete(&Orders, orderID).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted record.",
		"data":    Orders,
	})
}
