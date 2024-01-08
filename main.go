package main

import (
	"ginrestauth/database"
	"ginrestauth/middleware"
	"ginrestauth/models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"ginrestauth/controllers"
)

func main () {
	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Store{})
	database.Database.AutoMigrate(&models.Product{})
	database.Database.AutoMigrate(&models.Category{})
	database.Database.AutoMigrate(&models.Order{})

}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func serverApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.GET("verify-email/:code", controller.VerifyEmail)
	publicRoutes.POST("/login", controller.Login)
	publicRoutes.POST("/password-reset", controller.PasswordReset)
	publicRoutes.GET("/:password-reset-code", controller.PasswordResetConfirm)
	publicRoutes.POST("/:password-reset-code", controller.CreateNewPassword)

	protectedRoutes := router.Group("api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/store", controller.AddStore)
	protectedRoutes.PUT("/store", controller.UpdateStore)
	protectedRoutes.DELETE("/store/:id", controller.DeleteStore)
	protectedRoutes.GET("/store/:store_name", controller.DetailStore)

	protectedRoutes.POST("/product", controller.AddProduct)

	protectedRoutes.POST("/category", controller.AddCategory)
	
	router.Run(":8000")
	fmt.Println("server running on port 8000")
}

