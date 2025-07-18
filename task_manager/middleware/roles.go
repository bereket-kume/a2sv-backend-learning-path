package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if role, exists := ctx.Get("role"); !exists || role != "admin" {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Admin Access only"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
