package controllers

// func CreateProduct(ctx *gin.Context) {
// 	db := database.GetDB()
// 	Product := models.Product{}

// 	if err := ctx.ShouldBindJSON(&Product); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	err := db.Debug().Preload("ID").Create(&Product).Error
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error":   "Bad request",
// 			"message": err.Error(),
// 		})
// 	}
// }
