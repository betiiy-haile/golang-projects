package routes

import (
	"net/http"

	"example.com/events-api/models"
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Sign up a new user
// @Description Create a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /signup [post]
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

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
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