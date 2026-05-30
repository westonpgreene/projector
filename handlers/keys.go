package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"src/projector/db"
	"src/projector/models"
)

type CreateKeyRequest struct {
	Label string `json:"label" binding:"required"`
}

func CreateKey(ctx *gin.Context) {
	var req CreateKeyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plaintext, hash, err := models.GenerateKey()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate key"})
		return
	}

	apiKey := models.APIKey{
		ID:      uuid.New().String(),
		Label:   req.Label,
		KeyHash: hash,
		Active:  true,
	}
	if err := db.DB.Create(&apiKey).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save key"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"key":        plaintext,
		"label":      apiKey.Label,
		"created_at": apiKey.CreatedAt,
	})
}
