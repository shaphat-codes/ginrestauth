package controller

import (
	"ginrestauth/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"ginrestauth/utils"
	"ginrestauth/database"
)

func Register(context *gin.Context) {
	var input models.User
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !utils.IsEmailUnique(input.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Sorry, email already exists"})
		return
	}

	verificationCode := utils.GenerateVerificationCode()

	user := models.User{
		Email: input.Email,
		LastName: input.LastName,
		FirstName: input.FirstName,
		VerificationCode: verificationCode,
		Password: input.Password,
	}
	savedUser, err := user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	utils.SendVerificationEmail(user.Email, user.VerificationCode)
		
	context.JSON(http.StatusCreated, gin.H{"message": "A verification code has been sent to the procided email address.", "created user": savedUser})
}

func VerifyEmail(context *gin.Context) {
	verificationCode := context.Param("code")
	verifiedUser, err := utils.FindUserByVerificationCode(verificationCode)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifiedUser.IsVerified = true
	verifiedUser.VerificationCode = ""

	database.Database.Save(verifiedUser)
	context.JSON(http.StatusOK, gin.H{"message": "Email verification successful.", "verifiedUser": verifiedUser})	
}

func Login(context *gin.Context) {
	var input models.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := models.FindUserByEmail(input.Email)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidatePassword(input.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	
	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
