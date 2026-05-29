package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"src/projector/db"
	"src/projector/handlers"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	router := gin.Default()
	router.POST("/orders", handlers.CreateOrder)
	router.GET("/orders", handlers.GetOrders)
	router.PUT("/orders/:id", handlers.UpdateOrder)
	router.DELETE("/orders/:id", handlers.DeleteOrder)

	router.Run("localhost:8080")
}
