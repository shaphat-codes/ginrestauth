package controller

import (
	"ginrestauth/utils"
	"ginrestauth/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	
)

func AddStore(context *gin.Context) {
	var input models.Store
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
	savedStore, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedStore})
}

func UpdateStore(context *gin.Context) {
	var input models.Store
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingStore, err := models.FindStoreById(input.ID)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	input.OwnerID = user.ID
	updatedStore, err := existingStore.Update(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": updatedStore})
}

func DetailStore(context *gin.Context) {
	var products models.Product
	var categories models.Category

	storeName := context.Param("store_name")

	_, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingStore, err := models.FindStoreByName(storeName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the category_id from the query parameter
    categoryID, err := strconv.Atoi(context.Query("category_id"))
    if err != nil {
        categoryID = 0
        
    }

	var storeProducts []models.Product
    if categoryID == 0 {
        storeProducts, err = products.ProductsForStore(storeName)
    } else {
        storeProducts, err = products.ProductsForStoreAndCategory(storeName, categoryID)
    }
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	storeCategories, err := categories.CategoriesForStore(storeName)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"storeDetails": existingStore, "storeProducts": storeProducts, "storeCategories": storeCategories})
}

func DeleteStore(context *gin.Context) {
	var input models.Store
	id := context.Param("id")
	
	_, err := utils.CurrentUser(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	deletedStore, err := input.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": deletedStore})
}