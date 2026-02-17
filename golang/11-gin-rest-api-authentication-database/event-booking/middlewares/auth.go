package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsrkator/event-booking/utils"
)

func AuthMiddleware(context *gin.Context) {
	token := context.Request.Header.Get("authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
