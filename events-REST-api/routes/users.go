package routes

import (
	"net/http"

	"example.com/events-api/models"
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

// func signUp(context *gin.Context) {
// 	var user models.User

// 	err := context.ShouldBind(&user)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse request data"})
// 		return
// 	}

// 	err = user.Save()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't create user. Try again later"})
// 		return
// 	}
// 	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
// }
func signUp(c *gin.Context) {
    var user models.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request", "error": err.Error()})
        return
    }

    if err := user.Save(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't create user", "error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}


func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse request data"})
		return
	}

	err = user.ValidateCredentials() 
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})
}