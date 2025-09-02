package middleware

import (
	"net/http"
	"strings"

	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
authHeader := context.GetHeader("Authorization")
if authHeader == "" {
    context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!"})
    return
}

token := strings.TrimPrefix(authHeader, "Bearer ")

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized!"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}