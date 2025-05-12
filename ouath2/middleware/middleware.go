package middleware

import "github.com/gin-gonic/gin"

func CorsCredentials() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Credentials", "true")
		context.Next()
	}
}
