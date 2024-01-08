package controller

import (
	"ginrestauth/utils"
	"ginrestauth/models"
	"github.com/gin-gonic/gin"
	"net/http"
	
)

func AddProduct(context *gin.Context) {
	var input models.Product
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	input.OwnerID = user.ID
	savedProduct, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedProduct})
}