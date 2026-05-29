package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProjectType int
type Status int

const (
	Pending    Status = iota
	InProgress        = 1
	Completed         = 2
)

type Order struct {
	ID           uuid.UUID
	clientName   string
	projectType  ProjectType
	status       Status
	deliveryDate time.Time
}

func createOrder(ctx *gin.Context) {

}

func retrieveOrder(ctx *gin.Context) {

}

func updateOrder(ctx *gin.Context) {

}

func deleteOrder(ctx *gin.Context) {

}

func main() {
	router := gin.Default()

	router.POST("/orders", createOrder)
	router.GET("/orders", retrieveOrder)
	router.PUT("/orders", updateOrder)
	router.DELETE("/orders", deleteOrder)

	router.Run("localhost:8080")
}
