package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"src/projector/db"
	"src/projector/models"
)

func CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := models.Order{
		ClientName:   req.ClientName,
		ProjectType:  req.ProjectType,
		Status:       models.Pending,
		DeliveryDate: req.DeliveryDate,
	}
	if err := db.DB.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "order already exists"})
		return
	}
	ctx.JSON(http.StatusCreated, order)
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order
	if err := db.DB.Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve orders"})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order
	if err := db.DB.First(&order, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	var req models.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Model(&order).Updates(req).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update order"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	if result := db.DB.Where("id = ?", id).Delete(&models.Order{}); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	} else if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
