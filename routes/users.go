package routes

import (
	"event-planning/models"
	"event-planning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleSignUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to read data"})
		return
	}

	err = user.SaveUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User saved"})
}

func handleGetUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Users returned", "data": users})
}

func handleLogin(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data"})
		return
	}

	err = user.Login()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{ "message": "Invalid credentials" })
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "message": "Login successful", "token": token })
}