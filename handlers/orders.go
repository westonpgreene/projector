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
