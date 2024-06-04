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
	orders := []models.Order{}

	err := db.Debug().Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}
