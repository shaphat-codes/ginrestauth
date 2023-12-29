package main

import (
	"ginrestauth/database"
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
	
	router.Run(":8000")
	fmt.Println("server running on port 8000")
}

