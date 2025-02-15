package middlewear

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseHeaderMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500") // Replace with your origin or '*'
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle OPTIONS method for preflight requests
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		// ctx.Next()
	}
}
