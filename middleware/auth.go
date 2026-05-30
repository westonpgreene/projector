package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"src/projector/db"
	"src/projector/models"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		plaintext := ctx.GetHeader("X-API-Key")
		if plaintext == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing X-API-Key header"})
			return
		}

		var apiKey models.APIKey
		hash := models.HashKey(plaintext)
		if err := db.DB.Where("key_hash = ? AND active = ?", hash, true).First(&apiKey).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or inactive API key"})
			return
		}

		ctx.Next()
	}
}
