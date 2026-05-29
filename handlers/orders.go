package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"src/projector/db"
	"src/projector/models"
)

func CreateOrder(ctx *gin.Context) {
	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := models.Order{
		ClientName:   req.ClientName,
		ProjectType:  req.ProjectType,
		Status:       models.Pending,
		DeliveryDate: req.DeliveryDate,
	}
	db.DB.Create(&order)
	ctx.JSON(http.StatusCreated, order)
}

func GetOrders(ctx *gin.Context) {
	var orders []models.Order
	db.DB.Find(&orders)
	ctx.JSON(http.StatusOK, orders)
}

func UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	var req models.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Model(&order).Updates(req)
	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := db.DB.Delete(&models.Order{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "order deleted"})
}
